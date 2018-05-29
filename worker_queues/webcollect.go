package main

import (
	"fmt"
	"net/http"
)

func WebCollect(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		http.Error(w, "You must send some content.", http.StatusBadRequest)
		return
	}

	work := WorkRequest{Content: content}
	WorkQueue <- work
	fmt.Println("Work request queued from web handler", len(WorkQueue))

	w.WriteHeader(http.StatusCreated)
	return
}
