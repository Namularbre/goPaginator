package paginator

const DefaultNumberOfElemByPage int = 50

// Paginator is a struct to help you paginate things
type Paginator[T any] struct {
	Content            []T `json:"content"`
	ContentLen         int `json:"contentLen"`
	CurrentPage        int `json:"currentPage"`
	PageCount          int `json:"pageCount"`
	NumberOfElemByPage int `json:"numberOfElemByPage"`
}

// Page represents a page of the paginator
type Page[T any] struct {
	Content  []T `json:"content"`
	Previous int `json:"previous"`
	Next     int `json:"next"`
}

// NewDefaultPaginator creates a new paginator with default settings
func NewDefaultPaginator[T any](content []T, currentPage int) *Paginator[T] {
	contentLen := len(content)
	pageCount := contentLen / DefaultNumberOfElemByPage
	if contentLen%DefaultNumberOfElemByPage != 0 {
		pageCount++
	}
	return &Paginator[T]{
		Content:            content,
		ContentLen:         contentLen,
		CurrentPage:        currentPage,
		PageCount:          pageCount,
		NumberOfElemByPage: DefaultNumberOfElemByPage,
	}
}

// NewPaginator creates a paginator
func NewPaginator[T any](content []T, currentPage int, numberOfElemByPage int) *Paginator[T] {
	contentLen := len(content)
	pageCount := contentLen / numberOfElemByPage
	if contentLen%numberOfElemByPage != 0 {
		pageCount++
	}
	return &Paginator[T]{
		Content:            content,
		ContentLen:         contentLen,
		CurrentPage:        currentPage,
		NumberOfElemByPage: numberOfElemByPage,
		PageCount:          pageCount,
	}
}

// GetPage gets the page number x of the paginator
func (p *Paginator[T]) GetPage(page int) *Page[T] {
	if page < 0 || page >= p.PageCount {
		return nil
	}

	firstElemIndex := page * p.NumberOfElemByPage
	lastElemIndex := firstElemIndex + p.NumberOfElemByPage

	if firstElemIndex >= p.ContentLen {
		return nil
	}

	var content []T
	if lastElemIndex > p.ContentLen {
		content = p.Content[firstElemIndex:]
	} else {
		content = p.Content[firstElemIndex:lastElemIndex]
	}

	return &Page[T]{
		Content:  content,
		Previous: page - 1,
		Next:     page + 1,
	}
}
