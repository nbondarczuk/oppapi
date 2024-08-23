#
# This file provides all common technical go related targets.
#

go/tidy:
	go mod tidy -v

go/purge:
	go clean -modcache

go/fmt:
	go fmt ./...

go/vet:
	go vet ./...

go/lint:
	golangci-lint run

go/sec:
	gosec ./...

go/tar: clean
	tar -cvzf /tmp/$(TARGET).tar.gz .env .gitignore *

go/help:
	@echo
	@echo '*** Golang utility targets ***'
	@echo
	@echo 'Usage:'
	@echo '    make go/tidy                  run go tidy to resolve dependencies'
	@echo '    make go/purge                 clean dependencies cache'
	@echo '    make go/verify                run go verify'
	@echo '    make go/vet                   run go vet'
	@echo '    make go/linty                  run golangci-lint'
	@echo '    make go/tar                   run tar'
