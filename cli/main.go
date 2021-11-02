package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/subiz/header"
	"encoding/hex"
)

func main() {
	data, err := hex.DecodeString("10e00118f601200e306438f00150ff0160f00168f001706e78f0018001ff018801f001a001f001a801f001b001f001b801e001c001e001c80140d001f001d80140e00140f001ff01f801f0018002fe018802f00190020f9802c001a00240a802f001")
	if err != nil {
		panic(err)
	}

	session := &header.LoginSession{}
	if err := proto.Unmarshal(data, session); err != nil {
		panic(err)
	}

	fmt.Println(session)
}
