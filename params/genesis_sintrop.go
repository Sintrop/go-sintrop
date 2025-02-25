package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params/types/genesisT"
)

var SintropGenesisHash = common.HexToHash("0xb390a5cdb5c39d4ea8909b6822573b50630a1254ba352142e436a76e99702e32")

func DefaultSintropGenesisBlock() *genesisT.Genesis {
	return &genesisT.Genesis{
		Config:     SintropChainConfig,
		Nonce:      hexutil.MustDecodeUint64("0x0"),
		ExtraData:  hexutil.MustDecode("0x42"),
		GasLimit:   hexutil.MustDecodeUint64("0x2fefd8"),
		Difficulty: hexutil.MustDecodeBig("0x20000000"),
		Timestamp:  1740492658,
		Alloc: genesisT.GenesisAlloc{
			common.HexToAddress("366ae7da62294427c764870bd2a460d7ded29d30"): genesisT.GenesisAccount{
				Balance: big.NewInt(42),
			},
		},
	}
}
