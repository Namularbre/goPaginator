package paginator

import linq "github.com/Namularbre/goLinq"

const DefaultNumberOfElemByPage int = 50

// Paginator is a struct to help you paginate things
type Paginator[T any] struct {
	Content            []T `json:"content"`
	ContentLen         int `json:"contentLen"`
	CurrentPage        int `json:"currentPage"`
	PageCount          int `json:"pageCount"`
	NumberOfElemByPage int `json:"numberOfElemByPage"`
}

// Page represent a page of the paginator
type Page[T any] struct {
	Content  []T `json:"content"`
	Previous int `json:"previous"`
	Next     int `json:"next"`
}

// NewDefaultPaginator create a new paginator with default settings
func NewDefaultPaginator[T any](content []T, currentPage int) *Paginator[T] {
	contentLen := len(content)
	return &Paginator[T]{
		Content:     content,
		ContentLen:  contentLen,
		CurrentPage: currentPage,
		PageCount:   contentLen / DefaultNumberOfElemByPage,
	}
}

// NewPaginator create a paginator
func NewPaginator[T any](content []T, currentPage int, numberOfElemByPage int) *Paginator[T] {
	contentLen := len(content)
	return &Paginator[T]{
		Content:            content,
		ContentLen:         contentLen,
		CurrentPage:        currentPage,
		NumberOfElemByPage: numberOfElemByPage,
		PageCount:          contentLen / numberOfElemByPage,
	}
}

// GetPage get the page number x of the paginator
func (p *Paginator[T]) GetPage(page int) *Page[T] {
	firstElemIndex := uint(page * p.NumberOfElemByPage)

	if firstElemIndex > uint(p.ContentLen) {
		return nil
	}

	skipped := linq.Skip(p.Content, firstElemIndex)
	return &Page[T]{
		Content:  linq.Take(skipped, uint(p.NumberOfElemByPage)),
		Previous: page - 1,
		Next:     page + 1,
	}
}
