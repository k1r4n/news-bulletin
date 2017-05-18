package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/news-bulletin-cron/conf"
	"github.com/news-bulletin-cron/models"
)

// NewsAPI Object
type NewsAPI struct {
	conf *conf.Vars
	log  *logrus.Logger
}

// NewsAPIInstant create an object of NewsAPI service
func NewsAPIInstant(c *conf.Vars, l *logrus.Logger) NewsAPI {
	return NewsAPI{conf: c, log: l}
}

//FetchMediaList fetches list of medias from newsapi.org
func (n NewsAPI) FetchMediaList() ([]models.MediaList, []models.MediaList, error) {
	fmt.Println("Getting MediaList \n ")
	url := n.conf.ChannelEndPoint
	var mediaList models.MediaListResponse
	var mediaTopList []models.MediaList
	var mediaLatestList []models.MediaList
	resp, err := http.Get(url)
	if err != nil {
		n.log.Error(err)
		return mediaTopList, mediaLatestList, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		n.log.Error(err)
		return mediaTopList, mediaLatestList, err
	}
	err = json.Unmarshal(body, &mediaList)
	if err != nil {
		n.log.Error(err)
		return mediaTopList, mediaLatestList, err
	}
	for i := 0; i < len(mediaList.Sources); i++ {
		for j := 0; j < len(mediaList.Sources[i].SortBysAvailable); j++ {
			if mediaList.Sources[i].SortBysAvailable[j] == "top" {
				fmt.Println("Top channel \n\n", mediaList.Sources[i].Name)
				mediaTopList = append(mediaTopList, mediaList.Sources[i])
			}
			if mediaList.Sources[i].SortBysAvailable[j] == "latest" {
				fmt.Println("Latest channel \n\n", mediaList.Sources[i].Name)
				mediaLatestList = append(mediaLatestList, mediaList.Sources[i])
			}
		}
	}
	return mediaTopList, mediaLatestList, nil
}

// FetchArticleList fetchs all articles from newsapi.org
func (n NewsAPI) FetchArticleList(mediaTopList []models.MediaList, mediaLatestList []models.MediaList) ([]models.ArticleList, []models.ArticleList, error) {
	fmt.Println("Getting articles \n ")
	var articleList models.ArticleList
	var articleTopList []models.ArticleList
	var articleLatestList []models.ArticleList
	for i := 0; i < len(mediaTopList); i++ {
		fmt.Println("Getting top articles \n ")
		url := n.conf.ArticleEndPoint + "?source=" + mediaTopList[i].ID + "&sortBy=top&apiKey=" + n.conf.APIKey
		resp, err := http.Get(url)
		if err != nil {
			n.log.Error(err)
			var a []models.ArticleList
			var b []models.ArticleList
			return a, b, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			n.log.Error(err)
			var a []models.ArticleList
			var b []models.ArticleList
			return a, b, err
		}
		articleList = models.ArticleList{}
		err = json.Unmarshal(body, &articleList)

		articleTopList = append(articleTopList, articleList)
	}
	for i := 0; i < len(mediaLatestList); i++ {
		fmt.Println("Getting latest articles \n ")
		url := n.conf.ArticleEndPoint + "?source=" + mediaLatestList[i].ID + "&sortBy=latest&apiKey=" + n.conf.APIKey
		resp, err := http.Get(url)
		if err != nil {
			n.log.Error(err)
			var a []models.ArticleList
			var b []models.ArticleList
			return a, b, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			n.log.Error(err)
			var a []models.ArticleList
			var b []models.ArticleList
			return a, b, err
		}
		err = json.Unmarshal(body, &articleList)
		articleLatestList = append(articleLatestList, articleList)
	}
	return articleTopList, articleLatestList, nil
}
