package params

import (
	"testing"

)

// TestGenesisHashSequoia tests that SequoiaGenesisHash is the correct value for the genesis configuration.
func TestGenesisHashSequoia(t *testing.T) {
	genesis := DefaultSequoiaGenesisBlock()
	block := genesisToBlock(genesis, nil)
	if block.Hash() != SequoiaGenesisHash {
		t.Errorf("want: %s, got: %s", SequoiaGenesisHash.Hex(), block.Hash().Hex())
	}
}

