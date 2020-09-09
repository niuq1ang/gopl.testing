package api

import (
	"encoding/json"
	"log"

	"github.com/bangwork/mygo/2020/0324/config"
)

type SearchQueryInfo struct {
	Limit           int64                    `json:"limit"`
	Offset          int64                    `json:"offset"`
	Filters         []map[string]interface{} `json:"filters,omitempty"`
	FieldProjection []string                 `json:"fieldProjection"`
	Orders          []map[string]interface{} `json:"orders,omitempty"`
}

type QueryResult struct {
	BaseRes
	Data DataRes `json:"data"`
}

type DataRes struct {
	Total    int64             `json:"total"`
	Offset   int64             `json:"offset"`
	Limit    int64             `json:"limit"`
	DataList []json.RawMessage `json:"dataList"`
}

func (searchQuery SearchQueryInfo) queryAllPages(apiName string) []json.RawMessage {
	var results []json.RawMessage
	searchQuery.Limit = 100
	searchQuery.Offset = 0

	result := query(searchQuery, apiName)
	if result.ErrorCode != 0 {
		log.Printf("Call OpenAPI Error, errCode=[%d], errMsg=%s", result.ErrorCode, result.ErrorMessage)
		return results
	}
	results = append(results, result.Data.DataList...)
	pages := result.Data.Total / searchQuery.Limit

	if pages > 0 {
		for i := int64(0); i < pages-1; i++ {
			searchQuery.Offset = 100 * (1 + i)
			result := query(searchQuery, apiName)
			if result.ErrorCode != 0 {
				log.Printf("Call OpenAPI Error, errCode=[%d], errMsg=%s", result.ErrorCode, result.ErrorMessage)
				return results
			}
			results = append(results, result.Data.DataList...)
		}

		rest := result.Data.Total - pages*searchQuery.Limit
		if rest > 0 && pages > 0 {
			searchQuery.Offset = pages * searchQuery.Limit
			result := query(searchQuery, apiName)
			if result.ErrorCode != 0 {
				log.Printf("Call OpenAPI Error, errCode=[%d], errMsg=%s", result.ErrorCode, result.ErrorMessage)
				return results
			}
			results = append(results, result.Data.DataList...)
		}
	}
	return results
}

func query(searchQuery SearchQueryInfo, apiName string) QueryResult {
	var results QueryResult
	client := NewClient(config.Config.URL)
	token := GetToken()
	paramMap := map[string]interface{}{
		"corpAccessToken":   token.CorpAccessToken,
		"corpId":            token.CorpID,
		"currentOpenUserId": config.Config.AppUserId,
		"data": map[string]interface{}{
			"dataObjectApiName": apiName,
			"search_query_info": searchQuery,
		},
	}
	_, err := client.Post("/cgi/crm/v2/data/query", paramMap, &results)
	if err != nil {
		log.Printf("query data err: %+v", err)
	}
	return results
}

func queryDefined(searchQuery SearchQueryInfo, apiName string) QueryResult {
	var results QueryResult
	client := NewClient(config.Config.URL)
	token := GetToken()
	paramMap := map[string]interface{}{
		"corpAccessToken":   token.CorpAccessToken,
		"corpId":            token.CorpID,
		"currentOpenUserId": config.Config.AppUserId,
		"data": map[string]interface{}{
			"dataObjectApiName": apiName,
			"search_query_info": searchQuery,
		},
	}
	_, err := client.Post("/cgi/crm/custom/data/query", paramMap, &results)
	if err != nil {
		log.Printf("query data err: %+v", err)
	}
	return results
}
