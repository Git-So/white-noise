# 环境 linux_amd64

AppName = "white-noise"

init:
	mkdir -p cache
	git clone https://github.com/therecipe/env_linux_amd64_512.git cache/env_linux_amd64_512

rebuild:
	qtdeploy build desktop

build:
	make clear
	make rebuild

clear:
	-rm -rf vendor rcc* deploy linux
	go mod vendor
	go mod download
	cp -r cache/env_linux_amd64_512 vendor/github.com/therecipe/

run:
	make build
	make rerun

rerun:
	./deploy/linux/${AppName}

test:
	echo "🤪还没有写,正在学习QT"
