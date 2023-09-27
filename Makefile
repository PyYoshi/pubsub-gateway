setup-development:
	go install goa.design/goa/v3/cmd/goa@v3

generate:
	goa gen github.com/PyYoshi/pubsub-gateway/design -o ./
	# rm -rf ./cmd/gateway
	# rm -rf ./cmd/gateway-cli
	# rm -rf ./gcp.go
	# rm -rf ./healthz.go
	# rm -rf ./swagger.go
	goa example github.com/PyYoshi/pubsub-gateway/design -o ./

build-docker-images:
	docker build -t pyyoshi/pubsub-gateway-server:latest \
		-f ./dockerfiles/gateway/Dockerfile .

dev:
	cd cmd/
