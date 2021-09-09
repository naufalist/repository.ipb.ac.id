package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL       = "https://repository.ipb.ac.id"
	port          = "9000"
	ldap_username = "CHANGE_WITH_YOUR_LDAP_USERNAME"
	ldap_password = "CHANGE_WITH_YOUR_LDAP_PASSWORD"
)

type RepositoryApp struct {
	Client *http.Client
}

type Repository struct {
	Title   string            `json:"title"`
	Date    string            `json:"date"`
	Authors []string          `json:"authors"`
	Files   map[string]string `json:"files"`
}

func (app *RepositoryApp) login() {
	client := app.Client

	loginURL := baseURL + "/ldap-login"

	data := url.Values{
		"username":      {ldap_username},
		"ldap_password": {ldap_password},
	}

	response, err := client.PostForm(loginURL, data)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalln("Error loading HTTP response body. ", err)
	}

	switch document.Find(".alert.alert-danger").First().Text() {
	case "The user name and/or password supplied were not valid.":
		log.Fatalln("Auth failed: invalid username/pasword")
	default:
		log.Println("Auth success.")
	}

}

func (app *RepositoryApp) getRepositoryFile(w http.ResponseWriter, r *http.Request, fileURL string) {
	client := app.Client

	res, err := client.Get(fileURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer res.Body.Close()

	w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", res.Header.Get("Content-Length"))

	_, err = io.Copy(w, res.Body)
	if err != nil {
		http.Error(w, "Remote server error", 503)
		return
	}

	return

}

func getRepositoryDetails(w http.ResponseWriter, req *http.Request) interface{} {
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

		return Repository{
			Title:   title,
			Date:    date,
			Authors: authors,
			Files:   files,
		}

	}

	http.Error(w, req.Method+" isn't allowed.", http.StatusBadRequest)
	return nil

}

func handleApiRepository(w http.ResponseWriter, req *http.Request) {

	repositoryDetails := getRepositoryDetails(w, req)

	w.Header().Set("Content-Type", "application/json")

	result, err := json.Marshal(repositoryDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
	return
}

func handleRoot(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("index.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "POST":
		result := map[string]interface{}{
			"repository": getRepositoryDetails(w, req),
			"status":     true,
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		if err := tmpl.Execute(w, result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func handleGetRepositoryFile(w http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		link := req.FormValue("repository_file")

		res, err := http.Get(link)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Printf("status code error: '%d' '%s'\n", res.StatusCode, res.Status)
		}

		jar, _ := cookiejar.New(nil)

		repository := RepositoryApp{
			Client: &http.Client{Jar: jar},
		}

		repository.login()
		repository.getRepositoryFile(w, req, link)
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
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/get-repository-file", handleGetRepositoryFile)
	startGoServer()
}
