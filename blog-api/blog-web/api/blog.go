package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"blog-api/blog-web/forms"
	"blog-api/blog-web/global"
	"blog-api/blog-web/global/response"
	"blog-api/blog-web/proto"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将 grpc 的code转换成 http 状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务端错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

func HandleValidator(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func GetBlogList(ctx *gin.Context) {
	// 拨号连接blog grpc 服务器
	blogConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetBlogList] 连接 [blog服务失败]",
			"msg",
			err.Error(),
		)
	}
	// 生成 grpc 的 client 并调用接口
	blogSrvClient := proto.NewBlogClient(blogConn)
	// 获取 url 参数
	pn := ctx.DefaultQuery("pn", "0")
	pSize := ctx.DefaultQuery("psize", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := blogSrvClient.GetBlogList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询【blog列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)

	for _, value := range rsp.Data {

		blog := response.BlogResponse{
			Id:          value.Id,
			Title:       value.Title,
			Tag:         value.Tag,
			Description: value.Description,
			Content:     value.Content,
			ArticleId:   value.ArticleId,
			AddTime:     response.JsonTime(time.Unix(int64(value.AddTime), 0)),
			UpdateTime:  response.JsonTime(time.Unix(int64(value.UpdateTime), 0)),
		}

		result = append(result, blog)
	}
	ctx.JSON(http.StatusOK, result)
}

func ArticleCreate(ctx *gin.Context) {
	// 新建博客参数验证
	articleForm := forms.ArticleForm{}
	if err := ctx.ShouldBindJSON(&articleForm); err != nil {
		HandleValidator(ctx, err)
		return
	}
	// 拨号连接blog grpc 服务器
	blogConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetBlogList] 连接 [blog服务失败]",
			"msg",
			err.Error(),
		)
	}
	// 生成 grpc 的 client 并调用接口
	blogSrvClient := proto.NewBlogClient(blogConn)
	article, err := blogSrvClient.CreateArticle(context.Background(), &proto.CreateArticleInfo{
		Title:       articleForm.Title,
		Tag:         articleForm.Tag,
		Description: articleForm.Description,
		Content:     articleForm.Content,
		ArticleId:   articleForm.ArticleId,
	})
	if err != nil {
		zap.S().Errorf("[ArticleCreate] 创建【文章失败】: %s", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":    article.Id,
		"Title": article.Title,
	})
}

// func CreateArticle(ctx *gin.Context) {
// 	// 拨号连接用户 grpc 服务器
// 	blogConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
// 	if err != nil {
// 		zap.S().Errorw("[GetBlogList] 连接 [用户服务失败]",
// 			"msg",
// 			err.Error(),
// 		)
// 	}
// 	// 生成 grpc 的 client 并调用接口
// 	blogSrvClient := proto.NewBlogClient(blogConn)
// 	// 如何获取参数

// }
