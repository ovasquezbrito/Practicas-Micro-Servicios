package service

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"github.com/google/uuid"
	"google.golang.org/grpc/status"

	"github.com/techschool/pcbook/pb"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{store}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("recive a create-lapto request with id %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new lapto ID: %v", err)
		}
		laptop.Id = id.String()
	}

	// some heavy processing
	time.Sleep(6 * time.Second)

	// cuando el cliente interrumpe la petición

	if ctx.Err() == context.Canceled {
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Print("diadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "diadline is exceeded")
	}

	//aqui se debería guardar en la bases de dtos
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if err == ErrAlreadyExists {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to the store %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}
