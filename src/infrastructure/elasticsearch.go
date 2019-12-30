package infrastructure

import (
	"interfaces/repository"
)

type ElasticsearchHandler struct {
}

func NewElasticsearchHandler() repository.LogSearchEngineHandler {
	handler := new(ElasticsearchHandler)
	return handler
}

func (elasticsearchHandler *ElasticsearchHandler) Query(query string) (result []string, err error) {
	result = []string{"hoge","fuga"}
    return
}
