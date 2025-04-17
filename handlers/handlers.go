package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
	"text/template"
)

type Player struct {
	Name string
}

var player Player

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
			return
		}

		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoise, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoise)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "about.html", nil)
}

func restartValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
