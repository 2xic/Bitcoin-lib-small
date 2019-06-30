# Bitcoin

[![Build Status](https://travis-ci.org/2xic/bip-39.svg?branch=master)](https://travis-ci.org/2xic/bip-39)
[![Coverage Status](https://coveralls.io/repos/github/2xic/bip-39/badge.svg?branch=master)](https://coveralls.io/github/2xic/bip-39?branch=master)

Just visiting different parts of the Bitcoin stack, *not* writing a complete Bitcoin library. Look at [btcsuite](https://github.com/btcsuite) if you are looking for production software. This repo is mostly for me to learn, *not* to write production software. Can't experiment with elliptic curves if I write production software ;) (don't roll your own crypto)

# Features 
*Not far from an spv client!*
-	[BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki)
	-	"Mnemonic code for generating deterministic keys"
-	[BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
	-	"Hierarchical deterministic wallets"
		-	not fully implemented abillity to generate child key(s), need to write some more crypto code first :=)
-	[Part of the Bitcoin network protocol](https://en.bitcoin.it/wiki/Protocol_documentation)
	-	Conenct to a node and request block(s)
-	tx
	-	you can contruct a transaction
	-	implemented a subset of the Bitcoin script opcodes
		-	Featuring: [Peter Todd - sha1 Pinata](https://bitcointalk.org/index.php?topic=293382.0)

# Credit
- [jimmysong](https://github.com/jimmysong/programmingbitcoin) - the code releated to elliptic curve(secp256k1) is based off content from his programming bitcoin book. 

***No warranty is given. No complaints will be answered.***