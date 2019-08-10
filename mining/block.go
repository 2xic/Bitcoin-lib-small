package mining

import (
	"encoding/binary"
	"fmt"
	"github.com/2xic/bip-39/tx"
	"github.com/2xic/bip-39/secp256k1"
	"net/http"
	"io/ioutil"
	"math/big"
	"encoding/json"
)

type Blockheader struct{
	version []byte
	hashPrevBlock []byte
	hashMerkleRoot []byte
	Time []byte
	Bits []byte
	Nonce []byte
}

type TransactionInput struct {
	hash []byte 
	index []byte 

	signatureScript []byte 
	sequence []byte 
}

func FourByteAlign(number int) []byte{
	output := make([]byte, 4)
	binary.LittleEndian.PutUint32(output, uint32(number))
	return output
}

func revert(input []byte) []byte{
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func blockHeigthBIP34(heigt int) []byte{
	output := make([]byte, 8)
	binary.LittleEndian.PutUint64(output, uint64(heigt))
	return output[:compactSizeBIP(heigt)]
}

func compactSizeBIP(x int) int{
	n := 0;
	for (x != 0) {
		x >>= 8;
		n ++;
	}
	return n
}

func compactSize(x int) int{
	if(x < 0xFD){
		return 1;
	}else if(x < 0xFFFF){
		return 3;
	}else if(x < 0xFFFFFFFF){
		return 5;
	}else{
		return 9;
	}
}

func compactBytes(x int) []byte{
	output := make([]byte, 8)
	binary.LittleEndian.PutUint64(output, uint64(x))
	return output[:compactSize(x)]
}

/*
func generateCoinbase(script []byte) TransactionInput{
	return TransactionInput{
		make([]byte, 32),
		FourByteAlign(0xffffffff),
		script,
		FourByteAlign(0xffffffff)}
}
*/

type node struct{
	hash[]byte
	parent *node
}

func growMerkelTree(inputs []node) []node{
	if(len(inputs) == 1){
		return inputs
	}
	
	newLeafs := []node{}

	for i:= 0 ; i < len(inputs); i+=2{
		left := inputs[i]
		rigth := node{make([]byte, 32), nil}
		if((i + 1) < len(inputs)){
			rigth = inputs[i + 1]
		}else{
			rigth =  node{make([]byte, 32),
			nil}
		}

		parrentHash := append(left.hash, rigth.hash...)
		parrentNode := node{secp256k1.DoubleSha256(parrentHash), nil}

		left.parent = &parrentNode
		rigth.parent = &parrentNode

		newLeafs = append(newLeafs, parrentNode)
	}
	return growMerkelTree(newLeafs);
}

func txMerkelTree(txs []tx.Transaction) []node{
	nodes := []node{}
	for i:= 0; i < len(txs); i++{
		nodes = append(nodes, node{tx.TXid(txs[i]), nil})
	}
	merkelTree := growMerkelTree(nodes)
	return merkelTree
}

func serailizeBlockHeader(header Blockheader) string{
	blockSeralize := make([]byte, 0)
	blockSeralize = append(blockSeralize, header.version...)
	blockSeralize = append(blockSeralize, header.hashPrevBlock...)
	blockSeralize = append(blockSeralize, header.hashMerkleRoot...)
	blockSeralize = append(blockSeralize, header.Time...)
	blockSeralize = append(blockSeralize, header.Bits...)
	blockSeralize = append(blockSeralize, header.Nonce...)
	
	return fmt.Sprintf("%x", blockSeralize)
}

func generateExampleBlock(version int, previousBlock []byte, merkelRoot []byte, timestamp int, difficultiy int, nounce int, transactions []tx.Transaction) string{
	blockSeralize := make([]byte, 0)

	if(len(previousBlock) != 32){
		fmt.Println("error");
		return ""
	}

	if(len(merkelRoot) != 32){
		fmt.Println("error");
		return ""
	}

	blockSeralize = append(blockSeralize, FourByteAlign(version)...)
	blockSeralize = append(blockSeralize, previousBlock...)
	blockSeralize = append(blockSeralize, merkelRoot...)
	blockSeralize = append(blockSeralize, FourByteAlign(timestamp)...)
	blockSeralize = append(blockSeralize, FourByteAlign(difficultiy)...)
	blockSeralize = append(blockSeralize, FourByteAlign(nounce)...)	
	blockSeralize = append(blockSeralize, compactBytes(len(transactions))...)

	for i := 0; i < len(transactions); i++{
		blockSeralize = append(blockSeralize, tx.Serailizetx(transactions[i])...)
	}
	return fmt.Sprintf("%x", blockSeralize)
}

func generateBlockHash(block Blockheader) []byte{
	inputs := []byte{}
	inputs = append(inputs, block.version...)
	inputs = append(inputs, block.hashPrevBlock...)
	inputs = append(inputs, block.hashMerkleRoot...)
	inputs = append(inputs, block.Time...)
	inputs = append(inputs, block.Bits...)
	inputs = append(inputs, block.Nonce...)
	return secp256k1.DoubleSha256(inputs)
}

func mineBlock(block Blockheader, realTarget *big.Int, startNounce int, endNounce int) (found bool, responseBlock *Blockheader){
	difficultiyFound := new(big.Int)
	for i := startNounce; i < endNounce; i++{
		block.Nonce = FourByteAlign(i);
		difficultiyFound.SetBytes(revert(generateBlockHash(block)))

		if(difficultiyFound.Cmp(realTarget) == -1){
			return true, &block;
		}
	}
	return false, nil;
}

func requestUrl(url string) string{
	response, err := http.Get(url)
	if err != nil {
		panic("bad resposne")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("bad resposne")
	}

	return (string(body))
}

func requestBlock(block_hash string) (hash string, hashPrevious string, height float64, bits float64, version float64, time float64, nonce float64, merkel string){
	url := "https://blockchain.info/rawblock/" + block_hash
	fmt.Printf("Block url : %s\n", url)

	var result map[string]interface{}
	json.Unmarshal([]byte(requestUrl(url)), &result)

	height, bits, version, time, nonce = result["height"].(float64), result["bits"].(float64), result["ver"].(float64), result["time"].(float64), result["nonce"].(float64)

	return result["hash"].(string), result["prev_block"].(string), height, bits, version, time, nonce, result["mrkl_root"].(string)
}

func latestBlock() (hash string, hashPrevious string, height float64, bits float64, version float64, time float64, nonce float64, merkel string){
	latest_block_hash := requestUrl("https://blockchain.info/q/latesthash")
	return requestBlock(latest_block_hash)
}

func calculateDifficultyNumber(difficultiy float64) (*big.Int){
	maxTarget, _ := new(big.Int).SetString("00000000FFFF0000000000000000000000000000000000000000000000000000", 16)
	maxTargetFloat := new(big.Float).SetInt(maxTarget)

	target := big.NewFloat(difficultiy)
	target = target.Quo(maxTargetFloat, target)

	realTarget := new(big.Int)
	target.Int(realTarget)
	return realTarget
}

func calculateDifficultyBits(bits int) (*big.Int){
	hex := FourByteAlign(bits)
	hex = revert(hex)
	a := new(big.Int).SetBytes([]byte{hex[0]})
	b := new(big.Int).SetBytes(hex[1:])

	c := new(big.Int).Exp(big.NewInt(int64(2)), new(big.Int).Mul(big.NewInt(int64(8)), (new(big.Int).Sub(a, big.NewInt(int64(3))))), nil )

	result := new(big.Int).Mul(b,  c) //b)
	resultBytes := result.Bytes()

	target := make([]byte, 32-len(resultBytes))
	target = append(target, resultBytes...)

	targetNumber := new(big.Int).SetBytes(target)

	return targetNumber
}


