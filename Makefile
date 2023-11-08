mock:
	rm -rf internal/core/port/mocks && \
	mockery --dir=internal/core/port --all --output=internal/core/port/mocks && \
	go mod tidy
