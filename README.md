# 用户管理模块

​	当构建现代应用程序时，微服务架构已经成为了一种强大的架构范式。用户管理服务是一个典型的微服务示例，它体现了微服务架构的核心理念：组件化、低耦合、高内聚。本服务旨在为应用程序提供强大的用户管理功能，通过RESTful和gRPC接口，提供了便捷的用户注册、登录、注销等操作。

​	我们的设计理念是将用户管理功能模块化，将其拆分成一个独立的微服务，以降低系统内各模块之间的耦合度。通过RESTful和gRPC接口，其他项目可以轻松地与用户管理服务集成，无需关注底层实现细节，从而加速了应用程序的开发和维护。这种模块化的设计让我们的服务具有高度可扩展性，可以根据需要添加新的功能模块，而不会影响现有的系统。

​	用户管理服务的目标是提供简单、可靠、安全的用户管理功能，使应用程序能够专注于核心业务逻辑，而不必担心用户身份验证和授权等复杂问题。无论是开发Web应用、移动应用还是其他类型的应用，用户管理服务都可以轻松集成，提供一致的用户体验。

​	总之，我们的用户管理服务代表了现代微服务架构的精髓，通过模块化和松耦合的设计，为开发人员提供了一个强大的工具，使他们能够更快速、更灵活地构建应用程序，将精力集中在核心业务上，提高了整体开发效率和可维护性。



## 快速开始

### 源码安装

```bash
go mod tidy 
go run main.go --mode="simple" --grpc-port="50051" --rest-port="50050"
```

### 二进制文件安装

```bash
.\main.exe --mode="simple" --grpc-port="50051" --rest-port="50050"  // windows x64 启动
./main --mode="simple" --grpc-port="50051" --rest-port="50050"      // linux 系统启动
```

控制台出现以下界面即为启动成功。

![image-20230922105252799](https://mr-lai.oss-cn-zhangjiakou.aliyuncs.com/imgs/image-20230922105252799.png)



## gRPC调用

### protoc 文件

```protobuf
syntax = "proto3";  // Use proto3 syntax

package userService;  // Define the package name of the proto file
option go_package = "protoc/user_service";

service UserService {
  rpc Register (RegisterReq) returns (Stdout);
  rpc Login (LoginReq) returns (Stdout);
  rpc Logout (NoneReq) returns (Stdout);
  rpc RoleQuery (NoneReq) returns (Stdout);
  rpc GroupQuery (NoneReq) returns (Stdout);
  rpc PermissionQuery (NoneReq) returns (Stdout);
  rpc Edit (EditReq) returns (Stdout);
  rpc DelRole (DelRoleReq) returns (Stdout);
  rpc DelGroup (DelGroupReq) returns (Stdout);
  rpc CreateRole (CreateRoleReq) returns (Stdout);
  rpc CreateGroup (CreateGroupReq) returns (Stdout);
  rpc AddGroupMembers (AddGroupMembersReq) returns (Stdout);
  rpc ShowAllGroup (ShowAllGroupReq) returns (Stdout);
  rpc ShowAllRole (ShowAllRoleReq) returns (Stdout);
}
message RegisterReq {
  string username = 1;
  string password = 2;
  string email = 3;
  string phone = 4;
  string full_name = 5;
  string profilePicture = 6;
}
message Stdout {
  uint32 status_code = 1;
  string message = 2;
}
message LoginReq {
  string user_name = 1;
  string password = 3;
}
message NoneReq{}
message DelRoleReq{
  string role_name = 1;
}
message DelGroupReq{
  string group_name = 1;
}
message EditReq{
  string username = 1;
  string email = 2;
  string phone = 3;
  string full_name = 4;
  string profilePicture = 5;
}
message CreateRoleReq{
  string role_name = 1;
  string description = 2;
  string permissions = 3;
}
message CreateGroupReq{
  string group_name = 1;
  string description = 2;
  string permissions = 3;
}
message AddGroupMembersReq{
  string group_name = 1;
  string user_name = 2;
  string description = 3;
  string permissions = 4;
}
message ShowAllGroupReq{
  string group_name = 1;
}
message ShowAllRoleReq{
  string role_name = 1;
}
```

### 如何调用

#### Go

```bash
cd user_management_golang
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc # 安装GO的插件
protoc --go_out=. --go-grpc_out=. .\user_service.proto
```

#### Python

```bash
pip install grpcio grpcio-tools  # 安装Python相对应的插件
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. user_service.proto
```



## 路由说明

#### RESTful

| 序号 | 请求方式 | 路由地址  | 输入                                                         | 输出                                                         |
| ---- | -------- | --------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 1    | Post     | /register | username string<br />email string<br />phone string<br />full_name<br />profile_picture | StatusCode: 200 （注册成功）<br />StatusCode: 409   (用户已存在)<br />StatusCode: 400 （无效输入）<br />StatusCode: 500   (服务器未知错误) |
| 2    | Post     | /login    |                                                              |                                                              |
| 3    | Post     | /logout   |                                                              |                                                              |
|      |          |           |                                                              |                                                              |



#### gRPC

| 序号 | 路由地址 | 输入 | 输出 |
| ---- | -------- | ---- | ---- |
| 1    |          |      |      |
|      |          |      |      |
|      |          |      |      |

