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

package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"

	"github.com/marselsampe/accelbyte-gdpr-sdk/pkg/constant"
	pb "github.com/marselsampe/accelbyte-gdpr-sdk/pkg/pb"
)

type GDPRServiceServer struct {
	pb.UnimplementedGDPRServer
	DataGenerationHandler constant.DataGenerationHandler
	DataDeletionHandler   constant.DataDeletionHandler
}

func NewGDPRServiceServer() *GDPRServiceServer {
	rand.Seed(time.Now().Unix())

	return &GDPRServiceServer{}
}

func (s *GDPRServiceServer) DataGeneration(_ context.Context, req *pb.DataGenerationRequest) (*pb.DataGenerationResponse, error) {
	logrus.Info("Invoke DataGeneration")
	if s.DataGenerationHandler != nil {
		err := s.DataGenerationHandler()
		if err != nil {
			// TODO: handle error
		}
	}
	return &pb.DataGenerationResponse{}, nil
}

func (s *GDPRServiceServer) DataDeletion(_ context.Context, req *pb.DataDeletionRequest) (*pb.DataDeletionResponse, error) {
	logrus.Info("Invoke DataDeletion")
	if s.DataDeletionHandler != nil {
		err := s.DataDeletionHandler()
		if err != nil {
			// TODO: handle error
		}
	}
	return &pb.DataDeletionResponse{}, nil
}
