package main

import (
	"context"
	"fmt"

	pb "github.com/hanlray/at-search/proto/search"
	micro "github.com/micro/go-micro"
)

type service struct {
}

func (s *service) Search(ctx context.Context, req *pb.Query, res *pb.Response) error {
}

func main() {
	srv := micro.NewService(
		micro.Name("at.service.search"),
	)

	// Init will parse the command line flags.
	srv.Init()

	pb.RegisterSearchServiceServer(srv.Server(), &service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
