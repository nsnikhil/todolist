test: build
	go clean -testcache
	go test ./...

setup: deps copy-config migrate build test

copy-config:
	cp application.yaml.sample application.yaml

test-cover-html: build
	go clean -testcache
	mkdir -p out/
	go test ./... -coverprofile=out/coverage.out
	go tool cover -html=out/coverage.out

build:
	mkdir -p out/
	go build -o out/main

deps:
	go mod vendor

tidy:
	go mod tidy

clean:
	rm -rf out/

fmt:
	gofmt -l -s -w .

vet:
	go vet ./...

serve: build
	./out/main serve

docker-serve:
	docker-compose -f deployment/docker/docker-compose.yaml up

k8-serve:
	chmod +x deployment/k8/start.sh
	./deployment/k8/start.sh

k8-stop:
	chmod +x deployment/k8/stop.sh
	./deployment/k8/stop.sh

helm-serve:
	chmod +x deployment/helm/start.sh
	./deployment/helm/start.sh

helm-stop:
	chmod +x deployment/helm/stop.sh
	./deployment/helm/stop.sh

migrate:
	./out/main migrate

rollback:
	./out/main rollback

display-config:
	./out/main config