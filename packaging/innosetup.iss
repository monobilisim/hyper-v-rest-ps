[Setup]
AppName=HyperV REST PS
AppVersion=2.1
AppPublisher=Mono Bilisim
AppPublisherURL=https://mono.net.tr
AppSupportURL=https://github.com/monobilisim/hyper-v-rest-ps
AppUpdatesURL=https://github.com/monobilisim/hyper-v-rest-ps
DefaultDirName={code:GetProgramFiles}\hyper-v-rest-ps
UsePreviousAppDir=false
UninstallDisplayIcon={app}\hyper-v-rest-ps.exe
OutputBaseFilename=hyper-v-rest-ps-setup
Compression=lzma
SolidCompression=yes

[Code]
function GetProgramFiles(Param: string): string;
begin
  if IsWin64 then Result := ExpandConstant('{commonpf64}')
    else Result := ExpandConstant('{commonpf32}')
end;

procedure TaskKill(FileName: String);
var
  ResultCode: Integer;
begin
    Exec('taskkill.exe', '/f /im ' + '"' + FileName + '"', '', SW_HIDE,
     ewWaitUntilTerminated, ResultCode);
end;


[Files]
Source: "hyper-v-rest-ps.exe"; DestDir: "{app}"; BeforeInstall: TaskKill('hyper-v-rest-ps.exe')
Source: "cleanup.bat"; DestDir: "{app}"


[Icons]
;Name: "{group}\HyperV REST PS"; Filename: "{app}\hyper-v-rest-ps.exe"
;Name: "{group}\Uninstall"; Filename: "{uninstallexe}"


[Messages]
SetupAppTitle=HyperV REST PS


[Run]
Filename: "{app}\cleanup.bat"
Filename: "{app}\hyper-v-rest-ps.exe"; Description: "Install Service"; Parameters: --service=install
Filename: "{app}\hyper-v-rest-ps.exe"; Description: "Start Service"  ; Parameters: --service=start


[UninstallRun]
Filename: "{app}\cleanup.bat"
Filename: "{app}\hyper-v-rest-ps.exe"; Parameters: --service=stop
Filename: "{app}\hyper-v-rest-ps.exe"; Parameters: --service=uninstall
Filename: "{cmd}"; Parameters: "/C ""taskkill /im hyper-v-rest-ps.exe /f /t"
