#
# This file provides all common test targets.
#

.PHONY: test/unit
test/unit:
	go test -count=1 ./...

test/unit/verbose:
	go test -count=1 -v -cover ./...

test/unit/short:
	go test -count=1 -v -cover -short ./...

test/unit/cover:
	go test -count=1 -v -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

.PHONY: test/run
test/run:
	make -C ./test/run

test/help:
	@echo
	@echo '*** Test utility targets ***'
	@echo
	@echo 'Usage:'
	@echo '    make test/unit             run unit tests'
	@echo '    make test/unit/verbose     run unit tests with verbose level and coverage'
	@echo '    make test/unit/short       run unit tests with verbose level and coverage in short mode'
	@echo '    make test/unit/cover       run unit tests with coverage and start default broweser'
	@echo '    make test/run              run run tests'

