package main

import (
	"encoding/json"
	"expvar"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/kelseyhightower/envconfig"

	"github.com/353solutions/nlp"
)

// Config is application configuration
type Config struct {
	Port int
}

func parseConfig() (*Config, error) {
	cfg := &Config{
		Port: 8080, // default
	}
	err := envconfig.Process("NLPD", cfg)
	if err != nil {
		return nil, err
	}
	// TODO: Validate configuration
	return cfg, nil
}

// Add a "port" metric on which port the server is listening on

var (
	numTokenize = expvar.NewInt("num_tokenize")
)

func main() {
	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "usage: %s\n", name)
		fmt.Fprintf(os.Stderr, "NLP HTTP server\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	cfg, err := parseConfig()
	if err != nil {
		// Only main is allowed to exit the porgram
		log.Fatal(err)
	}

	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("server listening on %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	numTokenize.Add(1)
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("can't read request - %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokens := nlp.Tokenize(string(data))
	out, err := json.Marshal(tokens)
	if err != nil {
		log.Printf("can't marshal output - %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Sanity checks
	fmt.Fprintf(w, "OK\n")
}
