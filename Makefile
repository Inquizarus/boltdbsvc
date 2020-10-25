# Go parameters
export GOARM=${GOARM:-}

GOOS=linux

ifeq ($(UNAME_S),Darwin)
  GOOS = darwin
endif


ifeq ($(GOARCH),)
  GOARCH = amd64
  ifneq ($(UNAME_M), x86_64)
    GOARCH = 386
  endif
endif

my_d = $(shell pwd)
OUT_D = $(shell echo $${OUT_D:-$(my_d)/builds})

all: test build _build_linux_arm6 _build_linux_amd64
test: 
	go test -v ./...

clean: 
	rm -rf builds

.PHONY: build
build: _build
	@echo "==> build local version"
	@echo ""
	@mv $(OUT_D)/golbag_$(GOOS)_$(GOARCH) $(OUT_D)/golbag
	@echo "installed as $(OUT_D)/golbag"

.PHONY: _build
_build:
	@echo "=> building golbag via go build"
	@echo ""
	@OUT_D=${OUT_D} GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} scripts/build.sh
	@echo "built $(OUT_D)/golbag_$(GOOS)_$(GOARCH)"

.PHONY: _build_linux_amd64
_build_linux_amd64:
_build_linux_amd64: GOOS=linux
_build_linux_amd64: GOARCH=amd64
_build_linux_amd64: GOARM=
_build_linux_amd64: _build

.PHONY: _build_linux_arm6
_build_linux_arm6:
_build_linux_arm6: GOOS=linux
_build_linux_arm6: GOARCH=arm
_build_linux_arm6: GOARM=6
_build_linux_arm6: _build