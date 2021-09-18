package server

import (
	"log"
	"net"

	"github.com/liverday/codebank/core/infrastructure/grpc/pb"
	"github.com/liverday/codebank/core/infrastructure/grpc/service"
	"github.com/liverday/codebank/core/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	processTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer(processTransactionUseCase usecase.UseCaseTransaction) GRPCServer {
	return GRPCServer{
		processTransactionUseCase: processTransactionUseCase,
	}
}

func (g GRPCServer) Serve() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("could not start grpc server")
	}

	paymentService := service.NewPaymentService(g.processTransactionUseCase)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, paymentService)
	grpcServer.Serve(listener)
}
