.PHONY: test cover

test: ##@TestSuite Run all tests and writes cov.out coverage
	go test -timeout 30s -v ./... -cover -coverprofile=cov.out

cover: ##@TestSuite Open coverage tool in browser
	go tool cover -html=cov.out

.PHONY: doc

doc: ##@Documentation generates docs
	echo "Not implemented"