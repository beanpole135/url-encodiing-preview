run:
	go run main.go

package:
	export GOOS="windows"; go build -o bin/convert_win.exe main.go 
	export GOOS="darwin"; go build -o bin/convert_mac main.go 
	export GOOS="linux"; go build -o bin/convert_linux main.go 