package controller

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"strings"
	"time"
	pb "user_management_golang/protoc/user_service"
)

type GrpcController struct {
	pb.UserServiceServer // extend pb.GrpcController
}

func NewGrpcController() *GrpcController {
	return &GrpcController{}
}

func (s *GrpcController) Register(ctx context.Context, req *pb.RegisterReq) (*pb.Stdout, error) {
	// 在这里实现用户注册逻辑
	// 从 req 中获取注册请求的信息，处理注册逻辑
	// 创建并返回 RegisterResponse
	response := &pb.Stdout{
		Message: "OK.",
	}
	return response, nil
}

func (s *GrpcController) Login(ctx context.Context, req *pb.LoginReq) (*pb.Stdout, error) {
	// 实现登录逻辑

	return nil, nil
}

func (s *GrpcController) Logout(ctx context.Context, req *pb.NoneReq) (*pb.Stdout, error) {
	// 实现注销逻辑
	return nil, nil
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 在调用 gRPC 服务之前记录请求信息
	parts := strings.Split(info.FullMethod, "/")
	service := parts[len(parts)-2]
	method := parts[len(parts)-1]
	// 调用实际的 gRPC 处理程序
	resp, err := handler(ctx, req)
	// 在调用 gRPC 服务之后记录响应信息
	fmt.Printf("[gRPC] %service \"%s service\" | Status: 200", time.Now().Format("2006/01/02 - 15:04:05"), fmt.Sprintf("/%service/%service", service, method))
	return resp, err
}
