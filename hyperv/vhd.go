package hyperv

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
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
var pathListLock sync.Mutex

func VHD(c *gin.Context) {

	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if input == "all" {
		getAllVHDInfo(c)
		return
	}

	output := getVHDInfo(input, c)

	if len(output) < 1 {
		c.Data(returnResponse("No Disk found.", http.StatusOK, "failure", "error"))
		return
	}

	c.Data(returnResponse(output, http.StatusOK, "success", "VHD info is displayed in data field"))
}

func getAllVHDInfo(c *gin.Context) {
	var sizeList []VHDInfo
	var output []byte
	var err error
	pathListLock.Lock()
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
	pathListLock.Unlock()

	jsonOutput, err := json.Marshal(sizeList)
	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	c.Data(returnResponse(jsonOutput, http.StatusOK, "success", "VHD info is displayed in data field."))
}

func getVHDInfo(input string, c *gin.Context) []byte {
	output, err := utilities.CommandLine(`Get-VHD -Id ` + input + ` | ConvertTo-Json`)
	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return nil
	}
	return output
}

func UnmarshalVHDPath(data []byte) (VHDPath, error) {
	var r VHDPath
	err := json.Unmarshal(data, &r)
	return r, err
}

func Refresh() {
	path, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path, %s", err)
	}
	executablePath := filepath.Dir(path)

	pathListLock.Lock()
	err = Load(executablePath+"\\vhdpath.gob", &VHDPathList)
	if err != nil {
		fmt.Println("Unable to load VHDPathList from file. Loading from PowerShell.")
	} else {
		fmt.Println("Finished loading VHDPathList from disk.")
	}
	pathListLock.Unlock()

	fmt.Println("Reinitializing VHDPathList.")
	output, err := utilities.CommandLine(`Get-VM | Get-VMHardDiskDrive | Select-Object -Property Path, VMId | ConvertTo-Json`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Waiting for VHDPathList lock.")
	pathListLock.Lock()
	VHDPathList, err = UnmarshalVHDPath(output)
	pathListLock.Unlock()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Finished reinitializing VHDPathList.")

	err = Save(executablePath+"\\vhdpath.gob", VHDPathList)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func Init() {
	path, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path, %s", err)
	}
	executablePath := filepath.Dir(path)

	pathListLock.Lock()
	err = Load(executablePath+"\\vhdpath.gob", &VHDPathList)
	if err != nil {
		fmt.Println("Unable to load VHDPathList from file. Loading from PowerShell.")
	} else {
		fmt.Println("Finished loading VHDPathList from disk.")
	}
	pathListLock.Unlock()
}

func Save(path string, data VHDPath) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func Load(path string, data *VHDPath) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}
