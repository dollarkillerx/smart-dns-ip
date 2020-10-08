package main

import (
	"context"
	"log"
	"net"

	"github.com/dollarkillerx/smart-dns-ip/generate/smart_dns_ip"
	"github.com/dollarkillerx/smart-dns-ip/pkg/config"
	"github.com/dollarkillerx/smart-dns-ip/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	go rpcServer()
	httpServer()
}

func httpServer() {
	log.Println("RestfulAddr: ", config.BaseConfig.RestfulAddr)
	app := gin.New()
	app.Use(gin.Recovery())
	if config.BaseConfig.Debug {
		app.Use(gin.Logger())
	}

	app.GET("/ip_search", func(context *gin.Context) {
		ip := context.Query("ip")
		if ip == "" {
			context.JSON(400, gin.H{"message": "params error"})
			return
		}

		log.Println(ip)
		search, err := storage.IP2.MemorySearch(ip)
		if err != nil {
			context.JSON(500, err)
			return
		}

		context.JSON(200, search)
	})

	if err := app.Run(config.BaseConfig.RestfulAddr); err != nil {
		log.Fatalln(err)
	}
}

type server struct{}

func (s *server) IPSearch(ctx context.Context, request *smart_dns_ip.IPSearchRequest) (resp *smart_dns_ip.IPSearchResponse, err error) {
	search, err := storage.IP2.MemorySearch(request.Ip)
	if err != nil {
		return nil, err
	}
	return &smart_dns_ip.IPSearchResponse{
		CityId:  search.CityId,
		Country: search.Country,
		Region:  search.Region,
		City:    search.City,
		Isp:     search.ISP,
	}, nil
}

func rpcServer() {
	log.Println("GRPCAddr: ", config.BaseConfig.GRPCAddr)

	listen, err := net.Listen("tcp", config.BaseConfig.GRPCAddr)
	if err != nil {
		log.Fatalln(err)
	}

	ser := grpc.NewServer()
	smart_dns_ip.RegisterIPSearchServer(ser, &server{})
	reflection.Register(ser)

	if e := ser.Serve(listen); e != nil {
		log.Fatalln(e)
	}
}
