package datastorehandler

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", baseHandler)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {

}
