package gdprsdk

import (
	pb "github.com/marselsampe/accelbyte-gdpr-sdk/pkg/pb"
	"github.com/marselsampe/accelbyte-gdpr-sdk/pkg/service"
	"google.golang.org/grpc"
)

type GdprSDK struct {
	gdprServiceServer *service.GDPRServiceServer
}

func NewGdprSDK() *GdprSDK {
	return &GdprSDK{
		gdprServiceServer: service.NewGDPRServiceServer(),
	}
}

func (sdk GdprSDK) RegisterGRPC(server *grpc.Server) {
	pb.RegisterGDPRServer(server, sdk.gdprServiceServer)
}
