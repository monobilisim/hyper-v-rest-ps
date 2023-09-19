@ECHO OFF
setlocal enabledelayedexpansion
set "ServiceNames[0]=wmi-rest"
set "ServiceNames[1]=hyper-v-rest-wmi"
set "ServiceNames[2]=hyper-v-rest-ps"

for /l %%i in (0, 1, 2) do (
	set "ServiceName=!ServiceNames[%%i]!"
	sc query !ServiceName! | find "STATE" | find "RUNNING" >nul
	if !errorLevel! equ 0 (
		sc stop !ServiceName!
	)
	sc delete !ServiceName!
	taskkill /t /f /im !ServiceName!.exe
)

endlocal
