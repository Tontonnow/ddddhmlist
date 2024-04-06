#交叉编译
#Windows PowerShell
#年月日
$date=$(Get-Date -Format "yyyyMMdd")
$name="ddddhmlistclient"
#win 64
$env:CGO_ENABLED=0
$env:GOOS="windows"
$env:GOARCH="amd64"
go  build -o ./bin/$name-$date-win64.exe  -ldflags "-s -w " ./client.go
upx -9 ./bin/$name-$date-win64.exe

#win 32
$env:GOOS="windows"
$env:GOARCH="386"
go  build -o  ./bin/$name-$date-win32.exe -ldflags "-s -w " ./client.go
upx -9 ./bin/$name-$date-win32.exe

#linux 64
$env:GOOS="linux"
$env:GOARCH="amd64"
go  build -o ./bin/$name-$date-linux64 -ldflags "-s -w " ./client.go
upx -9 ./bin/$name-$date-linux64

#linux x86
$env:GOOS="linux"
$env:GOARCH="386"
go  build -o ./bin/$name-$date-linux32 -ldflags "-s -w " ./client.go
upx -9 ./bin/$name-$date-linux32
#linux arm
$env:GOOS="linux"
$env:GOARCH="arm"
go  build -o ./bin/$name-$date-linux-arm -ldflags "-s -w " ./client.go
upx -9 ./bin/$name-$date-linux-arm

#linux arm64
$env:GOOS="linux"
$env:GOARCH="arm64"
go  build -o ./bin/$name-$date-linux-arm64 -ldflags "-s -w " ./client.go
upx -9 ./bin/$name-$date-linux-arm64


$env:GOOS="darwin"
$env:GOARCH="amd64"
go  build -o ./bin/$name-$date-mac64 -ldflags "-s -w " ./client.go


#mac arm
$env:GOOS="darwin"
$env:GOARCH="arm"
go  build -o ./bin/$name-$date-mac-arm -ldflags "-s -w " ./client.go

#mac arm64
$env:GOOS="darwin"
$env:GOARCH="arm64"
go  build -o ./bin/$name-$date-mac-arm64 -ldflags "-s -w " ./client.go





