package utilities

import (
	"github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

var pwsh powershell.Shell

func InitPwsh() {
	var err error
	pwsh, err = powershell.New(&backend.Local{})
	if err != nil {
		panic(err)
	}
}

func CommandLine(ps string) ([]byte, error) {
	output, _, err := pwsh.Execute(ps)
	return []byte(output), err
}
