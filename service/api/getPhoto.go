package api

/*
go run ./cmd/webapi/
curl -v \
	-X GET \
	localhost:3000/photo/{1}
*/

import (
	"net/http"
	"strconv"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var pathID int64
	pathID, err := strconv.ParseInt(ps.ByName("pid"), 10, 64)

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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var photo Photo
	photo.FromDatabase(dbPhoto)
	filepath := photo.path()

	w.Header().Set("Content-Type", "image/*")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, filepath)
}
