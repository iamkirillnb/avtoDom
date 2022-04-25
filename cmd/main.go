package main

import (
	"flag"
	"fmt"
	"github.com/iamkirillnb/avtodom/internal"
	"github.com/iamkirillnb/avtodom/internal/handlers"
	"github.com/iamkirillnb/avtodom/internal/repos"
	"log"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config Path", "dev.yaml", "path to dev.yaml file")
}

func main() {
	flag.Parse()

	fmt.Println("AvtoDom started")

	cfg := internal.GetConfig(cfgPath)

	db := repos.NewPostgres(&cfg.Db)

	repo := repos.NewDbRepo(db)

	handler := handlers.NewGrpcApi(&cfg.Server, repo)

	httpClient := handlers.NewHttpClient(&cfg.Server, handler)

	checkUrl := "https://yandex.ru"
	err := httpClient.FollowTheLink(checkUrl)
	if err != nil {
		log.Println("http client request failed")
		log.Fatal(err)
	}



}

