package api

/*
go run ./cmd/webapi/
curl -v \
	-X DELETE \
	-H 'Authorization: Plutone' \
	localhost:3000/photo/{1}
*/

import (
	"net/http"
	"strconv"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := r.Header.Get("Authorization")
	dbUser, present, err := rt.db.GetUserID(userID)

	if err != nil || !validUsername(dbUser.Username) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !present {
		http.Error(w, "user not found", http.StatusNotFound)
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

	if dbPhoto.Owner != dbUser.ID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.DeletePhoto(pathID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var photo Photo
	photo.FromDatabase(dbPhoto)
	err = DeletePhotoFile(photo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
