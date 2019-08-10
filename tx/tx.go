
package tx

import (
	"fmt"
	"encoding/binary"
	"encoding/hex"
	"github.com/2xic/bip-39/secp256k1"
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


func inputs_and_outputs_bytes(inputs []TransactionInput, outputs[]OutputFormat) (inputs_bytes []byte, outputs_bytes []byte){
	
	for i:= 0; i < len(inputs); i++{
		inputs_bytes = append(inputs_bytes, inputs[i].hash...)
		inputs_bytes = append(inputs_bytes, inputs[i].index...)
		inputs_bytes = append(inputs_bytes, byte(inputs[i].scriptByteSize))
		inputs_bytes = append(inputs_bytes, inputs[i].signatureScript...)
		inputs_bytes = append(inputs_bytes, inputs[i].sequence...)
	}

	for i:= 0; i < len(outputs); i++{
		outputs_bytes = append(outputs_bytes, outputs[i].value...)
		outputs_bytes = append(outputs_bytes, byte(outputs[i].ScriptBytesPK))
		outputs_bytes = append(outputs_bytes, outputs[i].ScriptPK...)
	}
	return inputs_bytes, outputs_bytes
}

func construct(inputs []TransactionInput, outputs[]OutputFormat) string{
	txSerailize := make([]byte, 0)
	txSerailize = append(txSerailize, FourByteAlign(1)...)
	txSerailize = append(txSerailize, byte(len(inputs)))
/*
	for i:= 0; i < len(inputs); i++{
		txSerailize = append(txSerailize, inputs[i].hash...)
		txSerailize = append(txSerailize, inputs[i].index...)
		txSerailize = append(txSerailize, byte(inputs[i].scriptByteSize))
		txSerailize = append(txSerailize, inputs[i].signatureScript...)
		txSerailize = append(txSerailize, inputs[i].sequence...)
	}
*/
	inputs_bytes, outputs_bytes := inputs_and_outputs_bytes(inputs, outputs)

	txSerailize = append(txSerailize, inputs_bytes...)

	txSerailize = append(txSerailize, byte(len(outputs)))

	txSerailize = append(txSerailize, outputs_bytes...)

/*	for i:= 0; i < len(outputs); i++{
		txSerailize = append(txSerailize, outputs[i].value...)
		txSerailize = append(txSerailize, byte(outputs[i].ScriptBytesPK))
		txSerailize = append(txSerailize, outputs[i].ScriptPK...)
	}
*/
	txSerailize = append(txSerailize, FourByteAlign(0)...)
	return fmt.Sprintf("%x", txSerailize)
}

func Serailizetx(tx Transaction) []byte{
	txSerailize := make([]byte, 0)
	txSerailize = append(txSerailize, tx.version...)
	txSerailize = append(txSerailize, byte(tx.txInCount))

	//for i:= 0; i < len(tx.inputs); i++{
	txSerailize = append(txSerailize, tx.inputs...)//.hash...)
	//	txSerailize = append(txSerailize, tx.inputs[i].index...)
	//	txSerailize = append(txSerailize, byte(tx.inputs[i].scriptByteSize))
	//	txSerailize = append(txSerailize, tx.inputs[i].signatureScript...)
	//	txSerailize = append(txSerailize, tx.inputs[i].sequence...)
//	}

	txSerailize = append(txSerailize, byte(tx.txOutCount))
	//for i:= 0; i < len(tx.outputs); i++{
	txSerailize = append(txSerailize, tx.outputs...)//[i]...)//.value...)
	////	txSerailize = append(txSerailize, byte(tx.outputs[i].ScriptBytesPK))
	//	txSerailize = append(txSerailize, tx.outputs[i].ScriptPK...)
//	}

	txSerailize = append(txSerailize, tx.locktime...)
	return txSerailize

	//return fmt.Sprintf("%x", txSerailize)
}	


func TXid(tx Transaction) []byte{
	return secp256k1.DoubleSha256(Serailizetx(tx));
}


func serailizetxstring(tx Transaction) string{
	return fmt.Sprintf("%x", Serailizetx(tx))
}

func String2Bytes(tx_string string) Transaction{
	tx_bytes, _ := hex.DecodeString(tx_string)
	
	version := tx_bytes[:4]
	inputs_count := int(tx_bytes[4])


	index := 5
	inputs := []TransactionInput{}

	for i:= 0; i < inputs_count; i++{
		hash := tx_bytes[index:index+32]
		index += 32

		block_index := tx_bytes[index:index+4]
		index += 4

		script_size := int(tx_bytes[index])
		index += 1

		script := tx_bytes[index:index+script_size]

		index += script_size

		//fmt.Println(script_size)
		sequence := tx_bytes[index : index + 4]
		index += 4

		input := TransactionInput{
			hash,
			block_index,
			script_size,
			script,
			sequence}
		inputs = append(inputs, input)
	}

	//fmt.Println("YO")
	//fmt.Println(inputs_count)
	//fmt.Println(index)
	outputs_count := int(tx_bytes[index])
	index += 1
	//fmt.Println(outputs_count)
	//fmt.Println("BAGEL")

//	Exit(0)

	outputs := []OutputFormat{}
	for i:= 0; i < outputs_count; i++{
		value := tx_bytes[index:index+8]
		index += 8
		script_size := int(tx_bytes[index])
		index += 1
		

		script := tx_bytes[index:index+script_size]

		index += script_size

		outputs = append(outputs, OutputFormat{
			value,
			script_size,
			script})
	}

	//fmt.Println(len(outputs))
	inputsBytes, outputsBytes := inputs_and_outputs_bytes(inputs, outputs)


	/*
	inputsBytes := []byte{}
	for i:= 0; i < len(inputs); i++{
		inputsBytes = append(inputsBytes, inputs[i].hash...)
		inputsBytes = append(inputsBytes, inputs[i].index...)
		inputsBytes = append(inputsBytes, byte(inputs[i].scriptByteSize))
		inputsBytes = append(inputsBytes, inputs[i].signatureScript...)
		inputsBytes = append(inputsBytes, inputs[i].sequence...)
	}
	outputsBytes := []byte{}
	for i:= 0; i < len(outputs); i++{
		outputsBytes = append(outputsBytes, outputs[i].value...)
		outputsBytes = append(outputsBytes, byte(outputs[i].ScriptBytesPK))
		outputsBytes = append(outputsBytes, outputs[i].ScriptPK...)
	}*/

	locktime := tx_bytes[index:index + 4]

	new := Transaction{
		version,
		inputs_count,
		inputsBytes,
		outputs_count,
		outputsBytes,
		locktime}

	return new
}







