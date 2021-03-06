package substrate

import (
	"context"
	"fmt"

	"github.com/mailchain/go-substrate-rpc-client/types"
	"github.com/mailchain/mailchain/cmd/indexer/internal/actions"
	"github.com/pkg/errors"
)

type Block struct {
	txProcessor actions.Transaction
}

type TxOptions struct {
	Block *types.Block
}

func NewBlockProcessor(tx actions.Transaction) *Block {
	return &Block{
		txProcessor: tx,
	}
}

func (b *Block) Run(ctx context.Context, protocol, network string, blk interface{}) error {
	sbrtBlk, ok := blk.(*types.Block)
	if !ok {
		return errors.Errorf("tx %T must be go-substrate-rpc-client/types.Block", blk)
	}

	fmt.Println("block hash: ", sbrtBlk.Header.ParentHash.Hex())

	txs := sbrtBlk.Extrinsics
	for i := range txs {
		if !txs[i].Signature.Signer.IsAccountID {
			continue
		}

		if err := b.txProcessor.Run(ctx, protocol, network, &txs[i], &TxOptions{Block: sbrtBlk}); err != nil {
			return errors.Wrapf(err, "fails to process transaction: block=%s, transaction-index=%d", sbrtBlk.Header.ParentHash.Hex(), i)
		}
	}

	return nil
}
