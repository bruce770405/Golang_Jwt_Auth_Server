package data

import (
	"net/http"
)
import "login"

/**
server資源
*/
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	response := Response{"Gained access to protected resource"}
	login.JsonResponse(response, w)

}
