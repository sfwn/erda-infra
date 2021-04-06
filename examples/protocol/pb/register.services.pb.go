// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Source: greeter.proto, user.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterGreeterServiceImp greeter.proto
func RegisterGreeterServiceImp(regester transport.Register, srv GreeterServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterGreeterServiceHandler(regester, GreeterServiceHandler(srv), _ops.HTTP...)
	RegisterGreeterServiceServer(regester, srv, _ops.GRPC...)
}

// RegisterUserServiceImp user.proto
func RegisterUserServiceImp(regester transport.Register, srv UserServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterUserServiceHandler(regester, UserServiceHandler(srv), _ops.HTTP...)
	RegisterUserServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.infra.example.GreeterService",
		"erda.infra.example.UserService",
	)
}

var (
	greeterServiceClientType = reflect.TypeOf((*GreeterServiceClient)(nil)).Elem()
	greeterServiceServerType = reflect.TypeOf((*GreeterServiceServer)(nil)).Elem()

	userServiceClientType = reflect.TypeOf((*UserServiceClient)(nil)).Elem()
	userServiceServerType = reflect.TypeOf((*UserServiceServer)(nil)).Elem()
)

// GreeterServiceClientType .
func GreeterServiceClientType() reflect.Type { return greeterServiceClientType }

// GreeterServiceServerType .
func GreeterServiceServerType() reflect.Type { return greeterServiceServerType }

// UserServiceClientType .
func UserServiceClientType() reflect.Type { return userServiceClientType }

// UserServiceServerType .
func UserServiceServerType() reflect.Type { return userServiceServerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		greeterServiceClientType,
		userServiceClientType,
		// server types
		greeterServiceServerType,
		userServiceServerType,
	}
}
