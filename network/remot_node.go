package network


import (
	"net"
	"fmt"
	"github.com/2xic/bip-39/secp256k1"
	"encoding/binary"
	"time"
	"math/rand"
	"strings"
	"encoding/hex"
	"bytes"

)

func BytePadding(input int, size int) []byte{
	output := make([]byte, size)
	binary.BigEndian.PutUint32(output, uint32(input))
	return output
}

func BytePaddingL(input int, size int) []byte{
	output := make([]byte, size)
	binary.LittleEndian.PutUint32(output, uint32(input))
	return output
}

func SlicePadding(input []byte, size int) []byte{
	padding := (size - len(input))
	if(0 < padding){
		input = append(input, make([]byte, padding)...)
	}
	return input
}

func connect(target string) *net.Conn{
	conn, err := net.Dial("tcp", target)
	if err != nil {
		return nil;
	}
	return &conn
}

func formatMessage(message string, payload []byte) []byte{
	EncodedMessage := make([]byte, 0)
	//	magic bytes
	EncodedMessage = append(EncodedMessage, []byte{0xF9, 0xBE, 0xB4, 0xD9}...)
	EncodedMessage = append(EncodedMessage, SlicePadding([]byte(message), 12)...)

	EncodedMessage = append(EncodedMessage, BytePaddingL(len(payload), 4)...)
	EncodedMessage = append(EncodedMessage, secp256k1.ReturnCheckSum(payload)...)
	EncodedMessage = append(EncodedMessage, payload...)
	
	return EncodedMessage
}

func formatMessageOnly(message string) []byte{
	return formatMessage(message, make([]byte, 0))
}

func handshakePayload()[]byte{
	version := BytePaddingL(180002, 4)
	services := BytePadding(0, 8)
	timestamp := BytePaddingL(int(time.Now().UnixNano() / 1000000000), 8)

	addressMe := make([]byte, 8)
	addressMe = append(addressMe, SlicePadding([]byte("127.0.0.1"), 16)...)
	addressMe = append(addressMe, BytePadding(8333, 4)[2:]...)

	addressYou := make([]byte, 8)
	addressYou = append(addressYou, SlicePadding([]byte("127.0.0.1"), 16)...)
	addressYou = append(addressYou, BytePadding(8333, 4)[2:]...)

	nonce := make([]byte, 8)
	rand.Read(nonce)

	userAgent := make([]byte, 1)
	height := make([]byte, 4)

	payload := make([]byte, 0)
	payload = append(payload, version...)
	payload = append(payload, services...)
	payload = append(payload, timestamp...)
	payload = append(payload, addressMe...)
	payload = append(payload, addressYou...)

	payload = append(payload, nonce...)
	payload = append(payload, userAgent...)
	payload = append(payload, height...)

	return payload
}

func ByteReverse(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	return append(ByteReverse(input[1:]), input[0]) 
}

func getBlock(tx string) []byte{
	payload :=  make([]byte, 0)
	test, _ := hex.DecodeString(tx)
	payload = append(payload, byte(1))
	payload = append(payload, BytePaddingL(2, 4)...)
	payload = append(payload, ByteReverse(test)...)
	return payload
}

func recieveSocketLength(size int, connection net.Conn) []byte{
	response := make([]byte, size)
	_, _ = connection.Read(response)
	return response
}


func Byte2Int(data []byte) (ret int32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return 
}

func parseResponse(connection net.Conn) (command []byte, response[]byte){
	response = make([]byte, 24)
	_, err := connection.Read(response)
	if(err != nil){
		panic(err)
	}

//	magic := response[:4]
	command = response[4:16]
	lengthPayload := int(Byte2Int(response[16:20]))
//	checksum := response[20:]
	payload := recieveSocketLength(lengthPayload, connection)

	return command, payload
}


func handshakeNode(connection net.Conn){
	_, err := connection.Write(formatMessage("version", handshakePayload()))
	if(err != nil){
		panic(err)
	}
	//	version
	command, payload :=	parseResponse(connection)
	fmt.Println(string(command))
	fmt.Println(string(payload))

	//	verack
	command, payload =	parseResponse(connection)
	fmt.Println(string(command))
	fmt.Println(string(payload))

	//	acknowledgement
	_, err = connection.Write(formatMessageOnly("verack"))
	if(err != nil){
		panic(err)
	}
}

func Getblock() *[]byte{
	addrs, err := net.LookupHost("dnsseed.bluematt.me")
	if(err != nil){
		panic(err)		
	}
	target := addrs[4] + ":8333"
	nodeCon := connect(target)

	handshakeNode(*nodeCon)

	_, err = (*nodeCon).Write(formatMessage("getdata", getBlock("0000000000000000000541989b34cfd2e2acf776d6433ed797821734163031d9")))
	if(err != nil){
		panic(err)
	}

	command, payload := parseResponse(*nodeCon)

	i := 0
	maxTry := 100
	for !strings.HasPrefix(string(command), "block"){
		command, payload = parseResponse(*nodeCon)
		fmt.Println(string(command))
		if(maxTry < i){
			//return 
			break
		}
		i += 1
	}
	if(maxTry < i){
		return nil
	}
	return &payload
}
