.PHONY: clean
clean:
	rm -rf admin/v1/*.pb*.go admin status
	rm -rf api/v1/*.pb*.go admin status
	rm -rf metalstack/api/v1/*.pb*.go admin status

.PHONY: test
test:
	go test -v ./...
