// Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"

	pb "github.com/AccelByte/accelbyte-gdpr-sdk/pkg/pb"
)

type GRPCServer struct {
	pb.UnimplementedGDPRServer
}

func NewGDPRServiceServer() *GRPCServer {
	rand.Seed(time.Now().Unix())

	return &GRPCServer{}
}

func (s *GRPCServer) PersonalDataGeneration(_ context.Context, req *pb.PersonalDataRequest) (*pb.PersonalDataResponse, error) {
	logrus.Info("Invoke PersonalDataGeneration")
	return nil, nil
}

func (s *GRPCServer) DataDeletion(_ context.Context, req *pb.DataDeletionRequest) (*pb.DataDeletionResponse, error) {
	logrus.Info("Invoke DataDeletion")
	return nil, nil
}
