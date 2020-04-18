# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

build:
	@echo "########## Compilando nossa API ... "
	go build -ldflags="-s -w" -o gocep main.go
	upx gocep
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./gocep

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/gocep
	go build -ldflags="-s -w" -o gocep main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"
	@echo "fim"

deploy.aws:
	@echo "########## Compilando nossa API ... "
	sh deploy-aws.sh
	@echo "fim"
