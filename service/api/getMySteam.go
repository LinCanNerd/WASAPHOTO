package api

/*
go run ./cmd/webapi/
curl -v \
	-X GET \
	-H 'Authorization: Urano' \
	localhost:3000/stream
*/
import (
	"encoding/json"
	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID := r.Header.Get("Authorization")

	user, present, err := rt.db.GetUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	stream, err := rt.db.GetMyStream(user.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i, photo := range stream {
		likes, err := rt.db.GetLikesList(photo.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stream[i].Likes = likes

		comments, err := rt.db.GetCommentsList(photo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stream[i].Comments = comments
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(stream)

}
