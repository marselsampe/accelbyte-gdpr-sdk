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
	"github.com/marselsampe/accelbyte-gdpr-sdk/pkg/object"
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

func (sdk GdprSDK) SetDataGenerationHandler(handler object.DataGenerationHandler) {
	sdk.gdprServiceServer.DataGenerationHandler = handler
}

func (sdk GdprSDK) SetDataDeletionHandler(handler object.DataDeletionHandler) {
	sdk.gdprServiceServer.DataDeletionHandler = handler
}

func (sdk GdprSDK) RegisterGRPC(server *grpc.Server) {
	pb.RegisterGDPRServer(server, sdk.gdprServiceServer)
}
