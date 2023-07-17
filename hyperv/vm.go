package hyperv

import (
	"net/http"
	"wmi-rest/utilities"

	"github.com/gin-gonic/gin"
)

func VMS(c *gin.Context) {
	cmdline := `Get-WmiObject -namespace 'root\virtualization\v2' -class Msvm_ComputerSystem -Filter 'Caption="Virtual Machine"' | Select-Object -Property ElementName, InstallDate, Name, ProcessID | ConvertTo-Json`
	output, err := utilities.CommandLine(cmdline)

	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	if len(output) == 0 {
		c.Data(returnResponse("No VM found.", http.StatusInternalServerError, "failure", "error"))
		return
	}

	c.Data(returnResponse(output, http.StatusOK, "success", "VMs displayed in data field"))
}
