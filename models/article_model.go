package models

import (
	"context"
	"encoding/json"
	"go-blog/global"
	"go-blog/models/ctype"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// ArticleModel 文章表

type ArticleModel struct {
	ID        string `json:"id" struct:"id"`                 // es的id
	CreatedAt string `json:"created_at" struct:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" struct:"updated_at"` // 更新时间

	Title    string `json:"title" struct:"title"`                // 文章标题
	Keyword  string `json:"keyword,omit(list)" struct:"keyword"` // 关键字
	Abstract string `json:"abstract" struct:"abstract"`          // 文章简介
	Content  string `json:"content,omit(list)" struct:"content"` // 文章内容

	LookCount     int `json:"look_count" struct:"look_count"`         // 浏览量
	CommentCount  int `json:"comment_count" struct:"comment_count"`   // 评论量
	DiggCount     int `json:"digg_count" struct:"digg_count"`         // 点赞量
	CollectsCount int `json:"collects_count" struct:"collects_count"` // 收藏量

	UserID       uint   `json:"user_id" struct:"user_id"`               // 用户id
	UserNickName string `json:"user_nick_name" struct:"user_nick_name"` //用户昵称
	UserAvatar   string `json:"user_avatar" struct:"user_avatar"`       // 用户头像

	Category string `json:"category" struct:"category"`        // 文章分类
	Source   string `json:"source,omit(list)" struct:"source"` // 文章来源
	Link     string `json:"link,omit(list)" struct:"link"`     // 原文链接

	BannerID  uint   `json:"banner_id" struct:"banner_id"`   // 文章封面id
	BannerUrl string `json:"banner_url" struct:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags" struct:"tags"` // 文章标签
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  },
  "mappings": {
		"properties": {
			"title": { 
				"type": "text"
			},
			"keyword": { 
				"type": "keyword"
			},
			"abstract": { 
				"type": "text"
			},
			"content": { 
				"type": "text"
			},
			"look_count": {
				"type": "long"
			},
			"comment_count": {
				"type": "long"
			},
			"digg_count": {
				"type": "long"
			},
			"collects_count": {
				"type": "long"
			},
			"user_id": {
				"type": "long"
			},
			"user_nick_name": { 
				"type": "keyword"
			},
			"user_avatar": { 
				"type": "keyword"
			},
			"category": { 
				"type": "keyword"
			},
			"source": { 
				"type": "keyword"
			},
			"link": { 
				"type": "keyword"
			},
			"banner_id": {
				"type": "long"
			},
			"banner_url": { 
				"type": "keyword"
			},
			"tags": { 
				"type": "keyword"
			},
			"created_at":{
				"type": "date",
				"null_value": "null",
				"format": "[yyyy-MM-dd HH:mm:ss]"
			},
			"updated_at":{
				"type": "date",
				"null_value": "null",
				"format": "[yyyy-MM-dd HH:mm:ss]"
			}
		}
	}
}
`
}

// IndexExists 索引是否存在
func (a ArticleModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

// http://10.0.0.3:9200/article_index/blog/Al0jQ48BVn2YyNCXpCLI  article_index索引 blog类型 Al0jQ48BVn2YyNCXpCLI id号
// CreateIndex 创建索引
func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		// 有索引
		a.RemoveIndex()
	}
	// 没有索引
	// 创建索引
	// 需要IncludeTypeName，要不然不行 然后加上还有警告说type将被舍弃
	createIndex, err := global.ESClient.
		CreateIndex(a.Index()).
		Body(a.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}
	logrus.Infof("索引 %s 创建成功", a.Index())
	return nil
}

// RemoveIndex 删除索引
func (a ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}

// Create 添加的方法
func (a ArticleModel) Create() (err error) {
	// 不加Type是不可以的 报错
	indexResponse, err := global.ESClient.Index().
		Index(a.Index()).
		BodyJson(a).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// ISExistData 是否存在该文章
func (a ArticleModel) ISExistData() bool {
	res, err := global.ESClient.
		Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

func (a *ArticleModel) GetDataByID(id string) error {
	res, err := global.ESClient.
		Get().
		Index(a.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return err
	}
	err = json.Unmarshal(res.Source, a)
	return err
}
