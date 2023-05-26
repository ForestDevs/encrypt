package main

import (
	"context"
	"encrypt-decrypt/config"
	"encrypt-decrypt/internal/pb"
	"encrypt-decrypt/internal/service"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

func serveHTTP(cfg *config.AppConfig) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	appPort := fmt.Sprintf(":%d", cfg.AppPort)
	appProxyPort := fmt.Sprintf(":%d", cfg.AppProxyPort)

	err := pb.RegisterLivenessProbeHandlerFromEndpoint(ctx, mux, appPort, opts)
	if err != nil {
		fmt.Println(err)
	}

	err = pb.RegisterEncryptServiceHandlerFromEndpoint(ctx, mux, appPort, opts)
	if err != nil {
		fmt.Println(err)
	}

	err = pb.RegisterDecryptServiceHandlerFromEndpoint(ctx, mux, appPort, opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("server http listening at %v\n", appProxyPort)

	if err := http.ListenAndServe(appProxyPort, mux); err != nil {
		fmt.Println(err)
	}
}

func serveGRPC(cfg *config.AppConfig) {
	appPort := fmt.Sprintf(":%d", cfg.AppPort)
	lis, err := net.Listen("tcp", appPort)
	server := grpc.NewServer()
	if err != nil {
		fmt.Println(err)
	}

	livenessProbeService := service.NewLivenessProbeService()
	pb.RegisterLivenessProbeServer(server, livenessProbeService)

	encryptService := service.NewEncrypt(cfg.CryptoKey)
	pb.RegisterEncryptServiceServer(server, encryptService)

	decryptService := service.NewDecrypt(cfg.CryptoKey)
	pb.RegisterDecryptServiceServer(server, decryptService)

	fmt.Printf("server grpc listening at %v\n", lis.Addr())

	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	cfg, err := config.NewAppConfig()
	if err != nil {
		panic(err)
	}
	go serveGRPC(cfg)
	serveHTTP(cfg)
}
