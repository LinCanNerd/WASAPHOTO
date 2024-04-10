package api

/*
go run ./cmd/webapi/

curl -v \
	-X PUT \
	-H 'Authorization: Urano'\
	-H 'Content-Type: application/json' \
	-d '{"username": "Uranissimo"}' \
	localhost:3000/settings/
*/

import (
	"encoding/json"

	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID := r.Header.Get("Authorization")
	dbUser, present, err := rt.db.GetUserID(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !present {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	var newUser User
	newUser.FromDatabase(dbUser)
	err = json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !validUsername(newUser.Username) {
		http.Error(w, "invalid username", http.StatusBadRequest)
		return
	}

	_, present, err = rt.db.GetUsername(newUser.Username)

	if err != nil {
		http.Error(w, "username already taken", http.StatusBadRequest)
		return
	}
	if present {
		http.Error(w, "username already taken", http.StatusBadRequest)
		return
	}

	err = rt.db.UpdateUsername(newUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(newUser)
}
