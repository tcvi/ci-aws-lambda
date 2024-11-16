.PHONY: build clean deploy
init:
	serverless plugin install -n serverless-apigateway-service-proxy
build:
	#dep ensure -v
	env GOOS=linux GOARCH=amd64 go build -o hello main.go


deploy_dev: build
	script/deploy_aws.sh dev

deploy_prod: build
	script/deploy_aws.sh prod
