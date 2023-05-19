package client

import (
	"ai-medical/pb"
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
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

func GetResult() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	conn, err := NewConnection()
	if err != nil {
		return
	}
	c := pb.NewLandmarkClient(conn)
	f, err := os.ReadFile("/home/bot/Downloads/QQ_Image_1684497413039.jpg")
	feat, err := c.GetResult(ctx, &pb.Point{
		Name:  "QQ_Image_1684497413039.jpg",
		Image: f,
	})
	if err != nil {
		log.Error(err)
		return
	}
	println(feat)
}
