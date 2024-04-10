package api

/*
go run ./cmd/webapi/
curl -v \
	-X DELETE \
	-H 'Authorization: Urano' \
	localhost:3000/like/{6}
*/

import (
	"net/http"
	"strconv"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID := r.Header.Get("Authorization")
	userLike, present, err := rt.db.GetUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var pathID int64
	pathID, err = strconv.ParseInt(ps.ByName("pid"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbPhoto, present, err := rt.db.GetPhoto(pathID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Owner, present, err := rt.db.GetUserID(dbPhoto.Owner)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !present {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	IsBanned1, err1 := rt.db.BannedUserCheck(userLike, Owner)
	IsBanned2, err2 := rt.db.BannedUserCheck(Owner, userLike)

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if IsBanned1 || IsBanned2 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.UnlikePhoto(pathID, userLike.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
