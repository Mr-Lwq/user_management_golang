package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"user_management_golang/controller"
	"user_management_golang/dao"
	pb "user_management_golang/protoc/user_service"
	"user_management_golang/service"
)

func main() {
	var mode string
	var grpcPort string
	var restPort string
	var username, password, host, port, database string
	var pcu int

	flag.StringVar(&mode, "mode", "simple", "Select boot mode.")
	flag.StringVar(&restPort, "rest-port", "50050", "RESTful server port")
	flag.StringVar(&grpcPort, "grpc-port", "50051", "gRPC server port")
	flag.StringVar(&username, "username", "root", "MySQL DB's username")
	flag.StringVar(&password, "pd", "Zkyy2021", "MySQL DB's password")
	flag.StringVar(&host, "host", "10.1.38.245", "MySQL DB's host")
	flag.StringVar(&port, "port", "3307", "MySQL DB's port")
	flag.StringVar(&database, "db", "user_management", "MySQL DB's db name")
	flag.IntVar(&pcu, "pcu", 5, "Maximum number of concurrent online users")
	flag.Parse()

	service.PCU = pcu
	go service.CleanupExpiredTokens()

	var Db dao.ORM
	var err error
	switch mode {
	case "simple":
		Db, err = dao.NewMyBolt()
		if err != nil {
			fmt.Printf("Failed to connect BoltDB: %v\n", err)
		}
	case "mysql":
		Db, err = dao.NewMysql(dao.MysqlCfg{
			Username: username,
			Password: password,
			Host:     host,
			Port:     port,
			Database: database})
		if err != nil {
			fmt.Printf("Failed to connect MySql database: %v\n", err)
			return
		}
		fmt.Printf("Successfully connected to MySQL database at %s:%s, Database: %s\n", host, port, database)
	default:
		fmt.Printf("Failed to parse the %v\n mode because it does not exist", err)
		return
	}

	go func() {
		fmt.Printf("gRPC server started and listening on :%s port \n", grpcPort)
		listen, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen for gRPC: %v", err)
		}
		s := grpc.NewServer(
			grpc.UnaryInterceptor(controller.LoggingInterceptor),
		)
		gRPC := controller.NewGrpcController()
		pb.RegisterUserServiceServer(s, gRPC)
		if err := s.Serve(listen); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	go func() {
		server := service.Server{
			Db: Db,
		}
		fmt.Printf("RESTful server started and listening on :%s port \n", restPort)
		gin.SetMode(gin.ReleaseMode) // set release mode
		restAPI := controller.NewRestController(&server)
		r := gin.New()
		r.GET("/version", restAPI.Version)
		r.POST("/register", restAPI.Register)
		r.POST("/login", restAPI.Login)
		r.POST("/logout", restAPI.LogoutByToken)
		r.POST("/retrieve-token-for-user", restAPI.RetrieveTokenForUser)
		r.GET("/check-token-valid", restAPI.CheckTokenValid)
		r.GET("/search-role", restAPI.SearchRole)
		r.GET("/search-group", restAPI.SearchGroup)
		r.GET("/get-userid", restAPI.GetUserId)
		r.GET("/search-permission", restAPI.SearchPermission)
		r.PUT("/edit", restAPI.Edit)
		//r.DELETE("/del-role", restAPI.CheckTokenValid)
		//r.DELETE("/del-group", restAPI.CheckTokenValid)
		//r.POST("/create-role", restAPI.CheckTokenValid)
		//r.POST("/create-group", restAPI.CheckTokenValid)
		//r.POST("/add-group-members", restAPI.CheckTokenValid)
		//r.GET("/show-all-group", restAPI.CheckTokenValid)
		//r.GET("/show-all-role", restAPI.CheckTokenValid)
		if err := r.Run(":" + restPort); err != nil {
			log.Fatalf("Failed to serve RESTful: %v", err)
		}
	}()
	select {}
}
