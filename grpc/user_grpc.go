syntax = "proto3";      // একবারই লিখবে

package user;            // package name

option go_package = "golang-crud/proto/userpb"; // Go module + proto folder

// Service definition
service UserService {
  rpc AddUser (AddUserRequest) returns (UserResponse);
  rpc GetAllUsers (Empty) returns (UsersResponse);
  rpc GetUserById (GetUserRequest) returns (UserResponse);
  rpc UpdateUser (UpdateUserRequest) returns (UserResponse);
  rpc DeleteUser (DeleteUserRequest) returns (DeleteUserResponse);
  rpc TextSearch (TextSearchRequest) returns (UsersResponse);
}

// Request / Response messages
message AddUserRequest {
  string name = 1;
  string email = 2;
  string password = 3;
  string address = 4;
}

message GetUserRequest {
  int32 id = 1;
}

message UpdateUserRequest {
  int32 id = 1;
  string name = 2;
  string email = 3;
  string password = 4;
  string address = 5;
}

message DeleteUserRequest {
  int32 id = 1;
}

message TextSearchRequest {
  string query = 1;
}

message UserResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
  string password = 4;
  string address = 5;
}

message UsersResponse {
  repeated UserResponse users = 1;
}

message DeleteUserResponse {
  bool success = 1;
}

message Empty {}
