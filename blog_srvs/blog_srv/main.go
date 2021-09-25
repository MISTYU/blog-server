package main

import (
	"blog_srvs/blog_srv/handler"
	"blog_srvs/blog_srv/proto"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 8000, "端口号")

	flag.Parse()
	fmt.Println("ip: ", *IP)
	fmt.Println("port: ", *Port)
	server := grpc.NewServer()
	proto.RegisterBlogServer(server, &handler.BlogServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
