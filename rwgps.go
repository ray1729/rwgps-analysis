package rwgps_analysis

import (
	"net/http"

	"github.com/ray1729/gpx-utils/pkg/rwgps"
)

var handler *rwgps.RWGPSHandler

func init() {
	h, err := rwgps.NewHandler()
	if err != nil {
		panic(err)
	}
	handler = h
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	handler.ServeHTTP(w, r)
}
