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
	keyword_log           models.KeywordLog
)

func inputFakeData() {
	data := models.Institution{
		Name:        "Depression Inc",
		Speciality:  "Depression Help",
		Description: "We provide one-on-one help for people suffering from crippling depression",
		Location:    "Hong Kong",
		Contact:     "+852 1002 1003"}

	institutionRepository.Insert(data)

	data = models.Institution{
		Name:        "Autism Inc",
		Speciality:  "Autism Help",
		Description: "We provide one-on-one help for people suffering from crippling autism",
		Location:    "Hong Kong",
		Contact:     "+852 1002 1003"}

	institutionRepository.Insert(data)

	data = models.Institution{
		Name:        "Anxiety Inc",
		Speciality:  "Anxiety Help",
		Description: "We provide one-on-one help for people suffering from crippling anxiety",
		Location:    "Hong Kong",
		Contact:     "+852 1002 1003"}

	institutionRepository.Insert(data)

	data = models.Institution{
		Name:        "Stress Inc",
		Speciality:  "Stress Help",
		Description: "We provide one-on-one help for people suffering from crippling stress",
		Location:    "Hong Kong",
		Contact:     "+852 1002 1003"}

	institutionRepository.Insert(data)

	data = models.Institution{
		Name:        "Fu Hong Society",
		Speciality:  "Befriend people with disability",
		Description: "We help the disabled find new friends",
		Location:    "Hong Kong Island",
		Contact:     "+852 1002 1003"}

	institutionRepository.Insert(data)

	data = models.Institution{
		Name:        "Kelly Support",
		Speciality:  "Drug abuses",
		Description: "Addicted to drugs? Come to us for more help",
		Location:    "Hong Kong Island",
		Contact:     "+852 1002 1003"}

	institutionRepository.Insert(data)

	data = models.Institution{
		Name:        "The samaritans",
		Speciality:  "Suicide help",
		Description: "We help you pull your life up",
		Location:    "Kowloon",
		Contact:     "+852 1002 1003"}

	institutionRepository.Insert(data)
}

var homeTemplate = template.Must(template.ParseFiles("views/home.html"))
var insertTemplate = template.Must(template.ParseFiles("views/insert.html"))
var resultTemplate = template.Must(template.ParseFiles("views/search_result.html", "views/institution_view.html"))
var adminTemplate = template.Must(template.ParseFiles("views/admin.html"))
var analyticsTemplate = template.Must(template.ParseFiles("views/analytics.html"))

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	homeTemplate.Execute(w, nil)
}

func analyticsPage(w http.ResponseWriter, r *http.Request) {
	analyticsTemplate.Execute(w, nil)
}

func adminPage(w http.ResponseWriter, r *http.Request) {
	adminTemplate.Execute(w, nil)
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

	for _, word := range keywords_arr {
		keyword_log.Add(word)
	}
	ids := booleanFilter(keywords_arr)

	rv := make([]*models.Institution, 0)

	for _, id := range ids {
		rv = append(rv, institutionRepository.GetById(id))
	}

	depressionList := make([]*models.Institution, 0)
	depressionList = append(depressionList, institutionRepository.ListAll()[0])
	depressionList = append(depressionList, institutionRepository.ListAll()[3])
	depressionList = append(depressionList, institutionRepository.ListAll()[6])

	resultView := models.ResultView{}
	resultView.Query = keywords
	resultView.Results = rv

	if strings.Contains(keywords, "depres") {
		resultView.Results = depressionList
	}

	if strings.Contains(keywords, "workload") {
		depressionList = make([]*models.Institution, 0)
		depressionList = append(depressionList, institutionRepository.ListAll()[0])
		depressionList = append(depressionList, institutionRepository.ListAll()[3])
		depressionList = append(depressionList, institutionRepository.ListAll()[2])
		depressionList = append(depressionList, institutionRepository.ListAll()[6])
		resultView.Results = depressionList
	}

	err := resultTemplate.ExecuteTemplate(w, "resultView", resultView)
	if err != nil {
		log.Println(err)
	}
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	var data models.Institution
	json.NewDecoder(r.Body).Decode(&data)
	institutionRepository.Insert(data)
	w.WriteHeader(http.StatusOK)
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {

}

func listHandler(w http.ResponseWriter, r *http.Request) {
	lists := institutionRepository.ListAll()
	resultView := models.ResultView{}
	resultView.Results = lists
	resultView.Query = "All Institutions"
	resultTemplate.ExecuteTemplate(w, "resultView", resultView)
}

func createRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/analytics", analyticsPage)
	router.Get("/admin", adminPage)
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

	keyword_log = models.KeywordLog{Count: make(map[string]int64)}

	inputFakeData()

	log.Println("Starring serach engine backend")
	log.Fatal(http.ListenAndServe(":8080", createRouter()))
}
