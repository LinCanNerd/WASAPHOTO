package api

/*
go run ./cmd/webapi/
curl -v \
	-X DELETE \
	-H 'Authorization: Urano' \
	localhost:3000/photo/{6}/comment/{1}
*/

import (
	"net/http"
	"strconv"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID := r.Header.Get("Authorization")
	requestUser, present, err := rt.db.GetUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var photoID int64
	photoID, err = strconv.ParseInt(ps.ByName("pid"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbPhoto, present, err := rt.db.GetPhoto(photoID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var commentID int64
	commentID, err = strconv.ParseInt(ps.ByName("comment_id"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbComment, present, err := rt.db.GetCommentByID(commentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if dbComment.AuthorID != requestUser.ID && dbPhoto.Owner != requestUser.ID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.UncommentPhoto(commentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
