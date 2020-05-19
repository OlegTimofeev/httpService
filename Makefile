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
