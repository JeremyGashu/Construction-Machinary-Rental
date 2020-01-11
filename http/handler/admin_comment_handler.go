package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"time"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// AdminCommentHandler handles Admin handler admin requests
type AdminCommentHandler struct {
	tmpl       *template.Template
	CommentSrv admin.CommentService
}

// NewAdminCommentHandler initializes and returns new AdminCommentHandler
func NewAdminCommentHandler(T *template.Template, CS admin.CommentService) *AdminCommentHandler {
	return &AdminCommentHandler{tmpl: T, CommentSrv: CS}
}

// AdminComments handle requests on route /admin/Comments
func (ach *AdminCommentHandler) AdminComments(w http.ResponseWriter, r *http.Request) {
	Comments, err := ach.CommentSrv.Comments()

	if err != nil {
		panic(err)
	}
	ach.tmpl.ExecuteTemplate(w, "admin.comment.layout", Comments)
}

// func (ach *AdminCommentHandler) APIAdminComments(w http.ResponseWriter, r *http.Request) {
// 	Comments, err := ach.CommentSrv.Comments()
// 	if err != nil {
// 		panic(err)
// 	}
// 	json.NewEncoder(w).Encode(Comments)
// }

// AdminCommentsNew hanlde requests on route /admin/Comments/new
func (ach *AdminCommentHandler) AdminCommentsNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		ctg := entity.Comment{}
		ctg.UserName = r.FormValue("username")
		ctg.Message = r.FormValue("message")
		ctg.Email = r.FormValue("email")
		a, b, c := time.Now().Date()
		d := b.String()
		e := time.Now().Hour()
		f := time.Now().Minute()
		ctg.PlacedAt = strconv.Itoa(e) + ":" + strconv.Itoa(f) + "  " + strconv.Itoa(c) + "/" + d + "/" + strconv.Itoa(a)
		fmt.Println(ctg.UserName, ctg.Message, ctg.Email, ctg.PlacedAt)
		err := ach.CommentSrv.StoreComment(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/comment", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "admin.comment.new.layout", nil)

	}
}

// AdminCommentsUpdate handle requests on /admin/Comments/update
func (ach *AdminCommentHandler) AdminCommentsUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := ach.CommentSrv.Comment(id)
		if err != nil {
			panic(err)
		}

		ach.tmpl.ExecuteTemplate(w, "admin.comment.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.Comment{}
		ctg.UserName = r.FormValue("username")
		ctg.Message = r.FormValue("message")
		ctg.Email = r.FormValue("email")

		a, b, c := time.Now().Date()
		d := b.String()
		e := time.Now().Hour()
		ctg.PlacedAt = strconv.Itoa(e) + "  " + strconv.Itoa(c) + "/" + d + "/" + strconv.Itoa(a)
		fmt.Println(ctg.UserName, ctg.Message, ctg.Email, ctg.PlacedAt)
		err := ach.CommentSrv.UpdateComment(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/comment", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/comment", http.StatusSeeOther)
	}

}

// AdminCommentsDelete handle requests on route /admin/categories/delete
func (ach *AdminCommentHandler) AdminCommentsDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = ach.CommentSrv.DeleteComment(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/comment", http.StatusSeeOther)
}
