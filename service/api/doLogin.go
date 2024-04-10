package api

/*
go run ./cmd/webapi/
curl -v \
	-X POST \
	-H 'Content-Type: application/json' \
	-d '{"username":"Batman"}' \
	localhost:3000/session
*/

import (
	"encoding/json"
	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil || !validUsername(user.Username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbUser, present, err := rt.db.GetUserID(user.Username)

	if present {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusCreated)
		dbUser, err = rt.db.CreateUser(user.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	user.FromDatabase(dbUser)
	_ = json.NewEncoder(w).Encode(user)
}
