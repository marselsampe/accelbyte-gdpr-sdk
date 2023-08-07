// Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"

	pb "github.com/marselsampe/accelbyte-gdpr-sdk/pkg/pb"
)

type GDPRServiceServer struct {
	pb.UnimplementedGDPRServer
}

func NewGDPRServiceServer() *GDPRServiceServer {
	rand.Seed(time.Now().Unix())

	return &GDPRServiceServer{}
}

func (s *GDPRServiceServer) PersonalDataGeneration(_ context.Context, req *pb.PersonalDataRequest) (*pb.PersonalDataResponse, error) {
	logrus.Info("Invoke PersonalDataGeneration")
	return &pb.PersonalDataResponse{}, nil
}

func (s *GDPRServiceServer) DataDeletion(_ context.Context, req *pb.DataDeletionRequest) (*pb.DataDeletionResponse, error) {
	logrus.Info("Invoke DataDeletion")
	return &pb.DataDeletionResponse{}, nil
}
