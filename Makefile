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
	$(GO) run github.com/client9/misspell/cmd/misspell *

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
	$(GO) run github.com/jstemmer/go-junit-report/v2@latest < $< > $@

# Gitlab coverage report
coverage.xml: coverage.txt
	$(GO) run github.com/boumenot/gocover-cobertura@latest -cobertura < $< > $@

# Gitlab code quality report
gl-code-quality-report.json: staticcheck.json
	$(GO) run github.com/banyansecurity/golint-convert@latest -convert > $@

staticcheck.json:
	$(GO) run honnef.co/go/tools/cmd/staticcheck@latest -f json > $@

# Gitlab dependency report
govulncheck.sarif:
	$(GO) run golang.org/x/vuln/cmd/govulncheck@latest -format=sarif ./... > $@

.SUFFIXES: .peg .go

.peg.go:
	peg -noast -switch -inline -strict -output $@ $<

peg: grammar.go

.PHONY: total
total:
	awk -f total.awk < benches/all.txt
