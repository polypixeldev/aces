//go:build js && wasm

package main

import (
	"strings"
	"syscall/js"

	"bytes"

	"github.com/quackduck/aces"
)

func main() {
	js.Global().Set("encode", js.FuncOf(encode))
	js.Global().Set("decode", js.FuncOf(decode))
	select {}
}

func encode(this js.Value, args []js.Value) any {
	charset := args[0].String()
	byte1 := args[1].Int()
	byte2 := args[2].Int()
	byte3 := args[3].Int()
	byte4 := args[4].Int()
	ipBytes := []byte{byte(byte1), byte(byte2), byte(byte3), byte(byte4)}
	c, err := aces.NewCoding([]rune(charset))
	if err != nil {
		return err.Error()
	}

	output := new(bytes.Buffer)

	err = c.Encode(output, bytes.NewReader(ipBytes))
	if err != nil {
		return err.Error()
	}

	return output.String()
}

func decode(this js.Value, args []js.Value) any {
	charset := args[0].String()
	dataStr := args[1].String()

	c, err := aces.NewCoding([]rune(charset))
	if err != nil {
		return err.Error()
	}

	output := new(bytes.Buffer)

	err = c.Decode(output, strings.NewReader(dataStr))
	if err != nil {
		return err.Error()
	}

	return output.String()
}