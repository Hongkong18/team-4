package models

type InvertedIndex interface {
	GetDocumentsContainingKey(keyword string) []int64
	Insert(keyword string, id int64)
}
