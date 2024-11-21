package search

type Searcher interface {
	Search(data any, target int) bool
}
