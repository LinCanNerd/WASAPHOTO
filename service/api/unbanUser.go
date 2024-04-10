package api

/*
go run ./cmd/webapi/
curl -v \
	-X DELETE \
	-H 'Authorization: Plutone' \
	localhost:3000/ban/{Marte}
*/

import (
	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bannedID := ps.ByName("id")
	bannerID := r.Header.Get("Authorization")

	if bannerID == bannedID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	banner, present1, err1 := rt.db.GetUserID(bannerID)
	bannedUser, present2, err2 := rt.db.GetUserID(bannedID)

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !present1 || !present2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := rt.db.UnbanUser(banner, bannedUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
