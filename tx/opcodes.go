package tx

import (
	"fmt"
	"bytes"
	"math"
	"github.com/2xic/bip-39/secp256k1"
	"crypto/sha1"
)

var (
	OP_RIPEMD160 =	166	
	OP_SHA1 =	167	
	OP_SHA256 =	168	
	OP_HASH160 =	169	
	OP_HASH256 =	170	
	OP_CODESEPARATOR =	171	
	OP_CHECKSIG =	172	
	OP_CHECKSIGVERIFY =	173	
	OP_CHECKMULTISIG =	174	
	OP_CHECKMULTISIGVERIFY =	175	

	OP_PUSHDATA1 = 76
	OP_PUSHDATA2 = 77

	OP_EQUALVERIFY	= 136
	OP_DUP = 118
	OP_NOP = 97

	OP_0 = 0
	OP_1 = 1
	OP_2 = 82
	OP_3 = 83
	OP_4 = 84
	OP_5 = 85
	OP_6 = 86
	OP_7 = 87
	OP_8 = 88
	OP_9 = 89
	OP_10 = 90
	OP_11 = 91
	OP_12 = 92
	OP_13 = 93
	OP_14 = 94
	OP_15 = 95
	OP_16 = 96
	
	OP_FALSE = 0

	OP_VERIFY = 105
	OP_2DUP = 110

	OP_SWAP = 124

	OP_NOT = 145

	OP_ADD = 147
	OP_SUB = 148
	OP_EQUAL = 135

)

type machine struct{
	OPs []int
	data [][]byte
}

func (machine *machine) PushOP(op int) {
	machine.OPs = append(machine.OPs, op)
}

func (machine *machine) PopOP() int{
	size := len(machine.OPs)
	if size == 0 {
		panic("Emty op")
	}
	res := machine.OPs[0]
	machine.OPs = machine.OPs[1:]
	return res
}


func (machine *machine) PushData(data []byte) {
	machine.data = append(machine.data, data)
}

func (machine *machine) PopData() ([]byte) {
	size := len(machine.data)
	if size == 0 {
		panic("Empty data")
	}
	res := machine.data[size-1]
	machine.data = machine.data[:size-1]
	return res
}

func(machine *machine) PeekData(i int) ([]byte) {
	size := len(machine.data)
	if size == 0 || (size - i) < 0{
		panic("Empty data")
	}
	res := machine.data[size-i]
	return res
}


func createScript() machine{
	return machine{make([]int, 0),make([][]byte, 0)}
}

func int2bytes(input uint64) []byte{
	bytes := make([]byte, 0)
	index := uint64(0)
	for math.Pow(2, float64(index)) <= float64(input){
		bytes = append(bytes, byte((input >> index) & 0xFF))
		index += 8
	}
	if(input == 0){
		bytes = append(bytes, byte(0))
	}
	return bytes
}

func bytes2int(input []byte) (results uint64){
	for i := 0; i < len(input); i++{
		results += (uint64(input[i]) << uint64(i * 8))
	}
	return results
}

func(machine *machine)popN(c int) (results []byte){
	for i := 0; i < int(c); i++{
		data := machine.PopOP()
		results = append(results, byte(data))
	}
	return results
}

func(machine *machine)execute(){
	for 0 < len(machine.OPs){
		OP := machine.PopOP()
		switch{
			case OP == OP_NOP:
				fmt.Println("OP_NOP does nothing")
			case OP == OP_0, OP == OP_FALSE:
				machine.PushData(int2bytes(0))
			case OP_1 == OP:
				machine.PushData(int2bytes(1))
			case (OP_2 <= OP && OP <= OP_16):
				machine.PushData(int2bytes(uint64(OP) - 80))
			case OP == OP_PUSHDATA2, 
				 OP == OP_PUSHDATA1,
				 (1 <= OP && OP <= 75):				
				if(OP == OP_PUSHDATA1){
					a := machine.PopOP()
					machine.PushData(machine.popN(int(a)))
				}else if(OP == OP_PUSHDATA2){
					a := machine.PopOP()
					b := machine.PopOP()
					machine.PushData(machine.popN(int(uint64(a) + (uint64(b) << 8))))
				}else if((1 <= OP && OP <= 75)){
					machine.PushData(machine.popN(OP))
				}

			case OP == OP_2DUP:
				b := machine.PeekData(1)
				a := machine.PeekData(2)

				machine.PushData(b)				
				machine.PushData(a)

			case OP == OP_NOT:
				if(bytes2int(machine.PopData()) == 0){
					machine.PushData(int2bytes(1))
				}else{
					machine.PushData(int2bytes(0))
				}
			case OP == OP_VERIFY:
				if(!(bytes2int(machine.PopData()) == 1)){
					panic("invalid tx")
				}
			case OP == OP_SHA1:
				hash := secp256k1.RunHash(machine.PopData(), sha1.New())
				machine.PushData(hash)

			case OP == OP_SWAP:
				b := machine.PopData()
				a := machine.PopData()

				machine.PushData(b)
				machine.PushData(a)

			case OP == OP_ADD, 
				 OP == OP_SUB:
				b := bytes2int(machine.PopData())
				a := bytes2int(machine.PopData())
				
				if(OP == OP_ADD){
					c := a + b
					machine.PushData(int2bytes(c))
				}else{
					c := a - b
					machine.PushData(int2bytes(c))
				}
			case OP == OP_EQUAL:
				b := machine.PopData()
				a := machine.PopData()
					
				if(bytes.Compare(a, b) == 0){
					machine.PushData(int2bytes(1))
				}else{
					machine.PushData(int2bytes(0))
				}
			default:
				fmt.Println(OP)
				panic("unkown OP")
		}
	}
}

func(machine *machine) valid() bool{
	return (bytes2int(machine.PeekData(1)) == 1)
}

func UnlockScript(lock []int, unlock[]int) bool{
	combined := unlock
	combined = append(combined, lock...)

	scripts := createScript()
	for i:= 0; i < len(combined); i++{
		scripts.PushOP(combined[i])
	}
	scripts.execute()
	return scripts.valid()
}

