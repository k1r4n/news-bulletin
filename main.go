package main

import (
	"context"

	"github.com/Sirupsen/logrus"
	"github.com/news-bulletin-cron/conf"
	"github.com/news-bulletin-cron/models"
	"github.com/news-bulletin-cron/services"
	elastic "gopkg.in/olivere/elastic.v5"
)

var (
	errorLog          *logrus.Logger
	mediaList         models.MediaListResponse
	articleTopList    []models.ArticleList
	articleLatestList []models.ArticleList
	articleList       models.ArticleList
	mediaTopList      []models.MediaList
	mediaLatestList   []models.MediaList
	elasticClient     *elastic.Client
	err               error
	ctx               context.Context
)

//init function
func init() {
	errorLog = logrus.New()
	errorLog.Formatter = new(logrus.JSONFormatter)
	elasticClient, err = elastic.NewClient()
	if err != nil {
		panic(err)
	}
	ctx = context.Background()
}

func main() {
	conf := conf.Read()
	newsAPI := services.NewsAPIInstant(conf, errorLog, mediaList, mediaTopList, mediaLatestList, articleList, articleTopList, articleLatestList)
	mediaTopList, mediaLatestList, err := newsAPI.FetchMediaList()
	if err != nil {
		panic(err)
	}
	articleTopList, articleLatestList, err = newsAPI.FetchArticleList(mediaTopList, mediaLatestList)
	if err != nil {
		panic(err)
	}
	elastic := services.NewElasticInstant(ctx, conf, errorLog, elasticClient, articleTopList, articleLatestList, mediaTopList, mediaLatestList)
	err = elastic.UpdateDatabase()
	if err != nil {
		panic(err)
	}
}
