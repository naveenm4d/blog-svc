package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/naveenm4d/blog-svc/internal/app/repositories"
	"github.com/naveenm4d/blog-svc/internal/app/services"
	"github.com/naveenm4d/blog-svc/internal/config"
	"github.com/naveenm4d/blog-svc/internal/handlers"
	"github.com/naveenm4d/blog-svc/proto"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/naveenm4d/pkg/mongodb"
)

var (
	grpcServer *grpc.Server
)

func main() {
	appName := "blogs-svc"

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-c
		fmt.Printf("Got %s signal. Cancelling", sig)

		cancel()

		if grpcServer != nil {
			grpcServer.GracefulStop()
		}
	}()

	fmt.Println("Initializing mongo connection...")

	mongoClient, mErr := mongodb.CreateMongoClient(ctx, *config.Config.MongoDBEndpoint)

	if mErr != nil {
		panic(mErr)
	}

	defer mongoClient.Disconnect(ctx)

	fmt.Println("Initializing repository...")

	repository := repositories.NewPostsRepository(
		mongoClient.
			Database(*config.Config.MongoDBDatabase).
			Collection(*config.Config.MongoPostsCollection))

	fmt.Println("Initializing service...")

	postsService := services.NewPostService(repository)

	fmt.Println("Initializing handler...")

	handler := handlers.NewHandler(postsService)

	lis, lErr := net.Listen("tcp", fmt.Sprintf(":%s", *config.Config.GrpcPort))
	if lErr != nil {
		panic(fmt.Sprintf("failed to listen: %v", lErr))
	}

	fmt.Println("Initializing grpc server...")

	grpcServer = grpc.NewServer()

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("blog-svc", healthpb.HealthCheckResponse_SERVING)

	proto.RegisterBlogSvcServer(grpcServer, handler)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}

	<-ctx.Done()

	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	fmt.Print(appName)
}
