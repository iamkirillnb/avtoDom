package handlers

import (
	"context"
	"github.com/iamkirillnb/avtodom/internal"
	"github.com/iamkirillnb/avtodom/internal/repos"
	pc "github.com/iamkirillnb/avtodom/url_proto"
	"google.golang.org/grpc"
	"log"
	"net"
)


type GrpcApi struct {
	cfg *internal.ServerConfig
	repo *repos.DbRepo

	*grpc.Server
}


func NewGrpcApi(config *internal.ServerConfig, repo *repos.DbRepo) *GrpcApi {

	s := &GrpcApi{
		repo: repo,
		cfg:    config,
	}

	serv := grpc.NewServer()

	pc.RegisterUrlRedirectServer(serv, s)

	s.Server = serv

	return s

}

func (g GrpcApi) GetUrlOuter(ctx context.Context, request *pc.GetUrlOuterRequest) (*pc.GetUrlOuterResponse, error) {
	innerUrl := request.Inner
	redirectTo := g.repo.GetByInnerUrl(innerUrl)

	result := &pc.GetUrlOuterResponse{Outer: redirectTo}
	return result, nil
}


func (g GrpcApi) Start() error {
	l, err := net.Listen("tcp", g.cfg.Address())
	if err != nil {
		log.Printf("listen tcp %v failed\n", g.cfg.Address())
		return err
	}
	err = g.Serve(l)
	if err != nil {
		log.Println("serve failed")
		return err
	}
	return nil
}