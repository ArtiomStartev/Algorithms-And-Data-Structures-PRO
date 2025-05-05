package search

type Searcher interface {
	Search(data any, x int) bool
}
