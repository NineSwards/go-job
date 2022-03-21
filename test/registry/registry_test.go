package registry

import (
	"fmt"
	pb "github.com/NineSwards/go-job/core/server/registry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestRegistryServer(t *testing.T) {
	client, err := grpc.Dial(":3344", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(client *grpc.ClientConn) {
		err := client.Close()
		if err != nil {
			t.Log("close error")
		}
	}(client)
	// 初始化grpc客户端
	findClassClient := pb.NewRegisterClient(client)
	class, err := findClassClient.Registry(context.Background(), &pb.RegistryRequest{
		Name: "root",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(class.Message) // Chinese
}
