package main

import (
	"cmp"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"os"
)

var (
	log           = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	port          = cmp.Or(os.Getenv("PORT"), "8000")
	xServerHeader = cmp.Or(os.Getenv("X_SERVER_HEADER"), fmt.Sprintf("server#%d", rand.IntN(100)))
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", HTTPRequestAttr("http request", r))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-SERVER-HEADER", xServerHeader)
		body := `{"status":"ok"}`
		fmt.Fprint(w, body)
		log.Info("response", HTTPResponseAttr("http response", w, body))
	})

	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", HTTPRequestAttr("http request", r))
		name := cmp.Or(r.URL.Query().Get("name"), "World")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-SERVER-HEADER", xServerHeader)
		body := fmt.Sprintf(`{"message":"Hello %s"}`, name)
		fmt.Fprint(w, body)
		log.Info("response", HTTPResponseAttr("http response", w, body))
	})

	log.Info("listening on :" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Error("http listen and serve failed", slog.String("err", err.Error()))
	}
}
