
package mining

import (
	"testing"
	"encoding/hex"
	"strings"
	"bytes"
	"github.com/2xic/bip-39/tx"
	"encoding/binary"
)


func getTestBlock() Blockheader{
	prevBlock, _ := hex.DecodeString("0000000000000000a41299d0230519d425dbd368d2b5937485aa20741e13a020")
	merkel, _ := hex.DecodeString("181253a22d7a9287fe9b38b71d724ff51de008e99dda6610ed1b9d2676178c0b")


	testBlock := Blockheader{
		FourByteAlign(2),
		revert(prevBlock),
		revert(merkel),
		FourByteAlign(1396028577),
		FourByteAlign(419486617),
		FourByteAlign(2997251652)}
	
	return testBlock
}

func getTestBlockNew() Blockheader{
	prevBlock, _ := hex.DecodeString("0000000000000000001b89552219ad3fff373030a992f62f95563b578479481e")
	merkel, _ := hex.DecodeString("6dbc4a137442adbe87e8b9ccbb35cf9593692ba504f916121a16d32406086601")


	testBlock := Blockheader{
		FourByteAlign(0x20000000),
		revert(prevBlock),
		revert(merkel),
		FourByteAlign(1565359033),
		FourByteAlign(387723321),
		FourByteAlign(2763111713)}
	
	return testBlock
}

func Test_Merkel(t *testing.T){
	//	https://www.blockchain.com/btc/block/000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506	

	txs := []tx.Transaction{}
	txs = append(txs, tx.String2Bytes("01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff08044c86041b020602ffffffff0100f2052a010000004341041b0e8c2567c12536aa13357b79a073dc4444acb83c4ec7a0e2f99dd7457516c5817242da796924ca4e99947d087fedf9ce467cb9f7c6287078f801df276fdf84ac00000000"))
	txs = append(txs, tx.String2Bytes("0100000001032e38e9c0a84c6046d687d10556dcacc41d275ec55fc00779ac88fdf357a187000000008c493046022100c352d3dd993a981beba4a63ad15c209275ca9470abfcd57da93b58e4eb5dce82022100840792bc1f456062819f15d33ee7055cf7b5ee1af1ebcc6028d9cdb1c3af7748014104f46db5e9d61a9dc27b8d64ad23e7383a4e6ca164593c2527c038c0857eb67ee8e825dca65046b82c9331586c82e0fd1f633f25f87c161bc6f8a630121df2b3d3ffffffff0200e32321000000001976a914c398efa9c392ba6013c5e04ee729755ef7f58b3288ac000fe208010000001976a914948c765a6914d43f2a7ac177da2c2f6b52de3d7c88ac00000000"))
	txs = append(txs, tx.String2Bytes("0100000001c33ebff2a709f13d9f9a7569ab16a32786af7d7e2de09265e41c61d078294ecf010000008a4730440220032d30df5ee6f57fa46cddb5eb8d0d9fe8de6b342d27942ae90a3231e0ba333e02203deee8060fdc70230a7f5b4ad7d7bc3e628cbe219a886b84269eaeb81e26b4fe014104ae31c31bf91278d99b8377a35bbce5b27d9fff15456839e919453fc7b3f721f0ba403ff96c9deeb680e5fd341c0fc3a7b90da4631ee39560639db462e9cb850fffffffff0240420f00000000001976a914b0dcbf97eabf4404e31d952477ce822dadbe7e1088acc060d211000000001976a9146b1281eec25ab4e1e0793ff4e08ab1abb3409cd988ac00000000"))
	txs = append(txs, tx.String2Bytes("01000000010b6072b386d4a773235237f64c1126ac3b240c84b917a3909ba1c43ded5f51f4000000008c493046022100bb1ad26df930a51cce110cf44f7a48c3c561fd977500b1ae5d6b6fd13d0b3f4a022100c5b42951acedff14abba2736fd574bdb465f3e6f8da12e2c5303954aca7f78f3014104a7135bfe824c97ecc01ec7d7e336185c81e2aa2c41ab175407c09484ce9694b44953fcb751206564a9c24dd094d42fdbfdd5aad3e063ce6af4cfaaea4ea14fbbffffffff0140420f00000000001976a91439aa3d569e06a1d7926dc4be1193c99bf2eb9ee088ac00000000"))
	merkel := txMerkelTree(txs)[0].hash

//	fmt.Printf("%x", revert(merkel))

	true_merkel, _ := hex.DecodeString("f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766")

	if(!(bytes.Compare(revert(merkel), true_merkel) == 0)){
		t.Errorf("Error with merkel tree")
	}

}


func Test_Block(t *testing.T) {
	//	Testing on https://blockchain.info/block/00000000000000005ca6049e552e76ea082c8f0f6b6d94377b5c2857ae66d750?format=hex
	prevBlock, _ := hex.DecodeString("0000000000000000a41299d0230519d425dbd368d2b5937485aa20741e13a020")
	prevBlock = revert(prevBlock)

	version := 2
	timestamp := 1396028577
	difficultiy := 419486617
	nounce := 2997251652
	
	txs := []tx.Transaction{}
	txs = append(txs, tx.String2Bytes("01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff54034e78040d00456c6967697573005335b491fabe6d6d9c033cedbbbc5433e23248d308cdd869a5ca97a20599ba0f1382f835b2c382cf0400000000000000002f7330746573742f0050bf7902df7f000032000000ffffffff0100f90295000000001976a9145399c3093d31e4b0af4be1215d59b857b861ad5d88ac00000000"))

	merkel := txMerkelTree(txs)[0].hash


	block := generateExampleBlock(version,
									prevBlock,
									merkel,
									timestamp,
									difficultiy,
									nounce,
									txs)

	if(!(strings.Compare(block, "0200000020a0131e7420aa857493b5d268d3db25d4190523d09912a400000000000000000b8c1776269d1bed1066da9de908e01df54f721db7389bfe87927a2da2531218a1b4355399db0019446ea6b20101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff54034e78040d00456c6967697573005335b491fabe6d6d9c033cedbbbc5433e23248d308cdd869a5ca97a20599ba0f1382f835b2c382cf0400000000000000002f7330746573742f0050bf7902df7f000032000000ffffffff0100f90295000000001976a9145399c3093d31e4b0af4be1215d59b857b861ad5d88ac00000000") == 0)){
		t.Errorf(block)
		t.Errorf("Error with block")
	}

	if(!(bytes.Compare(blockHeigthBIP34(416236), []byte{0xec,0x59,0x06}) == 0) ){
		t.Errorf("Error with bip34 block heigth")		
	}
}

func Test_Hashing(t *testing.T){
	testBlock := getTestBlock()

	trueHash, _ := hex.DecodeString("00000000000000005ca6049e552e76ea082c8f0f6b6d94377b5c2857ae66d750")

	if(!(bytes.Compare(revert(generateBlockHash(testBlock)), trueHash) == 0)){
		t.Errorf("Error with merkel tree")
	}	
}

func Test_HashingDifficulty(t *testing.T){
	testBlock := getTestBlock()

	foundBlock, parsedBlock := mineBlock(testBlock, calculateDifficultyNumber(5006860589.29), 2997251650, 2997251655)
	if(!foundBlock){
		t.Errorf("Error with parsing difficultiy?")
	}

	calculatedNounce := binary.BigEndian.Uint32(parsedBlock.Nonce)
	if(!(calculatedNounce != 2997251652)){
		t.Errorf("Error with parsing POW?")
	}

}

func Test_HashingDifficultyBits(t *testing.T){
	testBlock := getTestBlock()

	foundBlock, parsedBlock := mineBlock(testBlock, calculateDifficultyBits(419486617), 2997251650, 2997251655)
	if(!foundBlock){
		t.Errorf("Error with parsing difficultiy?")
	}

	calculatedNounce := binary.BigEndian.Uint32(parsedBlock.Nonce)
	if(!(calculatedNounce != 2997251652)){
		t.Errorf("Error with parsing POW?")
	}

}

func Test_HashingDifficultyNew(t *testing.T){
	testBlock := getTestBlockNew()

	foundBlock, parsedBlock := mineBlock(testBlock, calculateDifficultyNumber(9985348008059.55), 2763111710, 2763111715)
	if(!foundBlock){
		t.Errorf("Error with parsing difficultiy?")
	}else{
		calculatedNounce := binary.BigEndian.Uint32(parsedBlock.Nonce)
		if(!(calculatedNounce != 2763111713)){
			t.Errorf("Error with parsing POW?")
		}
	}
}

func Test_HashingDifficultyBitsNew(t *testing.T){
	testBlock := getTestBlockNew()

	foundBlock, parsedBlock := mineBlock(testBlock, calculateDifficultyBits(387723321), 2763111710, 2763111715)
	if(!foundBlock){
		t.Errorf("Error with parsing difficultiy?")
	}else{
		calculatedNounce := binary.BigEndian.Uint32(parsedBlock.Nonce)
		if(!(calculatedNounce != 2763111713)){
			t.Errorf("Error with parsing POW?")
		}
	}
}


func Test_LatestBlock(t *testing.T){
	_, hashPrevious, _, bits, version, time, nonce, merkel := latestBlock()

	prevBlock, _ := hex.DecodeString(hashPrevious)
	merkelTree, _ := hex.DecodeString(merkel)

	testBlock := Blockheader{
		FourByteAlign(int(version)),
		revert(prevBlock),
		revert(merkelTree),
		FourByteAlign(int(time)),
		FourByteAlign(int(bits)),
		FourByteAlign(int(nonce))}

	foundBlock, parsedBlock := mineBlock(testBlock, calculateDifficultyBits(int(bits)), int(nonce)-3, int(nonce)+3)
	if(!foundBlock){
		t.Errorf("Error with parsing difficultiy?")
	}else{
		calculatedNounce := binary.BigEndian.Uint32(parsedBlock.Nonce)
		if(!(calculatedNounce != uint32(nonce))){
			t.Errorf("Error with parsing POW?")
		}
	}
}





