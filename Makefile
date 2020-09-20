lint: ## Lint the files
	$(GOLINT) -set_exit_status ${PKG_LIST}

format:
	go fmt ./...

test:
	@go test -v ${PKG_LIST}

test-coverage:
	@go test -v -coverpkg=./... -coverprofile=profile.cov ./...
	@go tool cover -func profile.cov

mock:
	mockgen \
		-source=./gofmt256.go \
		-package gofmt256mocks \
		-destination=./mocks/gofmt256.go