package arfblocks

import (
	"encoding/json"
	"net/http"
)

func OperateHttpRequest[T IEndoint](response http.ResponseWriter, endpoint T) {

	// Validation Phase
	validationErr := endpoint.Validate()
	if validationErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + validationErr.Error() + `" }`))
		return
	}

	// Verification Phase
	verificationErr := endpoint.Verificate()
	if verificationErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + verificationErr.Error() + `" }`))
		return
	}

	// Handling Phase
	data, err := endpoint.Handle()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(data)
}
