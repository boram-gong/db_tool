package es

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
)

var ctx = context.Background()

type QueryGroup struct {
	Term  map[string]string
	Match map[string]string
}

func NewEs(url string) (*elastic.Client, error) {
	ClientEs, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	_, _, err = ClientEs.Ping(url).Do(ctx)
	if err != nil {
		return nil, err
	}
	return ClientEs, nil
}

func CreateIndex(cli *elastic.Client, index string, mapping string) error {
	exists, err := cli.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		// 如果不存在，就创建
		createIndex, err := cli.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			return errors.New("Not Acknowledged")
		}
	}
	return nil
}

func Insert(cli *elastic.Client, index string, id string, body interface{}) error {
	_, err := cli.Index().Index(index).BodyJson(body).Id(id).Do(ctx)
	return err
}

func BulkInsert(cli *elastic.Client, index string, ids []string, data []interface{}) error {
	if len(ids) != len(data) {
		return errors.New("ids.len != data.len")
	}
	bulkRequest := cli.Bulk()
	for i, v := range data {
		req := elastic.NewBulkIndexRequest().Index(index).Id(ids[i]).Doc(v)
		bulkRequest = bulkRequest.Add(req)
	}
	_, err := bulkRequest.Do(ctx)
	return err
}

func MustQuery(cli *elastic.Client, index string, group QueryGroup, onlyTerm bool) ([]*elastic.SearchHit, error) {
	var query []elastic.Query
	if len(group.Term) != 0 {
		for k, v := range group.Term {
			query = append(query, elastic.NewTermQuery(k, v))
		}
	}
	if len(group.Match) != 0 && !onlyTerm {
		for k, v := range group.Match {
			query = append(query, elastic.NewMatchQuery(k, v))
		}
	}
	if len(query) != 0 {
		mustQuery := elastic.NewBoolQuery().Must(query...)
		searchResult, err := cli.Search().
			Index(index).
			Query(mustQuery).
			Pretty(true).
			Do(ctx)
		if err != nil {
			return nil, err
		}
		if searchResult.Hits.TotalHits.Value > 0 {
			return searchResult.Hits.Hits, nil
		}
	}
	return nil, nil
}

func MultiMustQuery(cli *elastic.Client, index string, groups []QueryGroup, onlyTerm bool) ([][]*elastic.SearchHit, error) {
	var (
		mQuery = cli.MultiSearch()
		result [][]*elastic.SearchHit
	)
	for _, group := range groups {
		var query []elastic.Query
		if len(group.Term) != 0 {
			for k, v := range group.Term {
				query = append(query, elastic.NewTermQuery(k, v))
			}
		}
		if len(group.Match) != 0 && !onlyTerm {
			for k, v := range group.Match {
				query = append(query, elastic.NewMatchQuery(k, v))
			}
		}
		if len(query) != 0 {
			mustQuery := elastic.NewBoolQuery().Must(query...)
			searchRequest := elastic.NewSearchRequest().Index(index).Query(mustQuery)
			mQuery.Add(searchRequest)
		}
	}

	MSearchResult, err := mQuery.Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}
	for _, searchResult := range MSearchResult.Responses {
		result = append(result, searchResult.Hits.Hits)
	}
	return result, nil
}
