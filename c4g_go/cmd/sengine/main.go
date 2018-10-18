package main

import (
	"c4g_go/inmemory_store"
	"c4g_go/models"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

var (
	institutionRepository models.InstitutionInterface
	invertedIndex         models.InvertedIndex
)

func inputFakeData() {

}

var homeTemplate = template.Must(template.ParseFiles("views/home.html"))
var insertTemplate = template.Must(template.ParseFiles("views/insert.html"))

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	homeTemplate.Execute(w, nil)
}

func insertPage(w http.ResponseWriter, r *http.Request) {
	insertTemplate.Execute(w, nil)
}

func intersect(list1, list2 []int64) (answer []int64) {
	i := 0
	j := 0

	for i != len(list1) && j != len(list2) {
		if list1[i] == list2[j] {
			answer = append(answer, list1[i])
			i++
			j++
		} else if list1[i] < list2[j] {
			i++
		} else {
			j++
		}
	}
	return
}

func booleanFilter(query []string) (docIDs []int64) {
	if len(query) == 0 {
		return
	}

	wordDoc := make(map[string][]int64)
	for _, word := range query {
		wordDoc[word] = invertedIndex.GetDocumentsContainingKey(word)
	}

	sort.Slice(query, func(i, j int) bool {
		return len(wordDoc[query[i]]) < len(wordDoc[query[j]])
	})

	docIDs = append(docIDs, wordDoc[query[0]]...)
	for i, word := range query {
		if i == 0 {
			continue
		}
		docIDs = intersect(docIDs, wordDoc[word])
	}

	return
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	keywords := r.URL.Query()["keyword"][0]
	keywords_arr := strings.Split(keywords, " ")

	ids := booleanFilter(keywords_arr)

	rv := make([]*models.Institution, 0)

	for _, id := range ids {
		rv = append(rv, institutionRepository.GetById(id))
	}

	json.NewEncoder(w).Encode(rv)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	var data models.Institution
	json.NewDecoder(r.Body).Decode(&data)
	institutionRepository.Insert(&data)
	w.WriteHeader(http.StatusOK)
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {

}

func listHandler(w http.ResponseWriter, r *http.Request) {
	lists := institutionRepository.ListAll()
	json.NewEncoder(w).Encode(lists)
	w.WriteHeader(http.StatusOK)
}

func createRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/insertPage", insertPage)
	router.Get("/home", helloWorldHandler)
	router.Get("/search", searchHandler)
	router.Post("/insert", insertHandler)
	router.Get("/listInstitutions", listHandler)
	router.Post("/verify", verifyHandler)
	return router
}

func main() {

	invertedIndex = inmemory_store.NewInvertedIndex()
	institutionRepository = inmemory_store.NewInstitutionService(invertedIndex)

	log.Println("Starring serach engine backend")
	log.Fatal(http.ListenAndServe(":8080", createRouter()))
}
