package models

import "sort"

type KeywordLog struct {
	Count map[string]int64
}

func (k *KeywordLog) Add(word string) {
	count := k.Count[word]
	count = count + 1
	k.Count[word] = count
}

func (k *KeywordLog) GetTopTen() []string {

	var keys []string
	for k := range k.Count {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return k.Count[keys[i]] < k.Count[keys[j]]
	})

	if len(keys) < 10 {
		return keys
	}

	return keys[:10]
}
