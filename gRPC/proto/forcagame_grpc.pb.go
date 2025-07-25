// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: proto/forcagame.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GameService_CriarJogo_FullMethodName       = "/ForcaGame.GameService/CriarJogo"
	GameService_EntrarJogo_FullMethodName      = "/ForcaGame.GameService/EntrarJogo"
	GameService_PalpitarLetra_FullMethodName   = "/ForcaGame.GameService/PalpitarLetra"
	GameService_PalpitarPalavra_FullMethodName = "/ForcaGame.GameService/PalpitarPalavra"
	GameService_PedirDica_FullMethodName       = "/ForcaGame.GameService/PedirDica"
	GameService_ObterEstado_FullMethodName     = "/ForcaGame.GameService/ObterEstado"
)

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	// Criação de jogos
	CriarJogo(ctx context.Context, in *CriarJogoRequest, opts ...grpc.CallOption) (*CriarJogoResponse, error)
	EntrarJogo(ctx context.Context, in *EntrarJogoRequest, opts ...grpc.CallOption) (*EntrarJogoResponse, error)
	// Ações no jogo
	PalpitarLetra(ctx context.Context, in *PalpitarLetraRequest, opts ...grpc.CallOption) (*AtualizacaoResponse, error)
	PalpitarPalavra(ctx context.Context, in *PalpitarPalavraRequest, opts ...grpc.CallOption) (*AtualizacaoResponse, error)
	PedirDica(ctx context.Context, in *DicaRequest, opts ...grpc.CallOption) (*AtualizacaoResponse, error)
	// Consulta de estado (opcional)
	ObterEstado(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*AtualizacaoResponse, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) CriarJogo(ctx context.Context, in *CriarJogoRequest, opts ...grpc.CallOption) (*CriarJogoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CriarJogoResponse)
	err := c.cc.Invoke(ctx, GameService_CriarJogo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) EntrarJogo(ctx context.Context, in *EntrarJogoRequest, opts ...grpc.CallOption) (*EntrarJogoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EntrarJogoResponse)
	err := c.cc.Invoke(ctx, GameService_EntrarJogo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) PalpitarLetra(ctx context.Context, in *PalpitarLetraRequest, opts ...grpc.CallOption) (*AtualizacaoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AtualizacaoResponse)
	err := c.cc.Invoke(ctx, GameService_PalpitarLetra_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) PalpitarPalavra(ctx context.Context, in *PalpitarPalavraRequest, opts ...grpc.CallOption) (*AtualizacaoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AtualizacaoResponse)
	err := c.cc.Invoke(ctx, GameService_PalpitarPalavra_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) PedirDica(ctx context.Context, in *DicaRequest, opts ...grpc.CallOption) (*AtualizacaoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AtualizacaoResponse)
	err := c.cc.Invoke(ctx, GameService_PedirDica_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ObterEstado(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*AtualizacaoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AtualizacaoResponse)
	err := c.cc.Invoke(ctx, GameService_ObterEstado_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility.
type GameServiceServer interface {
	// Criação de jogos
	CriarJogo(context.Context, *CriarJogoRequest) (*CriarJogoResponse, error)
	EntrarJogo(context.Context, *EntrarJogoRequest) (*EntrarJogoResponse, error)
	// Ações no jogo
	PalpitarLetra(context.Context, *PalpitarLetraRequest) (*AtualizacaoResponse, error)
	PalpitarPalavra(context.Context, *PalpitarPalavraRequest) (*AtualizacaoResponse, error)
	PedirDica(context.Context, *DicaRequest) (*AtualizacaoResponse, error)
	// Consulta de estado (opcional)
	ObterEstado(context.Context, *EstadoRequest) (*AtualizacaoResponse, error)
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGameServiceServer struct{}

func (UnimplementedGameServiceServer) CriarJogo(context.Context, *CriarJogoRequest) (*CriarJogoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CriarJogo not implemented")
}
func (UnimplementedGameServiceServer) EntrarJogo(context.Context, *EntrarJogoRequest) (*EntrarJogoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntrarJogo not implemented")
}
func (UnimplementedGameServiceServer) PalpitarLetra(context.Context, *PalpitarLetraRequest) (*AtualizacaoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PalpitarLetra not implemented")
}
func (UnimplementedGameServiceServer) PalpitarPalavra(context.Context, *PalpitarPalavraRequest) (*AtualizacaoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PalpitarPalavra not implemented")
}
func (UnimplementedGameServiceServer) PedirDica(context.Context, *DicaRequest) (*AtualizacaoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirDica not implemented")
}
func (UnimplementedGameServiceServer) ObterEstado(context.Context, *EstadoRequest) (*AtualizacaoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObterEstado not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}
func (UnimplementedGameServiceServer) testEmbeddedByValue()                     {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	// If the following call pancis, it indicates UnimplementedGameServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_CriarJogo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CriarJogoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CriarJogo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_CriarJogo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CriarJogo(ctx, req.(*CriarJogoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_EntrarJogo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntrarJogoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).EntrarJogo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_EntrarJogo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).EntrarJogo(ctx, req.(*EntrarJogoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_PalpitarLetra_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PalpitarLetraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).PalpitarLetra(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_PalpitarLetra_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).PalpitarLetra(ctx, req.(*PalpitarLetraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_PalpitarPalavra_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PalpitarPalavraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).PalpitarPalavra(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_PalpitarPalavra_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).PalpitarPalavra(ctx, req.(*PalpitarPalavraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_PedirDica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DicaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).PedirDica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_PedirDica_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).PedirDica(ctx, req.(*DicaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ObterEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstadoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ObterEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_ObterEstado_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ObterEstado(ctx, req.(*EstadoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ForcaGame.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CriarJogo",
			Handler:    _GameService_CriarJogo_Handler,
		},
		{
			MethodName: "EntrarJogo",
			Handler:    _GameService_EntrarJogo_Handler,
		},
		{
			MethodName: "PalpitarLetra",
			Handler:    _GameService_PalpitarLetra_Handler,
		},
		{
			MethodName: "PalpitarPalavra",
			Handler:    _GameService_PalpitarPalavra_Handler,
		},
		{
			MethodName: "PedirDica",
			Handler:    _GameService_PedirDica_Handler,
		},
		{
			MethodName: "ObterEstado",
			Handler:    _GameService_ObterEstado_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/forcagame.proto",
}
