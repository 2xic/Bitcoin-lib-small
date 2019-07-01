
package tx

import (
	"fmt"
	"encoding/binary"
)

type Transaction struct{
	version []byte // 4 bytes
	txInCount int
	inputs []byte
	txOutCount int
	outputs []byte
	locktime []byte // 4 bytes
}



type TransactionInput struct {
	hash []byte // 32 byte
	index []byte 

	scriptByteSize int // max 10 000
	signatureScript []byte 				//	more like the verificaiton script?
	sequence []byte // 4 byte
}

type OutputFormat struct {
	value []byte // int, max 8 byte
	ScriptBytesPK int // pubkey script bytes
	ScriptPK []byte

}

func FourByteAlign(number int) []byte{
	output := make([]byte, 4)
	binary.LittleEndian.PutUint32(output, uint32(number))
	return output
}

func EigthByteAlign(number int) []byte{
	output := make([]byte, 8)
	binary.LittleEndian.PutUint64(output, uint64(number))
	return output
}

func createOutput(value int, script []byte) OutputFormat{
	return OutputFormat{EigthByteAlign(value),
						len(script),
						script}
}

func createInput(txid []byte, outputIndex int, script []byte) TransactionInput{
	return TransactionInput{	txid,
								FourByteAlign(outputIndex),
								len(script),
								script,
								FourByteAlign(0xffffffff)}
}

func construct(inputs []TransactionInput, outputs[]OutputFormat) string{
	txSerailize := make([]byte, 0)
	txSerailize = append(txSerailize, FourByteAlign(1)...)
	txSerailize = append(txSerailize, byte(len(inputs)))

	for i:= 0; i < len(inputs); i++{
		txSerailize = append(txSerailize, inputs[i].hash...)
		txSerailize = append(txSerailize, inputs[i].index...)
		txSerailize = append(txSerailize, byte(inputs[i].scriptByteSize))
		txSerailize = append(txSerailize, inputs[i].signatureScript...)
		txSerailize = append(txSerailize, inputs[i].sequence...)
	}

	txSerailize = append(txSerailize, byte(len(outputs)))
	for i:= 0; i < len(outputs); i++{
		txSerailize = append(txSerailize, outputs[i].value...)
		txSerailize = append(txSerailize, byte(outputs[i].ScriptBytesPK))
		txSerailize = append(txSerailize, outputs[i].ScriptPK...)
	}

	txSerailize = append(txSerailize, FourByteAlign(0)...)
	return fmt.Sprintf("%x", txSerailize)
}	

