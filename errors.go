package header

import (
	"strconv"

	"github.com/subiz/log"
)

func E400(err error, code E, v ...interface{}) error {
	field := log.M{}
	for i, vv := range v {
		field[strconv.Itoa(i)] = vv
	}

	return log.Error(err, field, log.E_invalid_input, log.E(code.String()))
}

func E500(err error, code E, v ...interface{}) error {
	field := log.M{}
	for i, vv := range v {
		field[strconv.Itoa(i)] = vv
	}

	return log.Error(err, field, log.E_internal, log.E(code.String()))
}
