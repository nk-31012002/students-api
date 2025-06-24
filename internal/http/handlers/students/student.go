package students

import (
	"encoding/json"
	"errors"
	"github.com/nk-31012002/student-api/internal/types"
	"github.com/nk-31012002/student-api/internal/utils/response"
	"io"
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var students types.Student

		err := json.NewDecoder(r.Body).Decode(&students)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}

		slog.Info("Creating a new student")

		w.Write([]byte("Welome to the student api"))

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
