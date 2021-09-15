package handler

import (
	"blog_srvs/blog_srv/global"
	"blog_srvs/blog_srv/model"
	"blog_srvs/blog_srv/proto"

	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type BlogServer struct{}

func ModelToRsponse(blog model.Blog) proto.ArticleInfoReponse {
	// 在 grpc 的 message 中字段有默认值，不能随便赋值，序列化会出错
	blogInfoRsp := proto.ArticleInfoReponse{
		Id:          blog.ID,
		Title:       blog.Title,
		Tag:         blog.Tag,
		Description: blog.Description,
		Content:     blog.Content,
		ArticleId:   blog.ArticleId,
		AddTime:     blog.CreatedAt.Unix(),
		UpdateTime:  blog.UpdatedAt.Unix(),
	}
	return blogInfoRsp
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *BlogServer) GetBlogList(ctx context.Context, req *proto.PageInfo) (*proto.BlogListResponse, error) {
	// 获取blog列表
	var blogs []model.Blog
	result := global.DB.Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("blog列表")
	rsp := &proto.BlogListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&blogs)

	for _, user := range blogs {
		userInfoRsp := ModelToRsponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

func (s *BlogServer) UpdateArticle(ctx context.Context, req *proto.UpdateArticleInfo) (*empty.Empty, error) {
	// 更新文章
	var blog model.Blog
	result := global.DB.First(&blog, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "博客不存在")
	}

	blog.Title = req.Title
	blog.Tag = req.Tag
	blog.Description = req.Description
	blog.Content = req.Content

	result = global.DB.Save(&blog)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (s *BlogServer) GetArticleById(ctx context.Context, req *proto.IdRequest) (*proto.ArticleInfoReponse, error) {
	// 通过 id 查询 blog
	var blog model.Blog
	result := global.DB.First(&blog, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "博客不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModelToRsponse(blog)
	return &userInfoRsp, nil
}

func (s *BlogServer) CreateArticle(ctx context.Context, req *proto.CreateArticleInfo) (*proto.ArticleInfoReponse, error) {
	// 创建文章
	var blog model.Blog
	blog.ArticleId = req.ArticleId
	blog.Title = req.Title
	blog.Tag = req.Tag
	blog.Description = req.Description
	blog.Content = req.Content

	result := global.DB.Create(&blog)

	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	blogInfoRsp := ModelToRsponse(blog)
	return &blogInfoRsp, nil
}
