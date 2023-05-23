package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
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
	name := "/home/bot/Downloads/QQ_Image_1684497413039.jpg"
	fileContent, _ := os.ReadFile(name)
	content, _ := GetResult("QQ_Image_1684497413039.jpg", fileContent)
	os.WriteFile("test.jpg", content.Arrowimage, 0644)
}
