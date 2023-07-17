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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Data(http.StatusOK, "application/json", output)
}
