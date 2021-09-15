package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"blog_srvs/blog_srv/proto"
)

var blogClient proto.BlogClient

var conn *grpc.ClientConn

func Init() {
	// var err error
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	blogClient = proto.NewBlogClient(conn)
}

func TestGetBlogList() {
	rsp, err := blogClient.GetBlogList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, blog := range rsp.Data {
		fmt.Println(blog.Id)
	}
}

func TestCreateBlog() {
	for i := 0; i < 10; i++ {
		rsp, err := blogClient.CreateArticle(context.Background(), &proto.CreateArticleInfo{
			Title:       fmt.Sprintf("yiyue%d", i),
			Tag:         fmt.Sprintf("前端%d", i),
			Description: fmt.Sprintf("测试%d", i),
			Content:     "nyf&lch",
			ArticleId:   fmt.Sprintf("yiyue%d", i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func main() {
	Init()
	// TestCreateBlog()
	TestGetBlogList()

	conn.Close()
}
