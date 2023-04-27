package main

import (
	"github.com/ilyasoloviev99/category-service/internal/app/category_service"
	desc "github.com/ilyasoloviev99/category-service/pkg/category"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	categoryService := category_service.NewCategoryService()

	desc.RegisterCategoryServiceServer(grpcServer, categoryService)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
