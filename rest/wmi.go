//go:build windows

package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"wmi-rest/wmi"

	"github.com/gorilla/mux"
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

	data, err := wmi.VM(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(data) == 0 {
		httpError(w, errors.New("no VM found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "VM info is displayed in data field."
	resp.Data = json.RawMessage(data)

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func vms(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	data, err := wmi.VMs()
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(data) == 0 {
		httpError(w, errors.New("no VM found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "VMs are listed in data field."
	resp.Data = json.RawMessage(data)

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func vhd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		httpError(w, errors.New("name is missing in parameters"), http.StatusBadRequest, resp)
		return
	}

	data, err := wmi.Vhd(name)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError, resp)
		return
	}

	if len(data) == 0 {
		httpError(w, errors.New("no image info found"), http.StatusNotFound, resp)
		return
	}

	resp.Result = "success"
	resp.Message = "Image info is displayed in data field."
	resp.Data = json.RawMessage(data)

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}

func version(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp response

	resp.Result = "success"
	resp.Message = "Version is displayed in data field."
	resp.Data = "0.3.1"

	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	_, _ = w.Write(jsonResp)
}
