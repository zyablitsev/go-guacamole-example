package main

import (
	"bytes"
	"strconv"
)

// EncodeInstructions translates list of instructions
// such as []string{"opcode", "param1", "param2"} to []byte slice
// according Guacamole proto: "6.opcode,6.param1,6.param2;"
func EncodeInstructions(instructions []string) []byte {
	buf := &bytes.Buffer{}
	for i := 0; i < len(instructions); i++ {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.Write(encode([]byte(instructions[i])))
	}
	buf.WriteByte(';')
	return buf.Bytes()
}

func encode(instruction []byte) []byte {
	buf := &bytes.Buffer{}
	buf.Write([]byte(strconv.Itoa(len(instruction))))
	buf.WriteByte('.')
	buf.Write(instruction)
	return buf.Bytes()
}

// DecodeInstructions translates "6.opcode,6.param1,6.param2;" []byte slice
// of instructions in Guacamole proto format
// to list []string{"opcode", "param1", "param2"}
func DecodeInstructions(instructions []byte) []string {
	strings := make([]string, 0, 10) // preallocate try prevent extra allocations
	begin := 0
	for i := 0; i < len(instructions); i++ {
		if instructions[i] == ',' || instructions[i] == ';' {
			strings = append(strings, string(decode(instructions[begin:i])))
			begin = i
		}
	}
	return strings
}

func decode(instruction []byte) []byte {
	return instruction[bytes.IndexByte(instruction, '.')+1:]
}
