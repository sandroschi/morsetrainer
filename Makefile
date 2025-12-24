all: buildlinux buildwindows zip

buildwindows:
	GOOS=windows GOARCH=amd64 go build -o bin/win64/morsetrainer.exe

buildlinux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux64/morsetrainer

zip:
	zip -j bin/morsetrainer_win64.zip config.toml bin/win64/morsetrainer.exe
	zip -j bin/morsetrainer_linux64.gz config.toml bin/linux64/morsetrainer 