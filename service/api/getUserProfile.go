package api

/*
go run ./cmd/webapi/
curl -v \
	-X GET \
	-H 'Authorization: Zeus' \
	localhost:3000/users/{Ares}/
*/

import (
	"encoding/json"
	"net/http"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/LinCannn/WASA/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingID := r.Header.Get("Authorization")
	requestedID := ps.ByName("id")

	var followers []database.User
	var following []database.User
	var photos []database.Photo

	requester, present, err := rt.db.GetUserID(requestingID)

	if err != nil || !validUsername(requester.Username) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	requested, present, err := rt.db.GetUserID(requestedID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	followers, err = rt.db.GetFollowers(requested)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	following, err = rt.db.GetFollowing(requested)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bannedList, err := rt.db.GetBannedList(requested)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photos, err = rt.db.GetPhotosList(requested)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var FullUser database.FullUser
	FullUser.ID = requested.ID
	FullUser.Username = requested.Username
	FullUser.Followers = followers
	FullUser.Following = following
	FullUser.Banned = bannedList
	FullUser.Photo = photos

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(FullUser)
}
