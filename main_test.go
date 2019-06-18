package main

import (
	"testing"
)

func generate_set_bytes(value byte) (stream []byte){
	byte_stream := make([]byte, 32)
	for i:= 0; i < len(byte_stream); i++{
		byte_stream[i] = value
	}
	return byte_stream
}

func TestSum(t *testing.T) {
	words, err := generate_memonic_bytes(generate_set_bytes(0))
	if err != nil{
		t.Errorf("Error with vector 0")		
	}
	if words != "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art"{
		t.Errorf("Error with vector 0")
	}

	words, err = generate_memonic_bytes(generate_set_bytes(0x7F))
	if err != nil{
		t.Errorf("Error with vector 1")		
	}
	if words != "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth title"{
		t.Errorf("Error with vector 1, got %s", words)
	}

	words, err = generate_memonic_bytes(generate_set_bytes(0x80))
	if err != nil{
		t.Errorf("Error with vector 2")		
	}
	if words != "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless"{
		t.Errorf("Error with vector 2, got %s", words)
	}

	words, err = generate_memonic_bytes(generate_set_bytes(0xFF))
	if err != nil{
		t.Errorf("Error with vector 2")		
	}
	if words != "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote"{
		t.Errorf("Error with vector 2, got %s", words)
	}
}