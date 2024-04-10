package api

/*
go run ./cmd/webapi/
curl -v \
	-X POST \
	-H 'Content-Type: text/plain' \
	-H 'Authorization: Urano' \
	-d "Bellissimo " \
	localhost:3000/photo/{6}/comment
*/

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID := r.Header.Get("Authorization")
	author, present, err := rt.db.GetUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var pathID int64
	pathID, err = strconv.ParseInt(ps.ByName("pid"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo, present, err := rt.db.GetPhoto(pathID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Owner, present, err := rt.db.GetUserID(photo.Owner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	IsBanned1, err1 := rt.db.BannedUserCheck(author, Owner)
	IsBanned2, err2 := rt.db.BannedUserCheck(Owner, author)

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if IsBanned1 || IsBanned2 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	text := string(body)
	comment := Comment{
		AuthorID:       author.ID,
		AuthorUsername: author.Username,
		PhotoID:        photo.ID,
		Text:           text,
		Date:           time.Now()}

	dbComment, err := rt.db.CommentPhoto(comment.ToDatabase())
	comment.FromDatabase(dbComment)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)

}
