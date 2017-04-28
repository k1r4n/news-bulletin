package services

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/news-bulletin-cron/conf"
	"github.com/news-bulletin-cron/models"
	elastic "gopkg.in/olivere/elastic.v5"
)

// Elastic Object
type Elastic struct {
	conf              *conf.Vars
	log               *logrus.Logger
	elasticClient     *elastic.Client
	ctx               context.Context
	articleTopList    []models.ArticleList
	articleLatestList []models.ArticleList
	mediaTopList      []models.MediaList
	mediaLatestList   []models.MediaList
}

// NewElasticInstant creates an object of Elastic
func NewElasticInstant(ctx context.Context, c *conf.Vars, l *logrus.Logger, e *elastic.Client, at []models.ArticleList, al []models.ArticleList, mt []models.MediaList, ml []models.MediaList) Elastic {
	return Elastic{conf: c, log: l, elasticClient: e, ctx: ctx, articleTopList: at, articleLatestList: al, mediaTopList: mt, mediaLatestList: ml}
}

func (e Elastic) UpdateDatabase() {
	exist, err := e.elasticClient.IndexExists(e.conf.Index).Do(e.ctx)
	if err != nil {
		e.log.Error(err)
		return
	}
	if exist {
		_, err = e.elasticClient.DeleteIndex(e.conf.Index).Do(e.ctx)
		if err != nil {
			e.log.Error(err)
			return
		}
	}
	createIndex, err := e.elasticClient.CreateIndex(e.conf.Index).Do(e.ctx)
	if err != nil {
		e.log.Error(err)
		return
	}
	if createIndex.Acknowledged {
		bulkRequest := e.elasticClient.Bulk()
		for i := 0; i < len(e.mediaTopList); i++ {
			doc, err := json.Marshal(e.mediaTopList[i])
			if err != nil {
				e.log.Error(err)
				return
			}
			req := e.elasticClient.NewBulkIndexRequest().Index(e.conf.Index).Type(e.conf.ChannelTypeTop).Id(strconv.Itoa(i + 1)).Doc(doc)
		}
	}
}
