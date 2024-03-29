# Small Bitcoin lib

[![Build Status](https://travis-ci.org/2xic/Bitcoin-lib-small.svg?branch=master)](https://travis-ci.org/2xic/Bitcoin-lib-small)
[![Coverage Status](https://coveralls.io/repos/github/2xic/Bitcoin-lib-small/badge.svg?branch=master)](https://coveralls.io/github/2xic/Bitcoin-lib-small?branch=master)
[![Internet cash](https://img.shields.io/badge/BTC-33o9PdY67drhNARg5YxaK6zs85F124qooA-blue)](https://img.shields.io/badge/BTC-33o9PdY67drhNARg5YxaK6zs85F124qooA-blue)

Visiting different parts of the Bitcoin stack, *not* writing a complete Bitcoin library. I won't spend time writing code that I don't learn from. For instance I won't learn that much more by implementing the entire network protocol, so I will rather explore another part of the stack. Look at [btcsuite](https://github.com/btcsuite) if you are looking for production software. This repo is mostly for me to learn, *not* to write production software. Can't experiment with elliptic curves if I write production software ;) (don't roll your own crypto)

# Features 
*Put the pieces together and you have a basic full node*
-	[BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki)
	-	"Mnemonic code for generating deterministic keys"
-	[BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
	-	"Hierarchical deterministic wallets"
-	[Part of the Bitcoin network protocol](https://en.bitcoin.it/wiki/Protocol_documentation)
	-	Conenct to a node and request block(s)
-	tx
	-	contruct a transaction
	-	implemented a subset of the Bitcoin script opcodes
		-	Featuring: [Peter Todd - sha1 Pinata](https://bitcointalk.org/index.php?topic=293382.0)
-	mining
	-	make and verify blocks
		-	if you add support for inventory vectors in the network protocol you can also do lottery mining!

# Credit
- [jimmysong](https://github.com/jimmysong/programmingbitcoin) - the code releated to elliptic curve(secp256k1) is based off content from his programming bitcoin book. 
- [tyler-smith](https://github.com/tyler-smith/go-bip32) - code releated to expanding a compressed public key

***No warranty is given. No complaints will be answered.***