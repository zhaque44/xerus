fmt:
	@echo "Formatting code...changing files, if no files are changed, then no formatting is needed"
	@gofmt -l -w .

test:
	go test -v ./...