package flag

import "go-blog/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()
}
