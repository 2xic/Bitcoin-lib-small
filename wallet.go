

package main


import(
	"fmt"
	"github.com/2xic/bip-39/bip32"
	"github.com/2xic/bip-39/bip39"
)

func main(){	
	words, _ := bip39.GenerateMnemonic()
	seed := bip39.Memonic2Seed(words, "TREZOR")

	fmt.Println(words)

	MasterPrivateKey := bip32.MasterPrivateKey(seed)
	MasterPublicKey := bip32.MasterPublicKey(MasterPrivateKey)

	fmt.Println(MasterPrivateKey.Readable())
	fmt.Println(MasterPublicKey.Readable())
	fmt.Println(MasterPrivateKey.ReadableAddress())

	fmt.Println("==============" )

	ChildrenPrivate := bip32.ChildKey(MasterPrivateKey, 0)

	fmt.Println(ChildrenPrivate.Readable())
}