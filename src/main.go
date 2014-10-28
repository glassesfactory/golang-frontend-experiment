package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yosssi/ace"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {
	goji.Get("/", Root)
	goji.Get("/news", News)
	goji.Get("/assets/:dir/:file", Static)
	goji.Serve()
}

// Render レンダリングする
func Render(tmplName string, w http.ResponseWriter, r *http.Request) {
	current, err := os.Getwd()

	tpl, err := ace.Load("partials/base", tmplName, &ace.Options{
		BaseDir:       current,
		DynamicReload: true,
	})

	if err != nil {
		log.Print("some error...\n", err)
	}

	if err := tpl.Execute(w, nil); err != nil {
		log.Print("some error...\n", err)
	}
}

// Root トップページ
func Root(c web.C, w http.ResponseWriter, r *http.Request) {
	Render("Index", w, r)
}

// News トップページ
func News(c web.C, w http.ResponseWriter, r *http.Request) {
	Render("news/index", w, r)
}

// Static スタティックファイル
func Static(c web.C, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
