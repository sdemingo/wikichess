package http

import (
	"model/games"
	"net/http"
	"text/template"
)

var routes map[string]bool

var directRoutes map[string]bool

func init() {
	routes = make(map[string]bool)

	directRoutes = map[string]bool{
		"/":          true,
		"/logout":    true,
		"/admin":     true,
		"/games/new": true,
		"/games/add": true,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		routes["/"] = true
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		tmpl := template.Must(template.ParseFiles(baseTmpl))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		routes["/help"] = true
		AppHandler(w, r, Help)
	})
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		routes["/admin"] = true
		AppHandler(w, r, Admin)
	})
	http.HandleFunc("/games/new", func(w http.ResponseWriter, r *http.Request) {
		routes["/games/new"] = true
		AppHandler(w, r, games.NewGameHandler)
	})
	http.HandleFunc("/games/add", func(w http.ResponseWriter, r *http.Request) {
		routes["/games/add"] = true
		AppHandler(w, r, games.SaveGameHandler)
	})
}
