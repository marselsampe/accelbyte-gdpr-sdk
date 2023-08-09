/*
 * Copyright (c) 2023 AccelByte Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 *
 */

package gdprsdk

import (
	gdprGrpc "github.com/marselsampe/accelbyte-gdpr-sdk/pkg/grpc"
	"github.com/marselsampe/accelbyte-gdpr-sdk/pkg/object"
	pb "github.com/marselsampe/accelbyte-gdpr-sdk/pkg/pb"
	"google.golang.org/grpc"
)

func NewGdprGrpc() *GdprGrpc {
	return &GdprGrpc{
		gdprServiceServer: gdprGrpc.NewGDPRServiceServer(),
	}
}

type GdprGrpc struct {
	gdprServiceServer *gdprGrpc.GDPRServiceServer
}

func (sdk GdprGrpc) RegisterGRPC(server *grpc.Server) {
	pb.RegisterGDPRServer(server, sdk.gdprServiceServer)
}

func (sdk GdprGrpc) SetDataGenerationHandler(handler object.DataGenerationHandler) {
	sdk.gdprServiceServer.DataGenerationHandler = handler
}

func (sdk GdprGrpc) SetDataDeletionHandler(handler object.DataDeletionHandler) {
	sdk.gdprServiceServer.DataDeletionHandler = handler
}
