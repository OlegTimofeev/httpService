build:
	go mod download
	go build

build-docker:
	docker build -t taskService -f Dockerfile .

run-docker:
	docker run  -p 8080:8000 httpService


SWAGGER_IMAGE=quay.io/goswagger/swagger:v0.23.0

gen-taskService:
	docker run --rm -v `pwd`:/go/ -w /go/ -t $(SWAGGER_IMAGE) \
	generate server \
	--target=service \
	-f taskService.swagger.yml

	docker run --rm -v `pwd`:/go/ -w /go/ -t $(SWAGGER_IMAGE) \
	generate client \
	--target=service \
	-f taskService.swagger.yml

run queue:
	docker run -d -p 4222:4222 -p 8222:8222 -p 6222:6222 --name queue nats-streaming:0.14.0
