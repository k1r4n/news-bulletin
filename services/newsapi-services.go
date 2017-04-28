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
	conf              *conf.Vars
	log               *logrus.Logger
	mediaList         models.MediaListResponse
	mediaTopList      []models.MediaList
	mediaLatestList   []models.MediaList
	articleList       models.ArticleList
	articleTopList    []models.ArticleList
	articleLatestList []models.ArticleList
}

// NewsAPIInstant create an object of NewsAPI service
func NewsAPIInstant(c *conf.Vars, l *logrus.Logger, m models.MediaListResponse, mt []models.MediaList, ml []models.MediaList, a models.ArticleList, at []models.ArticleList, al []models.ArticleList) NewsAPI {
	return NewsAPI{conf: c, log: l, mediaList: m, mediaTopList: mt, mediaLatestList: ml, articleList: a, articleTopList: at, articleLatestList: al}
}

//FetchMediaList fetches list of medias from newsapi.org
func (n NewsAPI) FetchMediaList() ([]models.MediaList, []models.MediaList, error) {
	url := n.conf.ChannelEndPoint
	mediaList := n.mediaList
	mediaTopList := n.mediaTopList
	mediaLatestList := n.mediaLatestList
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
				mediaTopList = append(mediaTopList, mediaList.Sources[i])
			}
			if mediaList.Sources[i].SortBysAvailable[j] == "latest" {
				mediaLatestList = append(mediaLatestList, mediaList.Sources[i])
			}
		}
	}
	return mediaTopList, mediaLatestList, nil
}

// FetchArticleList fetchs all articles from newsapi.org
func (n NewsAPI) FetchArticleList() ([]models.ArticleList, []models.ArticleList, error) {
	articleList := n.articleList
	articleTopList := n.articleTopList
	fmt.Println(articleTopList)
	articleLatestList := n.articleLatestList
	for i := 0; i < len(n.mediaTopList); i++ {
		url := n.conf.ArticleEndPoint + "?source=" + n.mediaTopList[i].ID + "&sortBy=top&apiKey=" + n.conf.APIKey
		resp, err := http.Get(url)
		if err != nil {
			n.log.Error(err)
			a := n.articleTopList
			b := n.articleLatestList
			return a, b, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			n.log.Error(err)
			a := n.articleTopList
			b := n.articleLatestList
			return a, b, err
		}
		err = json.Unmarshal(body, &articleList)
		articleTopList = append(articleTopList, articleList)
	}
	for i := 0; i < len(n.mediaLatestList); i++ {
		url := n.conf.ArticleEndPoint + "?source=" + n.mediaLatestList[i].ID + "&sortBy=latest&apiKey=" + n.conf.APIKey
		resp, err := http.Get(url)
		if err != nil {
			n.log.Error(err)
			a := n.articleTopList
			b := n.articleLatestList
			return a, b, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			n.log.Error(err)
			a := n.articleTopList
			b := n.articleLatestList
			return a, b, err
		}
		err = json.Unmarshal(body, &articleList)
		articleLatestList = append(articleLatestList, articleList)
	}
	return articleTopList, articleLatestList, nil
}
