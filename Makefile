setup-development:
	go get -u goa.design/goa/v3/cmd/goa

generate:
	goa gen github.com/PyYoshi/cloud_pubsub_gateway/design -o ./
	# rm -rf ./cmd/gateway
	# rm -rf ./cmd/gateway-cli
	# rm -rf ./gcp.go
	# rm -rf ./healthz.go
	# rm -rf ./swagger.go
	goa example github.com/PyYoshi/cloud_pubsub_gateway/design -o ./

build-docker-images:
	docker build -t pyyoshi/pubsub-gateway-server:latest \
		-f ./dockerfiles/gateway/Dockerfile .

dev:
	cd cmd/
