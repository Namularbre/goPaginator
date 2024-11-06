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
	assertPageCount := len(input) / assertNumberOfElemByPage
	assert := Paginator[userForTesting]{
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
		},
		CurrentPage:        assertCurrentPage,
		NumberOfElemByPage: assertNumberOfElemByPage,
		PageCount:          assertPageCount,
		ContentLen:         len(input),
	}

	res := NewPaginator[userForTesting](input, assertCurrentPage, assertNumberOfElemByPage)

	if res.NumberOfElemByPage != assert.NumberOfElemByPage || res.PageCount != assert.PageCount || !slices.Equal(res.Content, assert.Content) || res.ContentLen != assert.ContentLen || res.CurrentPage != assertCurrentPage {
		t.Fatalf("err: the paginator is not build properly, time to debug !")
	}
}
