mock:
	rm -rf internal/core/port/mocks && \
	mockery --dir=internal/core/port --all --output=internal/core/port/mocks && \
	mockery --dir=internal/core/cus-service --all --inpackage && \
	go mod tidy
