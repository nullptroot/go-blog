package articleapi

import (
	"context"
	"encoding/json"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var DateCount = map[string]int{}

// ArticleCalendarView 按照日历分组
// @Tags 文章管理
// @Summary 按照日历分组
// @Description 按照日历分组
// @Router /api/articles/calander [get]
// @Produce json
// @Success 200 {object} res.Response{data=CalendarResponse}
func (ArticleApi) ArticleCalendarView(c *gin.Context) {

	// 时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")

	// 时间段搜索
	// 从今天开始，到去年的今天
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0)

	format := "2006-01-02 15:04:05"
	// lt 小于  gt 大于
	query := elastic.NewRangeQuery("created_at").
		Gte(aYearAgo.Format(format)).
		Lte(now.Format(format))

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}

	var data BucketsType
	_ = json.Unmarshal(result.Aggregations["calendar"], &data)

	var resList = make([]CalendarResponse, 0)
	for _, bucket := range data.Buckets {
		Time, _ := time.Parse(format, bucket.KeyAsString)
		DateCount[Time.Format("2006-01-02")] = bucket.DocCount
	}
	days := int(now.Sub(aYearAgo).Hours() / 24)
	for i := 0; i <= days; i++ {
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")

		count, _ := DateCount[day]
		resList = append(resList, CalendarResponse{
			Date:  day,
			Count: count,
		})
	}

	res.OkWithData(resList, c)

}
