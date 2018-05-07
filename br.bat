set GOPATH=c:\dev\projects\settingssvc
cd c:\dev\projects\settingssvc\src\cmd\settings-server
go install 


rem RUN
c:\dev\projects\settingssvc\bin\settings-server.exe --port=4201