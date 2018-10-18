package inmemory_store

type InvertedIndex struct {
	inverted_index map[string][]int64
}

func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{inverted_index: make(map[string][]int64)}
}

func (i *InvertedIndex) GetDocumentsContainingKey(keyword string) []int64 {
	return i.inverted_index[keyword]
}

func (i *InvertedIndex) Insert(key string, id int64) {
	arr, ok := i.inverted_index[key]
	if !ok {
		arr = make([]int64, 0)
	}

	arr = append(arr, id)
	i.inverted_index[key] = arr
}
