package client

import (
	"context"
	"fmt"
	"log"
	"time"

	// pb "github.com/souzera/learning-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed in connection: %v", err)
	}
	defer connection.Close()

	client = pb.NewUserServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Criar um usuários
	fmt.Println("=== Criando usuários ===")
	createSenna, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "Ayrton Senna",
		Email: "senna@example.com",
		Age:   29,
	})

	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Printf("Usuário criado: %+v\n", createSenna.User)
	fmt.Printf("Mensagem: %s\n\n", createSenna.Message)

	createProst, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "Alain Prost",
		Email: "prost@example.com",
		Age:   34,
	})

	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Printf("Usuário criado: %+v\n", createProst.User)
	fmt.Printf("Mensagem: %s\n\n", createProst.Message)

	// Buscar usuário por ID
	fmt.Println("=== Buscando usuário por ID ===")
	getResp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Erro ao buscar usuário: %v", err)
	}
	fmt.Printf("Usuário encontrado: %+v\n\n", getResp)

	// Listar todos os usuários
	fmt.Println("=== Listando todos os usuários ===")
	listResp, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
	if err != nil {
		log.Fatalf("Erro ao listar usuários: %v", err)
	}

	for i, user := range listResp.Users {
		fmt.Printf("Usuário %d: %+v\n", i+1, user)
	}

}
