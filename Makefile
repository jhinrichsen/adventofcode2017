GO ?= CGO_ENABLED=0 go

.PHONY: all
all: clean lint test

.PHONY: clean
clean:
	$(GO) clean # remove test results from previous runs so that tests are executed
	-rm \
		coverage.txt \
		coverage.xml \
		gl-code-quality-report.json \
		govulncheck.sarif \
		junit.xml \
		staticcheck.json \
		test.log

.PHONY: bench
bench:
	$(GO) test -run=^$ -bench=. -benchmem

.PHONY: lint
lint:
	test -z $(gofmt -l .)
	$(GO) vet
	$(GO) run honnef.co/go/tools/cmd/staticcheck@latest -version
	$(GO) run honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: prof
prof: cpu.profile
	$(GO) tool pprof $^

.PHONY: test
test:
	$(GO) test -run=Day -short -vet=all

.PHONY: sast
sast: coverage.xml gl-code-quality-report.json govulncheck.sarif junit.xml

coverage.txt test.log &:
	-$(GO) test -coverprofile=coverage.txt -covermode count -short -v | tee test.log

# Gitlab test report
junit.xml: test.log
	which go-junit-report || $(GO) install github.com/jstemmer/go-junit-report/v2@latest
	go-junit-report -version
	go-junit-report < $< > $@

# Gitlab coverage report
coverage.xml: coverage.txt
	which gocover-cobertura || $(GO) install github.com/boumenot/gocover-cobertura@latest
	gocover-cobertura < $< > $@

# Gitlab code quality report
gl-code-quality-report.json: staticcheck.json
	which golint-convert || $(GO) install github.com/banyansecurity/golint-convert@latest
	golint-convert > $@

staticcheck.json:
	which staticcheck || $(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -version
	-staticcheck -f json > $@

# Gitlab dependency report
govulncheck.sarif:
	which govulncheck || $(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck -version
	govulncheck -format=sarif ./... > $@


.SUFFIXES: .peg .go

.peg.go:
	peg -noast -switch -inline -strict -output $@ $<

peg: grammar.go

.PHONY: total
total:
	awk -f total.awk < benches/all.txt
