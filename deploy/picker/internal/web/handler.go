package web

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gopherland/ntnx_labs/deploy/picker/internal/dictionary"
)

var words dictionary.WordList

// Response greeting message format.
type Response struct {
	AssetDir, Dictionary string
	WordCount            int
}

// LoadHandler loads a dictionary in memory.
func LoadHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		http.Error(w, "you must provide a dictionary name", http.StatusExpectationFailed)
		return
	}

	var err error
	words, err = dictionary.LoadDefault(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	resp := Response{
		AssetDir:   "assets",
		Dictionary: name,
		WordCount:  len(words),
	}
	buff, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(buff))
}

// PickHandler pick out a random word.
func PickHandler(w http.ResponseWriter, r *http.Request) {
	if len(words) == 0 {
		http.Error(w, "no dictionary loaded", http.StatusInternalServerError)
		return
	}
	word := words[rand.Intn(len(words))]

	resp := struct {
		Word string
	}{
		Word: word,
	}
	buff, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(buff))
}
