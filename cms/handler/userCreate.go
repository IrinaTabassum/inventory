package handler

import (
	// "fmt"
	// "log"
	"net/http"
	// "strings"

	// validation "github.com/go-ozzo/ozzo-validation/v4"
	// "github.com/justinas/nosurf"
)

type UserForm struct {
	// ListUser  []storage.User
	// User      storage.User
	FormError map[string]error
	CSRFToken string
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// uf := storage.UserFilter{
	// 	Limit: 100,
	// }
	// listUser, err := h.storage.ListUser(uf)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "internal server error", http.StatusInternalServerError)
	// }

	// h.pareseCreateUserTemplate(w, UserForm{
	// 	ListUser:  listUser,
	// 	CSRFToken: nosurf.Token(r),
	// })
}

func (h Handler) StoreUser(w http.ResponseWriter, r *http.Request) {
	// if err := r.ParseForm(); err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "internal server error", http.StatusInternalServerError)
	// }
	// form := UserForm{}
	// user := storage.User{}
	// if err := h.decoder.Decode(&user, r.PostForm); err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "internal server error", http.StatusInternalServerError)
	// }
	// log.Printf("%+v", user)
	// return
	// form.User = user
	// if err := user.Validate(); err != nil {
	// 	if vErr, ok := err.(validation.Errors); ok {
	// 		for key, val := range vErr {
	// 			form.FormError[strings.Title(key)] = val
	// 		}
	// 	}
	// 	h.pareseCreateUserTemplate(w, form)
	// 	return
	// }

	// newUser, err := h.storage.CreateUser(user)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "internal server error", http.StatusInternalServerError)
	// }

	// http.Redirect(w, r, fmt.Sprintf("/users/%v/edit", newUser.ID), http.StatusSeeOther)
}

func (h Handler) pareseCreateUserTemplate(w http.ResponseWriter, data any) {
	// t := h.Templates.Lookup("create-user.html")
	// if t == nil {
	// 	log.Println("unable to lookup create-user template")
	// 	h.Error(w, "internal server error", http.StatusInternalServerError)
	// }

	// if err := t.Execute(w, data); err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "internal server error", http.StatusInternalServerError)
	// }
}

/* INSERT INTO users (first_name, last_name, username, email, password, status, created_at, updated_at, deleted_at, phone) VALUES ('Mark', 'Taylor', 'boropafigq', 'saluroraqe@mailinator.com', 'Pa$$w0rd!', true, '2023-02-03 17:35:01.577029', '2023-02-03 17:35:01.577029', null, '02840402') ON CONFLICT(username) DO UPDATE SET first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, username = EXCLUDED.username, email = EXCLUDED.email, password=EXCLUDED.password, status = EXCLUDED.status, updated_at = EXCLUDED.updated_at, phone = EXCLUDED.phone ; */