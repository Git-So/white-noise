# 环境 linux_amd64

AppName = "white-noise"

init:
	mkdir -p cache
#	git clone https://github.com/therecipe/env_linux_amd64_512.git cache/env_linux_amd64_512
	git clone https://github.com/therecipe/env_linux_amd64_513.git cache/env_linux_amd64_513

rebuild:
	qtdeploy build desktop

build:
	make clear
	make rebuild

clear:
	-rm -rf vendor deploy linux config/log
	-rm -rf rcc* moc* ./*/moc*
	go mod vendor
	go mod download
	cp -r cache/env_linux_amd64_513 vendor/github.com/therecipe/

run:
	make rebuild
	make rerun

rerun:
	./deploy/linux/${AppName}
	-rm -rf rcc* ./*/moc*

test:
ifeq ($(findstring _test.go,$(testFile)),_test.go)
# go test
# 预测试环境
	-@ln -s $(rootPath)/config $(filePath)/config
#开始测试
	-go test -timeout 30s -v $(testFile)
# 清理测试环境
	-@rm -rf $(filePath)/config
else
# app test
	@echo "🤪待写吧..."
endif
