package paginator

import (
	"slices"
	"testing"
)

type userForTesting struct {
	Id       int
	Username string
	Age      uint
}

func PrepareTestData() []userForTesting {
	return []userForTesting{
		{
			Id:       0,
			Username: "Dupond",
			Age:      69,
		},
		{
			Id:       1,
			Username: "Dupont",
			Age:      96,
		},
		{
			Id:       2,
			Username: "Alice",
			Age:      47,
		},
		{
			Id:       3,
			Username: "Bob",
			Age:      23,
		},
		{
			Id:       4,
			Username: "Caly",
			Age:      14,
		},
	}
}

func TestNewPaginator(t *testing.T) {
	input := PrepareTestData()
	assertCurrentPage := 0
	assertNumberOfElemByPage := 2
	assertPageCount := (len(input) + assertNumberOfElemByPage - 1) / assertNumberOfElemByPage
	assert := Paginator[userForTesting]{
		Content:            input,
		CurrentPage:        assertCurrentPage,
		NumberOfElemByPage: assertNumberOfElemByPage,
		PageCount:          assertPageCount,
		ContentLen:         len(input),
	}

	res := NewPaginator[userForTesting](input, assertCurrentPage, assertNumberOfElemByPage)

	if res.NumberOfElemByPage != assert.NumberOfElemByPage || res.PageCount != assert.PageCount || !slices.Equal(res.Content, assert.Content) || res.ContentLen != assert.ContentLen || res.CurrentPage != assertCurrentPage {
		t.Fatalf("err: the paginator is not built properly, time to debug!")
	}
}

func PreparePaginator() *Paginator[userForTesting] {
	return NewPaginator(PrepareTestData(), 0, 2)
}

func TestPaginator_GetPageZero(t *testing.T) {
	paginator := PreparePaginator()
	assertPage := Page[userForTesting]{
		Content: []userForTesting{
			{
				Id:       0,
				Username: "Dupond",
				Age:      69,
			},
			{
				Id:       1,
				Username: "Dupont",
				Age:      96,
			},
		},
		Previous: -1,
		Next:     1,
	}
	page := paginator.GetPage(0)

	if page.Next != assertPage.Next || page.Previous != assertPage.Previous || !slices.Equal(page.Content, assertPage.Content) {
		t.Fatalf("error with getPage(0), time to debug.")
	}
}

func TestPaginator_GetPageNil(t *testing.T) {
	paginator := NewPaginator(PrepareTestData(), 0, 2)
	page := paginator.GetPage(999)

	if page != nil {
		t.Fatalf("error, page should be nil, for %v", page)
	}
}
