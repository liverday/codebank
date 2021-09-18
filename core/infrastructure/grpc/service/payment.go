package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/liverday/codebank/core/dto"
	"github.com/liverday/codebank/core/infrastructure/grpc/pb"
	"github.com/liverday/codebank/core/usecase"
)

type PaymentService struct {
	processTransactionUseCase usecase.UseCaseTransaction
	pb.UnimplementedPaymentServiceServer
}

func NewPaymentService(useCaseTransaction usecase.UseCaseTransaction) *PaymentService {
	return &PaymentService{
		processTransactionUseCase: useCaseTransaction,
	}
}

func (p *PaymentService) Payment(ctx context.Context, in *pb.PaymentRequest) (*empty.Empty, error) {
	newTransaction := dto.NewTransaction{
		OwnerName:       in.GetCreditCard().GetOwnerName(),
		Number:          in.GetCreditCard().GetNumber(),
		ExpirationMonth: in.GetCreditCard().GetExpirationMonth(),
		ExpirationYear:  in.GetCreditCard().GetExpirationYear(),
		CVV:             in.GetCreditCard().GetCvv(),
		Amount:          in.GetAmount(),
		Store:           in.GetStore(),
		Description:     in.GetDescription(),
	}

	transaction, err := p.processTransactionUseCase.ProcessTransaction(newTransaction)

	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}

	if transaction.Status != "approved" {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, "Transaction Rejected by the bank")
	}

	return &empty.Empty{}, nil
}
