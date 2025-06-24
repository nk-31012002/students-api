package students

import (
	"encoding/json"
	"errors"
	"github.com/nk-31012002/student-api/internal/types"
	"io"
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var students types.Student

		err := json.NewDecoder(r.Body).Decode(&students)
		if errors.Is(err, io.EOF) {

		}

		slog.Info("Creating a new student")

		w.Write([]byte("Welome to the student api"))
	}
}
