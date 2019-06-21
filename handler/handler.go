package handler

import (
	"html/template"
	"net/http"
	"ClicksPhonebook/db"
	"ClicksPhonebook/schema"
	"ClicksPhonebook/util"
)

type Handler struct {
	Tmpl *template.Template
}

func NewHandler() *Handler  {
	return &Handler{
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	contacts, phones, err := db.List()

	err = h.Tmpl.ExecuteTemplate(w, "index.html", struct {
		Contacts map[int]schema.Contact
		Phones map[int][]string
	}{
		contacts,
		phones,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) AddForm(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	// в целям упрощения примера пропущена валидация
	err := 	db.AddContact(r.FormValue("title"),	r.FormValue("description"))
	if err != nil {
		util.ResponseOk(w,"Error add contact")
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
