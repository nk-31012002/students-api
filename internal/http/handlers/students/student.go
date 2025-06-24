package students

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/nk-31012002/student-api/internal/types"
	"github.com/nk-31012002/student-api/internal/utils/response"
	"io"
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a new student")

		var students types.Student

		err := json.NewDecoder(r.Body).Decode(&students)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//request validation
		if err := validator.New().Struct(students); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
