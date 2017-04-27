package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/news-bulletin-cron/conf"
	"github.com/news-bulletin-cron/models"
	"github.com/news-bulletin-cron/services"
)

var (
	errorLog          *logrus.Logger
	mediaList         models.MediaList
	articleTopList    []models.ArticleList
	articleLatestList []models.ArticleList
	mediaTopList      []models.MediaList
	mediaLatestList   []models.MediaList
)

//init function
func init() {
	errorLog = logrus.New()
	errorLog.Formatter = new(logrus.JSONFormatter)

}

func main() {
	conf := conf.Read()
	newsAPI := services.NewsAPIInstant(conf, errorLog, mediaList, mediaTopList, mediaLatestList, articleTopList, articleLatestList)
	newsAPI.FetchMediaList()
	fmt.Println(newsAPI, "Main")
}
