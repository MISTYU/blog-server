package forms

type ArticleForm struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Tag         string `form:"tag" json:"tag" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	ArticleId   string `form:"articleId" json:"articleId"`
}
