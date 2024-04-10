package api

/*
go run ./cmd/webapi/
curl -v \
	-X PUT \
	-H 'Authorization: Atena' \
	localhost:3000/following/{Ares}/
*/
import (
	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userToFollowId := ps.ByName("id")
	requesterId := r.Header.Get("Authorization")

	if requesterId == userToFollowId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requester, present1, err1 := rt.db.GetUserID(requesterId)
	userToFollow, present2, err2 := rt.db.GetUserID(userToFollowId)

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !present1 || !present2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	banned, err := rt.db.BannedUserCheck(requester, userToFollow)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.FollowUser(requester.ID, userToFollow.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
