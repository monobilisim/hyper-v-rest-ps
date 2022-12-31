//go:build windows

package rest

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"wmi-rest/wmi"
)

func vm(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		httpError(w, errors.New("name is missing in parameters"), http.StatusBadRequest, resp)
		return
	}

	vms, err := wmi.VM(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(vms) == 0 {
		httpError(w, errors.New("no VM found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "VM info is displayed in data field."
	resp.Data = vms

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func vms(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	v, err := wmi.VMs()
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(v) == 0 {
		httpError(w, errors.New("no VM found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "VMs are listed in data field."
	resp.Data = v

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func memory(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		httpError(w, errors.New("name is missing in parameters"), http.StatusBadRequest, resp)
		return
	}

	m, err := wmi.Memory(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(m) == 0 {
		httpError(w, errors.New("no memory info found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "Memory info is displayed in data field."
	resp.Data = m

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func processor(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		httpError(w, errors.New("name is missing in parameters"), http.StatusBadRequest, resp)
		return
	}

	p, err := wmi.Processor(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(p) == 0 {
		httpError(w, errors.New("no processor info found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "Processor info is displayed in data field."
	resp.Data = p

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func storage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		httpError(w, errors.New("name is missing in parameters"), http.StatusBadRequest, resp)
		return
	}

	s, err := wmi.Storage(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(s) == 0 {
		httpError(w, errors.New("no storage info found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "Storage info is displayed in data field."
	resp.Data = s

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}
