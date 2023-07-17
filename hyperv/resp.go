package hyperv

import "encoding/json"

type response struct {
	Result  string      `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func returnResponse(respData interface{}, status int, result, message string) (code int, contentType string, data []byte) {

	value, ok := respData.(string)
	if ok {
		respData, _ = json.Marshal(value)
	}
	resp := response{
		Result:  result,
		Message: message,
		Data:    json.RawMessage(respData.([]byte)),
	}
	jsonResp, _ := json.MarshalIndent(resp, "", "    ")
	return status, "application/json", jsonResp
}
