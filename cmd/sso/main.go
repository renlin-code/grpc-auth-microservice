package main

import (
	"fmt"

	"github.com/renlin-code/grpc-sso-microservice/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
