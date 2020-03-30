package upload

import (
	"fmt"
	"github.com/ATechnoHazard/audienceai-backend/api/views"
	"github.com/ATechnoHazard/audienceai-backend/internal/utils"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
)

func upVid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, handler, err := r.FormFile("video")
		fileName := r.FormValue("file_name")
		if err != nil {
			views.Wrap(err, w)
			return
		}
		defer file.Close()

		f, err := os.OpenFile(fmt.Sprintf("./videos/%s", handler.Filename), os.O_WRONLY|os.O_CREATE, 0744)
		if err != nil {
			views.Wrap(err, w)
			return
		}
		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			views.Wrap(err, w)
			return
		}

		msg := utils.Message(http.StatusAccepted, fmt.Sprintf("File %s uploaded successfully", fileName))
		utils.Respond(w, msg)
		return
	}
}

func MakeUpload(r *httprouter.Router) {
	r.HandlerFunc("POST", "/api/v1/upload", upVid())
}
