package services

import (
	"context"
	"fmt"

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

// UpdateDatabase update datebase by removing current news by latest news
func (e Elastic) UpdateDatabase() error {
	exist, err := e.elasticClient.IndexExists(e.conf.Index).Do(e.ctx)
	if err != nil {
		e.log.Error(err)
		return err
	}
	if exist {
		fmt.Println("Deleting existing index \n ")
		_, err = e.elasticClient.DeleteIndex(e.conf.Index).Do(e.ctx)
		if err != nil {
			e.log.Error(err)
			return err
		}
	}
	createIndex, err := e.elasticClient.CreateIndex(e.conf.Index).Do(e.ctx)
	if err != nil {
		e.log.Error(err)
		return err
	}
	if createIndex.Acknowledged {
		fmt.Println("Inserting data \n ")
		for i := 0; i < len(e.mediaTopList); i++ {
			put, err := e.elasticClient.Index().Index(e.conf.Index).Type(e.conf.ChannelTypeTop).Id(e.mediaTopList[i].ID).BodyJson(e.mediaTopList[i]).Do(e.ctx)
			if err != nil {
				e.log.Error(err)
				return err
			}
			fmt.Printf("Inserted Document %v to Index %s, Type %s \n ", put, e.conf.Index, e.conf.ChannelTypeTop)
		}
		for i := 0; i < len(e.mediaLatestList); i++ {
			put, err := e.elasticClient.Index().Index(e.conf.Index).Type(e.conf.ChannelTypeLatest).Id(e.mediaLatestList[i].ID).BodyJson(e.mediaLatestList[i]).Do(e.ctx)
			if err != nil {
				e.log.Error(err)
				return err
			}
			fmt.Printf("Inserted Document %v to Index %s, Type %s \n ", put, e.conf.Index, e.conf.ChannelTypeLatest)
		}
		for i := 0; i < len(e.articleTopList); i++ {
			put, err := e.elasticClient.Index().Index(e.conf.Index).Type(e.conf.ArticleTypeTop).Id(e.articleTopList[i].Source).BodyJson(e.articleTopList[i]).Do(e.ctx)
			if err != nil {
				e.log.Error(err)
				return err
			}
			fmt.Printf("Inserted Document %v to Index %s, Type %s \n ", put, e.conf.Index, e.conf.ArticleTypeTop)
		}
		for i := 0; i < len(e.articleLatestList); i++ {
			put, err := e.elasticClient.Index().Index(e.conf.Index).Type(e.conf.ArticleTypeLatest).Id(e.articleLatestList[i].Source).BodyJson(e.articleLatestList[i]).Do(e.ctx)
			if err != nil {
				e.log.Error(err)
				return err
			}
			fmt.Printf("Inserted Document %v to Index %s, Type %s \n", put, e.conf.Index, e.conf.ArticleTypeLatest)
		}
	}
	return nil
}
