package test

import (
	"fmt"
	executorReg "github.com/NineSwards/go-job/server/executor"
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
	// close
	defer func(client *grpc.ClientConn) {
		err := client.Close()
		if err != nil {
			t.Log("close error")
		}
	}(client)

	// 初始化grpc客户端
	jobClient := executorReg.NewJobExecutorRegisterClient(client)
	class, err := jobClient.ExecutorRegistry(context.Background(), &executorReg.RegistryRequest{
		ServerName: "Test",
		ServerList: []string{"1", "2", "3"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(class.Message) // Chinese
}

func TestUnRegistryServer(t *testing.T) {
	client, err := grpc.Dial(":3344", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	// close
	defer func(client *grpc.ClientConn) {
		err := client.Close()
		if err != nil {
			t.Log("close error")
		}
	}(client)

	// 初始化grpc客户端
	jobClient := executorReg.NewJobExecutorRegisterClient(client)
	class, err := jobClient.ExecutorUnRegistry(context.Background(), &executorReg.RegistryRequest{
		ServerName: "Test",
		ServerList: []string{"1", "2", "3"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(class.Message) // Chinese
}
