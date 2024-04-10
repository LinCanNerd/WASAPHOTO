package api

/*
go run ./cmd/webapi/
curl -v \
	-X POST \
	-H 'Authorization: Ares' \
	-H 'Content-Type: image/*' \
	--data-binary "@./photo-samples/atene1.jpg" \
	localhost:3000/photo
*/

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/LinCannn/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	photoData, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	format := http.DetectContentType(photoData)
	switch format {
	case "image/jpeg":
		format = "jpeg"
	case "image/png":
		format = "png"
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	binaryPhoto := photoData
	photo := Photo{
		Owner:    userID,
		Username: dbUser.Username,
		Format:   format,
		Date:     time.Now(),
	}

	dbPhoto, err := rt.db.CreatePhoto(photo.ToDatabase())
	photo.FromDatabase(dbPhoto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = CreatePhotoFile(photo, binaryPhoto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)

}
