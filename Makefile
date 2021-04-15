binary = hey

release:
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(binary)_windows_amd64
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(binary)_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(binary)_darwin_amd64

plugin-build:
	go build -buildmode=plugin -o=./plugin/bin/create_disbursement.so ./plugin/xendit/create_disbursement.go
	go build -buildmode=plugin -o=./plugin/bin/create_disbursement_payment_engine.so ./plugin/payment-engine/create_disbursement.go
	go build -buildmode=plugin -o=./plugin/bin/openpo_login.so ./plugin/openpo/login/login.go

all: plugin-build release
