package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"rspweb/service"
	"strconv"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

type Player struct {
	Name string
}

var player Player

func Index(w http.ResponseWriter, r *http.Request) {
	restartValues()
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValues()
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error al pársear el formulario", http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("playerName")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}
	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("playerChoice"))
	result := service.PlayRound(playerChoice)
	fmt.Println(result)
	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	write, err := w.Write(out)
	if err != nil {
		return
	}
	fmt.Println(write)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValues()
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error al renderizar la plantilla:", err)
		return
	}
}

func restartValues() {
	player = Player{}
}
