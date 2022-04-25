package handlers

import (
	"context"
	"fmt"
	"github.com/iamkirillnb/avtodom/internal"
	pc "github.com/iamkirillnb/avtodom/url_proto"
	"log"
	"net/http"
)

type HttpClient struct {
	grpc *GrpcApi
	cfg  *internal.ServerConfig
	client *http.Client
}

func NewHttpClient(config *internal.ServerConfig, api *GrpcApi) *HttpClient {
	return &HttpClient{
		grpc:   api,
		cfg:    config,
		client: http.DefaultClient,
	}
}
func (h *HttpClient) FollowTheLink(url string) error {
	ps := &pc.GetUrlOuterRequest{Inner: url}
	result, err := h.grpc.GetUrlOuter(context.Background(), ps)
	if err != nil {
		log.Println("grpc get url outer failed")
		return err
	}
	u := result.Outer
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Println("new request failed")
		return err
	}
	res, err := h.client.Do(req)
	if err != nil {
		log.Println("client do request failed")
		return err
	}
	fmt.Println(res.Header)
	return nil
}

