test: copy-config build migrate
	go test ./...

copy-config:
	cp application.yaml.sample application.yaml

test-cover-html:
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

migrate:
	./out/main migrate

rollback:
	./out/main rollback

display-config:
	./out/main config