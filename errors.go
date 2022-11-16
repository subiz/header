package header

// This package lets you see which line of code has created an error along with its call stack.
//
//     err := readDatabase()
//     fmt.Println(err.(*errors.Error).Stack)
//
//
//    account/core/account.go:26
//    /vendor/github.com/subiz/header/account/account.pb.go:3306
//    /vendor/github.com/subiz/goutils/grpc/grpc.go:86
//    /vendor/github.com/subiz/goutils/grpc/grpc.go:87
//    /vendor/github.com/subiz/header/account/account.pb.go:3308
//    /vendor/google.golang.org/grpc/server.go:681

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// wrapErr converts an unknown error to an `*Error`, information of the
// old error stored in Root field.
// this function dont create a new error if the root error is already *Error
//   instead it attach class and code to the root (but only when those
//   properties of the root error is not defined.
//   Otherwise, this function create a brand new *Error.
func wrapErr(root error, class int, code E, v ...interface{}) *Error {
	if root == nil {
		// no root
		return newError(class, code, v)
	}

	mye, ok := root.(*Error)
	if !ok {
		// casting to err failed
		// dont give up yet, fallback to json
		errstr := root.Error()
		if strings.HasPrefix(errstr, "#ERR ") {
			roote := &Error{}
			if er := json.Unmarshal([]byte(errstr[len("#ERR "):]), roote); er == nil {
				if roote.Code != "" && roote.Class != 0 { // valid err
					return roote
				}
			}
		}

		e := newError(class, code, append(v, errstr))
		e.Root = root.Error()
		return e
	}

	if mye == nil {
		// no root
		return newError(class, code, v)
	}

	// if root existed, dont try to create a new error
	if mye.Code == "" {
		mye.Code = code.String()
	}

	if class != 0 && mye.Class == 0 {
		mye.Class = int32(class)
	}

	if len(v) > 0 {
		e := newError(class, code, v)
		mye.Description += "\n" + e.Description
	}
	return mye
}

// Errorf creates default *Error with description
// This method does not include stacktrace into returned object
func Errorf(format string, v ...interface{}) *Error {
	desc := fmt.Sprintf(format, v...)
	return &Error{Description: desc, Class: int32(500), Code: E_undefined.String()}
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func newError(class int, code E, v ...interface{}) *Error {
	var format, message string
	if len(v) == 0 {
		format = ""
	} else {
		var ok bool
		format, ok = v[0].(string)
		if !ok {
			format = strings.Repeat("%v", len(v))
		} else {
			v = v[1:]
		}
	}
	message = fmt.Sprintf(format, v...)

	e := &Error{}
	e.Description = message
	e.Class = int32(class)
	e.Stack = getStack(1)
	e.Code = code.String()
	return e
}

func E422(err error, code E, v ...interface{}) *Error {
	return wrapErr(err, 422, code, v...)
}

func E400(err error, code E, v ...interface{}) *Error {
	return wrapErr(err, 400, code, v...)
}

func E500(err error, code E, v ...interface{}) *Error {
	return wrapErr(err, 500, code, v...)
}

// FromString unmarshal an error string to *Error
func FromString(err string) *Error {
	if !strings.HasPrefix(err, "#ERR ") {
		return E500(nil, E_undefined, err)
	}
	e := &Error{}
	if er := json.Unmarshal([]byte(err[len("#ERR "):]), e); er != nil {
		return E500(er, E_invalid_json, "%s, %s", er, err)
	}
	return e
}

// Error returns string representation of an Error
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	b, _ := json.Marshal(e)
	return "#ERR " + string(b)
}

// getStack returns 10 closest stacktrace, included file paths and line numbers
// it will ignore all system path, path which is vendor is striped to /vendor/
// skip: number of stack ignored
func getStack(skip int) string {
	stack := make([]uintptr, 10)
	var sb strings.Builder
	// skip one system stack, the this current stack line
	length := runtime.Callers(4+skip, stack[:])
	for i := 0; i < length; i++ {
		pc := stack[i]
		// pc - 1 because the program counters we use are usually return addresses,
		// and we want to show the line that corresponds to the function call
		f := runtime.FuncForPC(pc)
		file, line := f.FileLine(pc - 1)
		// dont report system path

		file = trimToPrefix(file, "/vendor/")

		if isSystemPath(file) {
			continue
		}

		if isIgnorePath(file) {
			continue
		}

		// trim out common provider since most of go projects are hosted
		// in single host, there is no need to include them in the call stack
		// remove them help keeping the call stack smaller, navigatiing easier
		if !strings.HasPrefix(file, "/vendor") {
			file = trimOutPrefix(file, "/git.subiz.net/")
			file = trimOutPrefix(file, "/github.com/")
			file = trimOutPrefix(file, "/gopkg.in/")
		}

		sb.WriteString(file)
		sb.WriteString(":")
		sb.WriteString(strconv.Itoa(line))
		sb.WriteString(",")
	}
	return sb.String()
}

// isIgnorePath indicates whether a path is just noise, that excluding the path does not
// affect error context
func isIgnorePath(path string) bool {
	if strings.HasPrefix(path, "/vendor/google.golang.org/") {
		return true
	}

	if strings.HasPrefix(path, "/vendor/github.com/gin-gonic") {
		return true
	}

	if strings.HasSuffix(path, ".pb.go") {
		return true
	}
	return false
}

// isSystemPath tells whether a file is in system golang packages
func isSystemPath(path string) bool {

	return strings.HasPrefix(path, "/usr/local/go/src")
}

// trimToPrefix removes all the characters before the prefix
// its return the original string if not found prefix in str
func trimToPrefix(str, prefix string) string {
	i := strings.Index(str, prefix)
	if i < 0 {
		return str
	}
	return str[i:]
}

// trimOutPrefix removes all the characters before AND the prefix
// its return the original string if not found prefix in str
func trimOutPrefix(str, prefix string) string {
	i := strings.Index(str, prefix)
	if i < 0 {
		return str
	}
	return str[i+len(prefix):]
}

func GetErrCode(err error) string {
	myerr, ok := err.(*Error)
	if !ok {
		return E_undefined.String()
	}
	return myerr.Code
}
