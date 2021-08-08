init:
	go install honnef.co/go/tools/cmd/staticcheck@latest

lint:
	test -z $(gofmt -l .)
	golint -set_exit_status
	go vet
	staticcheck
