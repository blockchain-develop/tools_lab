package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tools_lab/graphql-bitquery/graph/generated"
	"github.com/tools_lab/graphql-bitquery/graph/model"
)

func (r *filecoinBlockResolver) WinCount(ctx context.Context, obj *model.FilecoinBlock, date *model.DateSelector, time *model.DateTimeSelector, height *model.BlockSelector, blockHash *model.HashSelector, miner []*model.AddressSelector, blockIndex *model.IntegerSelector) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *filecoinBlockResolver) Wincount(ctx context.Context, obj *model.FilecoinBlock) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Algorand(ctx context.Context, network *model.AlgorandNetwork) (*model.Algorand, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Binance(ctx context.Context) (*model.Binance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Bitcoin(ctx context.Context, network *model.BitcoinNetwork) (*model.Bitcoin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cardano(ctx context.Context, network *model.CardanoNetwork) (*model.Cardano, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Conflux(ctx context.Context, network *model.ConfluxNetwork) (*model.Conflux, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Diem(ctx context.Context, network *model.DiemNetwork) (*model.Libra, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Elrond(ctx context.Context, network *model.ElrondNetwork) (*model.Elrond, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Eos(ctx context.Context, network *model.EosNetwork) (*model.Eos, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ethereum(ctx context.Context, network *model.EthereumNetwork) (*model.Ethereum, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ethereum2(ctx context.Context, network *model.Ethereum2Network) (*model.Ethereum2, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Filecoin(ctx context.Context, network *model.FilecoinNetwork) (*model.Filecoin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Harmony(ctx context.Context, network *model.HarmonyNetwork) (*model.Harmony, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hedera(ctx context.Context, network *model.HederaNetwork) (*model.Hedera, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Offchain(ctx context.Context) (*model.Offchain, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ripple(ctx context.Context, network *model.RippleNetwork) (*model.Ripple, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Search(ctx context.Context, string string, limit *int, offset *int, network *model.Network) ([]*model.Result, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Solana(ctx context.Context, network *model.SolanaNetwork) (*model.Solana, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Stellar(ctx context.Context, network *model.StellarNetwork) (*model.Stellar, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tron(ctx context.Context, network *model.TronNetwork) (*model.Tron, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Ethereum(ctx context.Context, network *model.EthereumNetwork) (<-chan *model.Ethereum, error) {
	panic(fmt.Errorf("not implemented"))
}

// FilecoinBlock returns generated.FilecoinBlockResolver implementation.
func (r *Resolver) FilecoinBlock() generated.FilecoinBlockResolver { return &filecoinBlockResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type filecoinBlockResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
