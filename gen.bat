cd C:\dev\projects\settingssvc\src
set GOPATH=C:\dev\projects\settingssvc
..\bin\goswagger generate server -f ..\..\settings\swagger\swagger.yaml -A settings

REM swagger generate server [-f ./swagger.json] -A [application-name [--principal [principal-name]]
