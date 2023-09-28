package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"user_management_golang/controller"
	pb "user_management_golang/protoc/user_service"
)

func main() {
	var mode string
	var grpcPort string
	var restPort string

	// 使用 flag 包定义命令行参数
	flag.StringVar(&mode, "mode", "simple", "Select boot mode.")
	flag.StringVar(&grpcPort, "grpc-port", "50051", "gRPC server port")
	flag.StringVar(&restPort, "rest-port", "50050", "RESTful server port")
	flag.Parse()

	// 启动 gRPC 服务器
	go func() {
		fmt.Printf("gRPC server started and listening on :%s port \n", grpcPort)
		listen, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen for gRPC: %v", err)
		}
		s := grpc.NewServer(
			grpc.UnaryInterceptor(controller.LoggingInterceptor),
		)
		pb.RegisterUserServiceServer(s, &controller.GrpcService{})
		if err := s.Serve(listen); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// 启动 RESTful 服务器
	go func() {
		fmt.Printf("RESTful server started and listening on :%s port \n", restPort)
		gin.SetMode(gin.ReleaseMode) // set release mode
		restAPI := controller.NewRestService()
		r := gin.New()
		r.POST("/register", restAPI.Register)
		// 添加其他 RESTful 路由
		if err := r.Run(":" + restPort); err != nil {
			log.Fatalf("Failed to serve RESTful: %v", err)
		}
	}()

	// 防止程序提前退出
	select {}
}
