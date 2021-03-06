package main

import (
	"net/http"
	"os/exec"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func enableCustomHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "Access-Token")
}

func date(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	enableCustomHeaders(&w)
	accessKey := r.Header.Get("Access-Token")
	if accessKey != "yobro123" {
		//w.WriteHeader(http.StatusForbidden)

	}
	cmd := exec.Command("date")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/date", date)
	http.ListenAndServe(":8081", mux)
}
