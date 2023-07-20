package hyperv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wmi-rest/utilities"

	"github.com/gin-gonic/gin"
)

type VHDPath []VHDPathElement

type VHDPathElement struct {
	Path string `json:"Path"`
	VMID string `json:"VMId"`
}

type VHDInfo struct {
	Size int64  `json:"Size"`
	Id   string `json:"Id"`
}

var VHDPathList VHDPath

func VHD(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if input == "all" {
		//  Get-VHD -Path C:\Hyper-V\VM243.vhdx |Select-Object -Property Size | ConvertTo-Json

		var sizeList []VHDInfo
		var output []byte
		var err error
		sizeList = make([]VHDInfo, len(VHDPathList))
		for i, v := range VHDPathList {
			output, err = utilities.CommandLine(`Get-VHD -Path '` + v.Path + `' | Select-Object -Property Size | ConvertTo-Json`)
			if err != nil {
				c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
				return
			}
			json.Unmarshal(output, &sizeList[i])

			sizeList[i].Id = v.VMID

		}

		jsonOutput, err := json.Marshal(sizeList)
		if err != nil {
			c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
			return
		}

		c.Data(returnResponse(jsonOutput, http.StatusOK, "success", "VHD info is displayed in data field."))
		return
	}

	output, err := utilities.CommandLine(`Get-VHD -Id ` + input + ` | ConvertTo-Json`)
	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	if len(output) < 1 {
		c.Data(returnResponse("No Disk found.", http.StatusOK, "failure", "error"))
		return
	}

	c.Data(returnResponse(output, http.StatusOK, "success", "VHD info is displayed in data field"))
}

func UnmarshalVHDPath(data []byte) (VHDPath, error) {
	var r VHDPath
	err := json.Unmarshal(data, &r)
	return r, err
}

func Initialize() {
	output, err := utilities.CommandLine(`Get-VM | Get-VMHardDiskDrive | Select-Object -Property Path, VMId | ConvertTo-Json`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	VHDPathList, err = UnmarshalVHDPath(output)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Finished initializing VHDPathList.")
}
