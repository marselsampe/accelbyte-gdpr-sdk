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

package grpc

import (
	"context"
	"math/rand"
	"time"

	"github.com/marselsampe/accelbyte-gdpr-sdk/pkg/object"
	pb "github.com/marselsampe/accelbyte-gdpr-sdk/pkg/pb"
	"github.com/marselsampe/accelbyte-gdpr-sdk/pkg/utils"
	"github.com/sirupsen/logrus"
)

type GDPRServiceServer struct {
	pb.UnimplementedGDPRServer
	DataGenerationHandler object.DataGenerationHandler
	DataDeletionHandler   object.DataDeletionHandler
}

func NewGDPRServiceServer() *GDPRServiceServer {
	rand.Seed(time.Now().Unix())

	return &GDPRServiceServer{}
}

func (s *GDPRServiceServer) DataGeneration(ctx context.Context, req *pb.DataGenerationRequest) (*pb.DataGenerationResponse, error) {
	if req.Namespace == "" || req.UserId == "" || req.UploadUrl == "" {
		return &pb.DataGenerationResponse{
			Success: false,
			Message: "required payload is empty",
		}, nil
	}

	namespace := req.Namespace
	userID := req.UserId

	if s.DataGenerationHandler != nil {
		logrus.Errorf("[DataGeneration worker] Failed executing DataGenerationHandler. Error: %s", err)
		resultBytes, err := s.DataGenerationHandler(namespace, userID)
		if err != nil {
			logrus.Errorf("[DataGeneration worker] Failed executing DataGenerationHandler. Error: %s", err)
			return &pb.DataGenerationResponse{
				Success: false,
				Message: err.Error(),
			}, nil
		}
		if resultBytes == nil {
			return &pb.DataGenerationResponse{Success: true}, nil
		}

		// upload file into storage
		err = utils.UploadFile(ctx, req.UploadUrl, resultBytes)
		if err != nil {
			logrus.Errorf("[DataGeneration worker] Failed uploading file. Error: %s", err)
			return &pb.DataGenerationResponse{
				Success: false,
				Message: "Failed uploading file. Error: " + err.Error(),
			}, nil
		}
	}

	return &pb.DataGenerationResponse{Success: true}, nil
}

func (s *GDPRServiceServer) DataDeletion(_ context.Context, req *pb.DataDeletionRequest) (*pb.DataDeletionResponse, error) {
	if req.Namespace == "" || req.UserId == "" {
		return &pb.DataDeletionResponse{
			Success: false,
			Message: "required payload is empty",
		}, nil
	}

	namespace := req.Namespace
	userID := req.UserId

	if s.DataDeletionHandler != nil {
		logrus.Infof("Invoke DataDeletion for namespace [%s] userId [%s]", namespace, userID)
		err := s.DataDeletionHandler(namespace, userID)
		if err != nil {
			logrus.Errorf("[DataGeneration worker] Failed executing DataDeletionHandler. Error: %s", err)
			return &pb.DataDeletionResponse{
				Success: false,
				Message: err.Error(),
			}, nil
		}
	}

	return &pb.DataDeletionResponse{Success: true}, nil
}
