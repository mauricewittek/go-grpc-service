package grpc

import (
	"context"
	"log"
	"net"

	"github.com/mauricewittek/go-grpc-service/internal/rocket"
	rkt "github.com/mauricewittek/tutorial-protos/rocket/v1"
	"google.golang.org/grpc"
)

type RocketService interface {
	GetRocket(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

type Handler struct {
	RocketService RocketService
}

func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
		return err
	}

	return nil
}

func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	rocket, err := h.RocketService.GetRocket(ctx, req.Id)
	if err != nil {
		log.Print("Failed to retrieve rocket by ID")
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	newRkt, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID:   req.Rocket.Id,
		Type: req.Rocket.Type,
		Name: req.Rocket.Name,
	})
	if err != nil {
		log.Print("failed to insert rocket into db")
		return &rkt.AddRocketResponse{}, err
	}

	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   newRkt.ID,
			Type: newRkt.Type,
			Name: newRkt.Name,
		},
	}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	err := h.RocketService.DeleteRocket(ctx, req.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{}, err
	}

	return &rkt.DeleteRocketResponse{
		Status: "successfully deleted rocket",
	}, nil
}
