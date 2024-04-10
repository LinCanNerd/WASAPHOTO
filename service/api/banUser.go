package api

/*
go run ./cmd/webapi/
curl -v \
	-X PUT \
	-H 'Authorization: Urano' \
	localhost:3000/ban/{Ares}/
*/

import (
	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	isBanned, err := rt.db.BannedUserCheck(banner, bannedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.BanUser(banner, bannedUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err1 = rt.db.UnfollowUser(banner.ID, bannedUser.ID)
	err2 = rt.db.UnfollowUser(bannedUser.ID, banner.ID)

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.NoMoreFriends(banner, bannedUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
