package main

import (
	"bytes"
	"testing"
)

var encodeDecodeTests = [][2]string{
	{"opcode", "6.opcode"},
	{"param1", "6.param1"},
	{"param2", "6.param2"},
}

func TestEncode(t *testing.T) {
	for _, test := range encodeDecodeTests {
		result := encode([]byte(test[0]))
		if bytes.Compare(result, []byte(test[1])) != 0 {
			t.Errorf("Encode %s expected %q got %q", test[0], test[1], string(result))
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range encodeDecodeTests {
		result := decode([]byte(test[1]))
		if bytes.Compare(result, []byte(test[0])) != 0 {
			t.Errorf("Decode %s expected %q got %q", test[1], test[0], string(result))
		}
	}
}
