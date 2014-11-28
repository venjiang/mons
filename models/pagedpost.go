package models

type PagedPost struct {
	Posts   []*Post
	Records int
	Pages   int
}
