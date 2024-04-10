package api

/*
go run ./cmd/webapi/
curl -v \
	-X GET \
	-H 'Authorization: Zeus' \
	localhost:3000/photo/{1}/like
*/

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLikesList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	pathId := r.Header.Get("Authorization")

	photoID, err := strconv.ParseInt(ps.ByName("pid"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo, present, err := rt.db.GetPhoto(photoID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	requesting, present1, err1 := rt.db.GetUserID(pathId)
	requested, present2, err2 := rt.db.GetUserID(photo.Owner)
	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !present1 || !present2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	isBanned1, err1 := rt.db.BannedUserCheck(requesting, requested)
	isBanned2, err2 := rt.db.BannedUserCheck(requested, requesting)
	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isBanned1 || isBanned2 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	likes, err := rt.db.GetLikesList(photoID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(likes)
}
