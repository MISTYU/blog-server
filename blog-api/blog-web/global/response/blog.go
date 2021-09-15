package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type BlogResponse struct {
	Id          int32    `json:"id"`
	Title       string   `json:"title"`
	Tag         string   `json:"tag"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	ArticleId   string   `json:"articleId"`
	AddTime     JsonTime `json:"addTime"`
	UpdateTime  JsonTime `json:"updateTime"`
}
