package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

func recoverMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Errorf(
					"err: %v trace %v", err, debug.Stack(),
				)
			}
		}()
		next(w, r)
	}
}

func post(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeHttpError(w, http.StatusBadRequest, "only POST method is supported")
			return
		}
		next(w, r)
	}
}

func writeHttpError(resp http.ResponseWriter, status int, comment string) {
	body := struct {
		Error string `json:"error"`
	}{
		Error: comment,
	}
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(status)
	err := json.NewEncoder(resp).Encode(body)
	if err != nil {
		log.Errorf("json encode error: %v", err)
	}
}
