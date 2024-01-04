COLOUR_GREEN=\033[0;32m
COLOUR_RED=\033[0;31m
COLOUR_BLUE=\033[0;34m
END_COLOUR=\033[0m

.DEFAULT_GOAL := all

cleanup:
	@echo "$(COLOUR_BLUE)\nCleaning Amazonia executables$(END_COLOUR)"
	rm -rf bin/

startmsg:
	@echo "$(COLOUR_BLUE)\nBuilding Amazonia$(END_COLOUR)"

build:
	go build -o bin/main main.go

run:
	@echo "$(COLOUR_GREEN)\nRunning Amazonia. Logging with $(COLOUR_BLUE)logs $(COLOUR_GREEN)file$(END_COLOUR)\n"
	rm -rf logs*
	go run main.go | tee logs

# For more read this 
# https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures
compile: cleanup
	@echo "$(COLOUR_BLUE)\nCompiling for every OS and Platform$(END_COLOUR)"
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/main-darwin-arm64 main.go
	@echo "$(COLOUR_GREEN)\nCompile successful$(END_COLOUR)"

endmsg:
	@echo "$(COLOUR_GREEN)\nBuild successful$(END_COLOUR)"

all: startmsg cleanup build endmsg