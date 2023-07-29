BINARY=minimalchi

Version=`git tag -l | tail -1`
GitStatus=`git status -s`
BuildTime=`date +%FT%T%z`
BuildGoVersion=`go version`


LDFLAGS=-ldflags "-w -s \
-X 'main.BinVersion=${Version}' \
-X 'main.GitStatus=${GitStatus}' \
-X 'main.BuildTime=${BuildTime}' \
-X 'main.BuildGoVersion=${BuildGoVersion}' \
"

d:
	env GIT_TERMINAL_PROMPT=1 go get

t:
	env GIT_TERMINAL_PROMPT=1 go mod tidy

test:
	go test ./... -v

deps: d t

run:
	go run main.go serve

compile:
	CGO_ENABLED=0 go build -o releases/linux/${BINARY}${Version} ${LDFLAGS}

clean:
	if [ -f releases/linux/${BINARY}${Version} ] ; then rm releases/linux/${BINARY}${Version}; fi
	if [ -f releases/macos/${BINARY}${Version} ] ; then rm releases/macos/${BINARY}${Version}; fi
	if [ -f releases/windows/${BINARY}${Version}.exe ] ; then rm releases/windows/${BINARY}${Version}.exe; fi

buildredmond:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o releases/windows/${BINARY}${Version}.exe

buildapple:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o releases/macos/${BINARY}${Version}

build: clean deps tidy compile

all: build buildapple buildredmond
