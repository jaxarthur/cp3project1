build :	main.go
		env GOOS=windows GOARCH=amd64 go build -o project1win.exe
		env GOOS=darwin GOARCH=amd64 go build -o project1mac
		env GOOS=linux GOARCH=amd64 go build -o project1linux