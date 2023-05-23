package client

import (
	"ai-medical/pb"
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func NewConnection() (*grpc.ClientConn, error) {
	addr := "124.221.136.185:50001"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetResult(filename string, content []byte) (*pb.Feature, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	conn, err := NewConnection()
	if err != nil {
		return nil, err
	}
	c := pb.NewLandmarkClient(conn)
	feat, err := c.GetResult(ctx, &pb.Point{
		Name:  filename,
		Image: content,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return feat, err
}
