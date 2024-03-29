
package tx

import (
	"testing"
)

func Test_Script(t *testing.T) {
	//	Just a basic Bitcoin script from Andreas Antonopoulos book.
	scripts := createScript()
	scripts.PushOP(OP_2)
	scripts.PushOP(OP_7)
	scripts.PushOP(OP_ADD)
	scripts.PushOP(OP_3)
	scripts.PushOP(OP_SUB)
	scripts.PushOP(OP_1)	
	scripts.PushOP(OP_ADD)
	scripts.PushOP(OP_7)
	scripts.PushOP(OP_EQUAL)
	scripts.execute()

	if(!(scripts.valid())){
		t.Errorf("something is wrong with script execution")
	}

	if(!UnlockScript([]int{OP_3, OP_ADD, OP_5, OP_EQUAL}, []int{OP_2})){
		t.Errorf("something is wrong with UnlockScript execution")		
	}

	//	Peter Todd - sha1 Pinata
	//	https://bitcointalk.org/index.php?topic=293382.0
	//	actual tx https://www.blockchain.com/btc/tx/8d31992805518fd62daa3bdd2a5c4fd2cd3054c9b3dca1d78055e9528cff6adc?show_adv=true
	opcodePinataFix := []int{OP_PUSHDATA2, 0x40, 0x01, 0x25, 0x50, 0x44, 0x46, 0x2d, 0x31, 0x2e, 0x33, 0x0a, 0x25, 0xe2, 0xe3, 0xcf, 0xd3, 0x0a, 0x0a, 0x0a, 0x31, 0x20, 0x30, 0x20, 0x6f, 0x62, 0x6a, 0x0a, 0x3c, 0x3c, 0x2f, 0x57, 0x69, 0x64, 0x74, 0x68, 0x20, 0x32, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x20, 0x33, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x54, 0x79, 0x70, 0x65, 0x20, 0x34, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x20, 0x35, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x20, 0x36, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x20, 0x37, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x38, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x42, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x20, 0x38, 0x3e, 0x3e, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x0a, 0xff, 0xd8, 0xff, 0xfe, 0x00, 0x24, 0x53, 0x48, 0x41, 0x2d, 0x31, 0x20, 0x69, 0x73, 0x20, 0x64, 0x65, 0x61, 0x64, 0x21, 0x21, 0x21, 0x21, 0x21, 0x85, 0x2f, 0xec, 0x09, 0x23, 0x39, 0x75, 0x9c, 0x39, 0xb1, 0xa1, 0xc6, 0x3c, 0x4c, 0x97, 0xe1, 0xff, 0xfe, 0x01, 0x7f, 0x46, 0xdc, 0x93, 0xa6, 0xb6, 0x7e, 0x01, 0x3b, 0x02, 0x9a, 0xaa, 0x1d, 0xb2, 0x56, 0x0b, 0x45, 0xca, 0x67, 0xd6, 0x88, 0xc7, 0xf8, 0x4b, 0x8c, 0x4c, 0x79, 0x1f, 0xe0, 0x2b, 0x3d, 0xf6, 0x14, 0xf8, 0x6d, 0xb1, 0x69, 0x09, 0x01, 0xc5, 0x6b, 0x45, 0xc1, 0x53, 0x0a, 0xfe, 0xdf, 0xb7, 0x60, 0x38, 0xe9, 0x72, 0x72, 0x2f, 0xe7, 0xad, 0x72, 0x8f, 0x0e, 0x49, 0x04, 0xe0, 0x46, 0xc2, 0x30, 0x57, 0x0f, 0xe9, 0xd4, 0x13, 0x98, 0xab, 0xe1, 0x2e, 0xf5, 0xbc, 0x94, 0x2b, 0xe3, 0x35, 0x42, 0xa4, 0x80, 0x2d, 0x98, 0xb5, 0xd7, 0x0f, 0x2a, 0x33, 0x2e, 0xc3, 0x7f, 0xac, 0x35, 0x14, 0xe7, 0x4d, 0xdc, 0x0f, 0x2c, 0xc1, 0xa8, 0x74, 0xcd, 0x0c, 0x78, 0x30, 0x5a, 0x21, 0x56, 0x64, 0x61, 0x30, 0x97, 0x89, 0x60, 0x6b, 0xd0, 0xbf, 0x3f, 0x98, 0xcd, 0xa8, 0x04, 0x46, 0x29, 0xa1, OP_PUSHDATA2, 0x40, 0x01, 0x25, 0x50, 0x44, 0x46, 0x2d, 0x31, 0x2e, 0x33, 0x0a, 0x25, 0xe2, 0xe3, 0xcf, 0xd3, 0x0a, 0x0a, 0x0a, 0x31, 0x20, 0x30, 0x20, 0x6f, 0x62, 0x6a, 0x0a, 0x3c, 0x3c, 0x2f, 0x57, 0x69, 0x64, 0x74, 0x68, 0x20, 0x32, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x20, 0x33, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x54, 0x79, 0x70, 0x65, 0x20, 0x34, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x20, 0x35, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x20, 0x36, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x20, 0x37, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x38, 0x20, 0x30, 0x20, 0x52, 0x2f, 0x42, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x20, 0x38, 0x3e, 0x3e, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x0a, 0xff, 0xd8, 0xff, 0xfe, 0x00, 0x24, 0x53, 0x48, 0x41, 0x2d, 0x31, 0x20, 0x69, 0x73, 0x20, 0x64, 0x65, 0x61, 0x64, 0x21, 0x21, 0x21, 0x21, 0x21, 0x85, 0x2f, 0xec, 0x09, 0x23, 0x39, 0x75, 0x9c, 0x39, 0xb1, 0xa1, 0xc6, 0x3c, 0x4c, 0x97, 0xe1, 0xff, 0xfe, 0x01, 0x73, 0x46, 0xdc, 0x91, 0x66, 0xb6, 0x7e, 0x11, 0x8f, 0x02, 0x9a, 0xb6, 0x21, 0xb2, 0x56, 0x0f, 0xf9, 0xca, 0x67, 0xcc, 0xa8, 0xc7, 0xf8, 0x5b, 0xa8, 0x4c, 0x79, 0x03, 0x0c, 0x2b, 0x3d, 0xe2, 0x18, 0xf8, 0x6d, 0xb3, 0xa9, 0x09, 0x01, 0xd5, 0xdf, 0x45, 0xc1, 0x4f, 0x26, 0xfe, 0xdf, 0xb3, 0xdc, 0x38, 0xe9, 0x6a, 0xc2, 0x2f, 0xe7, 0xbd, 0x72, 0x8f, 0x0e, 0x45, 0xbc, 0xe0, 0x46, 0xd2, 0x3c, 0x57, 0x0f, 0xeb, 0x14, 0x13, 0x98, 0xbb, 0x55, 0x2e, 0xf5, 0xa0, 0xa8, 0x2b, 0xe3, 0x31, 0xfe, 0xa4, 0x80, 0x37, 0xb8, 0xb5, 0xd7, 0x1f, 0x0e, 0x33, 0x2e, 0xdf, 0x93, 0xac, 0x35, 0x00, 0xeb, 0x4d, 0xdc, 0x0d, 0xec, 0xc1, 0xa8, 0x64, 0x79, 0x0c, 0x78, 0x2c, 0x76, 0x21, 0x56, 0x60, 0xdd, 0x30, 0x97, 0x91, 0xd0, 0x6b, 0xd0, 0xaf, 0x3f, 0x98, 0xcd, 0xa4, 0xbc, 0x46, 0x29, 0xb1}//, 0x08, 0x6e, 0x87, 0x91, 0x69, 0xa7, 0x7c, 0xa7, 0x87}
	opcodePinata := []int{OP_2DUP, OP_EQUAL, OP_NOT, OP_VERIFY, OP_SHA1, OP_SWAP, OP_SHA1, OP_EQUAL}

	if(!UnlockScript(opcodePinata, opcodePinataFix )){
		t.Errorf("something is wrong with UnlockScript execution")		
	}
}