package api

/*
go run ./cmd/webapi/
curl -v \
	-X GET \
	-H 'Authorization: Urano' \
	"localhost:3000/search/?username=Ares"
*/

import (
	"encoding/json"
	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUserByUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID := r.Header.Get("Authorization")

	_, present, err := rt.db.GetUserID(userID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	username := r.URL.Query().Get("username")
	usersList, err := rt.db.SearchUser(username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(usersList)
}
