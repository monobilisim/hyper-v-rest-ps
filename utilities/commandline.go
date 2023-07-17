package utilities

import (
	"github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

func initPwsh() powershell.Shell {
	pwsh, err := powershell.New(&backend.Local{})
	if err != nil {
		panic(err)
	}
	return pwsh
}

var pwsh powershell.Shell = initPwsh()

func CommandLine(ps string) ([]byte, error) {
	output, _, err := pwsh.Execute(ps)
	return []byte(output), err
}
