package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL = "https://repository.ipb.ac.id"
	port    = "9000"
)

type Repository struct {
	Title   string            `json:"title"`
	Date    string            `json:"date"`
	Authors []string          `json:"authors"`
	Files   map[string]string `json:"files"`
}

func handleApiRepository(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method == "POST" {
		link := req.FormValue("link")

		res, err := http.Get(link)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Printf("status code error: '%d' '%s'\n", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Println(err.Error())
		}

		title := doc.Find(".page-header.first-page-header").First().Text()

		date := doc.Find(".simple-item-view-date.word-break.item-page-field-wrapper.table").Contents().Text()
		date = strings.Trim(date, "\nDate")

		authors := []string{}
		doc.Find(".simple-item-view-authors").Contents().Each(func(i int, s *goquery.Selection) {
			if goquery.NodeName(s) == "div" {
				authors = append(authors, s.Text())
			}
		})

		files := make(map[string]string)
		doc.Find(".item-page-field-wrapper.table.word-break").Contents().Each(func(i int, s *goquery.Selection) {
			if goquery.NodeName(s) == "div" {
				if href, exist := s.Children().Attr("href"); exist {
					link := s.Text()
					link = strings.Trim(link, "\n ")
					files[link] = baseURL + href
				}
			}
		})

		repository := Repository{
			Title:   title,
			Date:    date,
			Authors: authors,
			Files:   files,
		}

		result, err := json.Marshal(repository)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return

	}

	http.Error(w, req.Method+" isn't allowed.", http.StatusBadRequest)
}

func startGoServer() {
	log.Printf("web server started at :%s\n", port)

	server := new(http.Server)
	server.Addr = ":" + port

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	http.HandleFunc("/api/repository", handleApiRepository)
	startGoServer()
}
