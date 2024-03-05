package rpc

import (
	"context"
	"github.com/kpyramid/bitcoin-inscribe/proto"
)

type Inscribe struct {
	proto.UnimplementedInscribeServiceServer
}

func (Inscribe) GenerateInscribeNFTParams(context.Context, *proto.GenerateInscribeNFTParamsRequest) (*proto.GenerateInscribeNFTParamsResponse, error) {
	panic("Todo")
}

func (Inscribe) LaunchInscribe(context.Context, *proto.LaunchInscribeRequest) (*proto.LaunchInscribeResponse, error) {
	panic("Todo")
}

func (Inscribe) GetInscribeInfo(context.Context, *proto.GetInscribeInfoRequest) (*proto.GetInscribeInfoResponse, error) {
	panic("Todo")
}
