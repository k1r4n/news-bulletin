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
	mediaList         models.MediaList
	mediaTopList      []models.MediaList
	mediaLatestList   []models.MediaList
	articleTopList    []models.ArticleList
	articleLatestList []models.ArticleList
}

// NewsAPIInstant create and object of NewsAPI service
func NewsAPIInstant(c *conf.Vars, l *logrus.Logger, m models.MediaList, mt []models.MediaList, ml []models.MediaList, at []models.ArticleList, al []models.ArticleList) NewsAPI {
	return NewsAPI{conf: c, log: l, mediaList: m, mediaTopList: mt, mediaLatestList: ml, articleTopList: at, articleLatestList: al}
}

//FetchMediaList fetches list of medias from newsapi.org
func (n NewsAPI) FetchMediaList() ([]models.MediaList, []models.MediaList, error) {
	url := n.conf.ChannelEndPoint
	mediaList := n.mediaList
	mediaTopList := n.mediaTopList
	mediaLatestList := n.mediaLatestList
	resp, err := http.Get(url)
	if err != nil {
		n.log.Error("FetchMediaList: Error getting api result. Error code:", err)
		return mediaTopList, mediaLatestList, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		n.log.Error("FetchMediaList: Read error. Error Code:", err)
	}
	err = json.Unmarshal(body, &mediaList)
	fmt.Println(mediaList)
	return n.mediaTopList, n.mediaLatestList, nil
}
