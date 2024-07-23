.PHONY: build deploy clean

build:
	GOOS=linux GOARCH=amd64 go build -o ./bootstrap ./cmd/main.go

deploy:
	cd ../deploy-auth && cdk deploy

deploy-swap:
	cd ../deploy-auth && cdk deploy --hotswap

clean:
	rm -rf ./bootstrap
	
