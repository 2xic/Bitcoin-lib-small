package main

import (
	"fmt"
	"math/rand" 
	"strconv"
	"crypto/sha256"
	"errors"
	"github.com/2xic/bip-39/wordlist"
	"time"
	"strings"
)

func paddingByte(input string, size int) (output string){
	for i := len(input); i < size; i++{
		input = "0" + input 
	}
	return input
}

func generateRandomBytes(byteCount int) (randomBytes []byte){
	rand.Seed(time.Now().UnixNano())
	random := make([]byte, byteCount)
    rand.Read(random)
    return random
}

func generateMnemonicBytes(randomBytes []byte) (wordsOut string, err error){
	lengthBytes := len(randomBytes)

	if(lengthBytes < 16 || 32 < lengthBytes || lengthBytes % 4 != 0){
		return "", errors.New("invalid length on byte array")
	}

	bits := ""
	for i := 0; i < lengthBytes; i++ {
		bits += paddingByte(strconv.FormatInt(int64(randomBytes[i]), 2), 8)
	}

	checksum := generateChecksum(randomBytes)

	bits += checksum

	//words := ""
	for i := 0; i < len(bits); i+=11 {
		if 1 < i {
			wordsOut += " "
		}
		index, err := strconv.ParseInt(bits[i:i+11], 2, 64)
		if(err != nil){
			return "", errors.New("invalid byte array")
		}
		wordsOut += wordlist.Get(index)
	}
	return wordsOut, nil
}

func generateMnemonic() (words string, err error){
	strength := 256
	if strength % 32 != 0{
		return "", errors.New("error with strength size")
	}
	return generateMnemonicBytes(generateRandomBytes(strength / 8))
}


func generateChecksum(byteArray []byte) (checksum string){	
	hash := sha256.New()
	hash.Write(byteArray)
	return paddingByte(strconv.FormatInt(int64(hash.Sum(nil)[0]), 2), 8)
}

func splitString(start int, end int, input string) (output string){
	for i := start; i < end && i < len(input); i++{
		output += string(input[i])
	}
	return output
}

func verifyMnemonic(mnemonic string) (valid bool, err error){
	words := strings.Split(strings.TrimSpace(mnemonic), " ")
	bits := ""
	for i := 0; i < len(words); i++{
		bits += paddingByte(strconv.FormatInt(int64(wordlist.LookUp(words[i])), 2), 11)
	}

	entropySize := (len(bits) / 33) * 32
	entropy := splitString(0, entropySize, bits)
	checksum := splitString(entropySize, len(bits), bits)

	byteStream := make([]byte, 32)
	for i, j := 0, 0; i < len(entropy); i, j = i + 8, j + 1 {
		index, err := strconv.ParseInt(entropy[i:i+8], 2, 64)
		if(err != nil){
			return false, errors.New("error with entropy parsing")
			break
		}
		byteStream[j] = byte(index)
	}

	if(strings.Compare(generateChecksum(byteStream), checksum) == 0){
		return true, nil
	}
	return false, nil
}

func main(){
	words, err := generateMnemonic()
	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println(words)
	}
	valid, err := verifyMnemonic(words)
	if(err != nil){
		fmt.Println(err)
	}else{
		if(valid){
			fmt.Println("valid")
		}else{
			fmt.Println("invalid")
		}
	}
}

