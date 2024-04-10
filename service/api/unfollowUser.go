package api

/*
go run ./cmd/webapi/
curl -v \
	-X DELETE \
	-H 'Authorization: Paperino' \
	localhost:3000/following/{topolino}
*/

import (
	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userToUnfollowId := ps.ByName("id")
	requesterId := r.Header.Get("Authorization")

	if requesterId == userToUnfollowId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requester, present1, err1 := rt.db.GetUserID(requesterId)
	userToUnfollow, present2, err2 := rt.db.GetUserID(userToUnfollowId)

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !present1 || !present2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := rt.db.UnfollowUser(requester.ID, userToUnfollow.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
