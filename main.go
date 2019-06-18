package main

import (
	"fmt"
	"math/rand" 
	"strconv"
	"crypto/sha256"
	"errors"
	"github.com/2xic/bip-39/wordlist"
	"time"
)

func padding_byte(input string, size int) (output string){
	for i := len(input); i < size; i++{
		input = "0" + input 
	}
	return input
}

func generate_random_bytes(byte_count int) (random_bytes []byte){
	rand.Seed(time.Now().UnixNano())
	random := make([]byte, byte_count)
    rand.Read(random)
    return random
}

func generate_memonic_bytes(random_bytes []byte) (words_out string, err error){
	length_bytes := len(random_bytes)

	if(length_bytes < 16 || 32 < length_bytes || length_bytes % 4 != 0){
		return "", errors.New("invalid length on byte array")
	}

	bits := ""
	for i := 0; i < length_bytes; i++ {
		bits += padding_byte(strconv.FormatInt(int64(random_bytes[i]), 2), 8)
	}

	checksum := generate_checksum(random_bytes)
	bits += checksum

	words := ""
	for i := 0; i < len(bits); i+=11 {
		if 1 < i {
			words += " "
		}
		index, err := strconv.ParseInt(bits[i:i+11], 2, 64)
		if(err != nil){
			return "", errors.New("invalid byte array")
		}
		words += wordlist.Get(index)
	}
	return words, nil
}

func generate_memonic() (words string, err error){
	strength := 256
	if strength % 32 != 0{
		return "", errors.New("error with strength size")
	}
	return generate_memonic_bytes(generate_random_bytes(strength / 8))
}


func generate_checksum(byte_array []byte) (checksum string){	
	hash := sha256.New()
	hash.Write(byte_array)
	return padding_byte(strconv.FormatInt(int64(hash.Sum(nil)[0]), 2), 8)
}

func main(){
	words, err := generate_memonic()
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(words)
}

