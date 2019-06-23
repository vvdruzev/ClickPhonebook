package handler

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"ClickPhonebook/db"
	"ClickPhonebook/schema"
	"ClickPhonebook/util"
	"strconv"
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
	err := 	db.AddContact(r.FormValue("firstname"),	r.FormValue("lastname"))
	if err != nil {
		util.ResponseOk(w,"Error add contact")
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) AddFormPhone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item  := &schema.Phone{}

	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ResponseOk(w,"Bad id")
	}
	item.Id = id
	err = h.Tmpl.ExecuteTemplate(w, "addphone.html", item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddPhone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ResponseOk(w,"Bad id")
	}

	// в целям упрощения примера пропущена валидация
	err = 	db.AddPhone(id,r.FormValue("PhoneNumber"))
	if err != nil {
		util.ResponseOk(w,"Error add contact")
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
