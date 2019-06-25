package handler

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"ClickPhonebook/db"
	"ClickPhonebook/schema"
	"ClickPhonebook/util"
	"strconv"
	"fmt"
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
		util.ResponseError(w,500,"Can't add contact")
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) AddFormPhone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item  := &schema.Phone{}

	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ResponseError(w,500,"Bad id")
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
		util.ResponseError(w,500,"Bad id")
	}

	// в целям упрощения примера пропущена валидация
	err = 	db.AddPhone(id,r.FormValue("PhoneNumber"))
	if err != nil {
		util.ResponseError(w,500,"Can't add Phone")
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) Edit (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ResponseOk(w,"Bad id")
	}
	contacts, phones, err := db.SelectItem(id)

	err = h.Tmpl.ExecuteTemplate(w, "edit.html", &struct {
		Contacts schema.Contact
		Phones []string
	}{
		contacts[id],
		phones[id],
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	r.ParseForm()
	fmt.Println(r.FormValue("phonenumber"))
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ResponseOk(w,"Bad id")
	}
	var contact schema.Contact
	contact.Id = id
	contact.FirstName = r.FormValue("firstname")
	contact.LastName  = r.FormValue("lastname")
	//var phonenumbers map[string][]string
	phonenumbers := r.Form["phonenumber"]
	err = db.Update(contact,phonenumbers)
	if err != nil  {
		util.ResponseError(w,404,"Can't update contact")
	}

	http.Redirect(w, r, "/", http.StatusFound)

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ResponseOk(w,"Bad id")
	}

	err = db.Delete(id)
	if err != nil {
		util.ResponseOk(w,"Error delete contact")
	}

	//w.Header().Set("Content-type", "application/json")
	resp := `{"affected": ` + strconv.Itoa(int(id)) + `}`
	//w.Write([]byte(resp))
	util.ResponseOk(w, resp)

}

