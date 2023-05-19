package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestConnection(t *testing.T) {
	addr := "124.221.136.185:50001"
	_, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error(err)
	}
}

func TestGetResult(t *testing.T) {
	GetResult()
}
