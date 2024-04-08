package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "See https://github.com/gladkikhartem/csv2xls")
	})

	r.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		id := fmt.Sprint(rnd.Int63())
		d, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "err reading body: %v", err)
			w.WriteHeader(400)
			return
		}
		_ = os.MkdirAll(fmt.Sprintf("%v", id), 0777)
		err = os.WriteFile(fmt.Sprintf("%v/Default", id), d, 0777)
		if err != nil {
			fmt.Fprintf(w, "err writing file: %v", err)
			w.WriteHeader(400)
			return
		}
		out, err := exec.Command("ssconvert", fmt.Sprintf("%v/Default", id), fmt.Sprintf("%v.xls", id)).CombinedOutput()
		if err != nil {
			fmt.Fprintf(w, "err converting file: %v %v", err, string(out))
			w.WriteHeader(400)
			return
		}
		d, err = os.ReadFile(fmt.Sprintf("%v.xls", id))
		if err != nil {
			fmt.Fprintf(w, "err reading file: %v ", err)
			w.WriteHeader(400)
			return
		}
		_, _ = w.Write(d)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
