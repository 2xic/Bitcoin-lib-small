package base58

import(
	"math/big"
)

func ConvertString(inputString string) string{
	stringBytes := []byte(inputString)
	return ConvertBytes(stringBytes)
}

func ConvertBytes(inputBytes []byte) (output string){
	alphabet := string("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

	prefix := ""
	for i:= 0; i < len(inputBytes); i++{
		if(inputBytes[i] == 0){
			prefix += "1"
		}else{
			break;
		}
	}

	byteNumber := new(big.Int)
	byteNumber.SetBytes(inputBytes)

	for byteNumber.Cmp(big.NewInt(0)) == 1{
		alphabetIndex := new(big.Int)
		alphabetIndex.Mod(byteNumber, big.NewInt(58))
		byteNumber.Div(byteNumber, big.NewInt(58))

		output = string(alphabet[alphabetIndex.Int64()]) + output
	}
	return prefix + output
}



