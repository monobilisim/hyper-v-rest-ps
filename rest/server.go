//go:build windows

package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{port: port}
}

type response struct {
	Result  string      `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (s *Server) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/vms", vms)
	r.HandleFunc("/vms/{name}", vm)
	r.HandleFunc("/vms/{name}/memory", memory)
	r.HandleFunc("/vms/{name}/processor", processor)
	r.HandleFunc("/vms/{name}/storage", storage)

	err := http.ListenAndServe(":"+strconv.Itoa(s.port), r)
	if err != nil {
		log.Fatalln(err)
	}
}

func httpError(w http.ResponseWriter, err error, code int, resp response) {
	w.WriteHeader(code)
	resp.Result = "error"
	resp.Message = err.Error()
	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}
