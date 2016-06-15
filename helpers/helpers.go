package helpers

import (
	"net/http"

	"github.com/dixonwille/RussellCatering/models"
)

// Write is used to write content to the writer with the status
// Accepts a model and with turn it into json before writing.
func Write(writer http.ResponseWriter, status int, content interface{}) {
	obj, ok := content.(models.Publicer)
	if ok {
		content = obj.Public()
	}
	cont, err := models.Jsonify(content)
	if err != nil {
		msg := models.NewError(err.Error())
		Write(writer, http.StatusInternalServerError, msg)
		return
	}
	writer.WriteHeader(status)
	writer.Write(cont)
}
