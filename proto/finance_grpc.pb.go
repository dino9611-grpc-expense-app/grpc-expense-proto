// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/finance.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExpenseClient is the client API for Expense service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpenseClient interface {
	GetFinanceDailyByUserId(ctx context.Context, in *FinanceReq, opts ...grpc.CallOption) (*FinanceRes, error)
	GetFinanceMonthlyByUserId(ctx context.Context, in *FinanceReq, opts ...grpc.CallOption) (*FinanceRes, error)
	GetFinanceTotalByUserId(ctx context.Context, in *FinanceReq, opts ...grpc.CallOption) (*FinanceRes, error)
	AddIncome(ctx context.Context, in *IncomeReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddExpense(ctx context.Context, in *ExpenseReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type expenseClient struct {
	cc grpc.ClientConnInterface
}

func NewExpenseClient(cc grpc.ClientConnInterface) ExpenseClient {
	return &expenseClient{cc}
}

func (c *expenseClient) GetFinanceDailyByUserId(ctx context.Context, in *FinanceReq, opts ...grpc.CallOption) (*FinanceRes, error) {
	out := new(FinanceRes)
	err := c.cc.Invoke(ctx, "/Expense/GetFinanceDailyByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseClient) GetFinanceMonthlyByUserId(ctx context.Context, in *FinanceReq, opts ...grpc.CallOption) (*FinanceRes, error) {
	out := new(FinanceRes)
	err := c.cc.Invoke(ctx, "/Expense/GetFinanceMonthlyByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseClient) GetFinanceTotalByUserId(ctx context.Context, in *FinanceReq, opts ...grpc.CallOption) (*FinanceRes, error) {
	out := new(FinanceRes)
	err := c.cc.Invoke(ctx, "/Expense/GetFinanceTotalByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseClient) AddIncome(ctx context.Context, in *IncomeReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Expense/AddIncome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseClient) AddExpense(ctx context.Context, in *ExpenseReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Expense/AddExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpenseServer is the server API for Expense service.
// All implementations must embed UnimplementedExpenseServer
// for forward compatibility
type ExpenseServer interface {
	GetFinanceDailyByUserId(context.Context, *FinanceReq) (*FinanceRes, error)
	GetFinanceMonthlyByUserId(context.Context, *FinanceReq) (*FinanceRes, error)
	GetFinanceTotalByUserId(context.Context, *FinanceReq) (*FinanceRes, error)
	AddIncome(context.Context, *IncomeReq) (*emptypb.Empty, error)
	AddExpense(context.Context, *ExpenseReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedExpenseServer()
}

// UnimplementedExpenseServer must be embedded to have forward compatible implementations.
type UnimplementedExpenseServer struct {
}

func (UnimplementedExpenseServer) GetFinanceDailyByUserId(context.Context, *FinanceReq) (*FinanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinanceDailyByUserId not implemented")
}
func (UnimplementedExpenseServer) GetFinanceMonthlyByUserId(context.Context, *FinanceReq) (*FinanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinanceMonthlyByUserId not implemented")
}
func (UnimplementedExpenseServer) GetFinanceTotalByUserId(context.Context, *FinanceReq) (*FinanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinanceTotalByUserId not implemented")
}
func (UnimplementedExpenseServer) AddIncome(context.Context, *IncomeReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddIncome not implemented")
}
func (UnimplementedExpenseServer) AddExpense(context.Context, *ExpenseReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExpense not implemented")
}
func (UnimplementedExpenseServer) mustEmbedUnimplementedExpenseServer() {}

// UnsafeExpenseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpenseServer will
// result in compilation errors.
type UnsafeExpenseServer interface {
	mustEmbedUnimplementedExpenseServer()
}

func RegisterExpenseServer(s grpc.ServiceRegistrar, srv ExpenseServer) {
	s.RegisterService(&Expense_ServiceDesc, srv)
}

func _Expense_GetFinanceDailyByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServer).GetFinanceDailyByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Expense/GetFinanceDailyByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServer).GetFinanceDailyByUserId(ctx, req.(*FinanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Expense_GetFinanceMonthlyByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServer).GetFinanceMonthlyByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Expense/GetFinanceMonthlyByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServer).GetFinanceMonthlyByUserId(ctx, req.(*FinanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Expense_GetFinanceTotalByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServer).GetFinanceTotalByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Expense/GetFinanceTotalByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServer).GetFinanceTotalByUserId(ctx, req.(*FinanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Expense_AddIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncomeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServer).AddIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Expense/AddIncome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServer).AddIncome(ctx, req.(*IncomeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Expense_AddExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServer).AddExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Expense/AddExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServer).AddExpense(ctx, req.(*ExpenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Expense_ServiceDesc is the grpc.ServiceDesc for Expense service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Expense_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Expense",
	HandlerType: (*ExpenseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFinanceDailyByUserId",
			Handler:    _Expense_GetFinanceDailyByUserId_Handler,
		},
		{
			MethodName: "GetFinanceMonthlyByUserId",
			Handler:    _Expense_GetFinanceMonthlyByUserId_Handler,
		},
		{
			MethodName: "GetFinanceTotalByUserId",
			Handler:    _Expense_GetFinanceTotalByUserId_Handler,
		},
		{
			MethodName: "AddIncome",
			Handler:    _Expense_AddIncome_Handler,
		},
		{
			MethodName: "AddExpense",
			Handler:    _Expense_AddExpense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/finance.proto",
}
