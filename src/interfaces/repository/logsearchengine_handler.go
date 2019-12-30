package repository

type LogSearchEngineHandler interface {
    Query(string) ([]string, error)
}