package gdprsdk

import (
	pb "github.com/marselsampe/accelbyte-gdpr-sdk/pkg/pb"
	"github.com/marselsampe/accelbyte-gdpr-sdk/pkg/service"
	"google.golang.org/grpc"
)

type GdprSDK struct {
	grpcServer *service.GRPCServer
}

func NewGdprSDK() *GdprSDK {
	return &GdprSDK{
		grpcServer: &service.GRPCServer{},
	}
}

func (sdk GdprSDK) RegisterGRPC(grpcServer *grpc.Server) {
	pb.RegisterGDPRServer(grpcServer, sdk.grpcServer)
}
