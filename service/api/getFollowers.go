package api

/*
go run ./cmd/webapi/
curl -v \
	-X GET \
	-H 'Authorization: Urano' \
	localhost:3000/users/{Urano}/followers/
*/

import (
	"encoding/json"
	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingID := r.Header.Get("Authorization")
	requestedID := ps.ByName("id")

	requester, present1, err1 := rt.db.GetUserID(requestingID)
	requested, present2, err2 := rt.db.GetUserID(requestedID)

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !present1 || !present2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	isBanned1, err1 := rt.db.BannedUserCheck(requester, requested)
	isBanned2, err2 := rt.db.BannedUserCheck(requested, requester)
	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isBanned1 || isBanned2 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	followers, err := rt.db.GetFollowers(requested)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(followers)
}
