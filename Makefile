
help:
	@echo "Available commands:"
	@echo "make install			Install dependencies."
	@echo "make test			Run tests."
	@echo "make coverage			Show coverage in html."
	@echo "make clean			Clean build files."

install:
	@echo "Make: Install"
	glide up

.PHONY: test
test:
	@echo "Make: Test"
	./scripts/test.sh

coverage:
	@echo "Make: Coverage"
	./scripts/coverage.sh

clean:
	@echo "Make: Clean"
	rm -rf vendor
	touch coverage.out && rm coverage.out
	touch coverage.html && rm coverage.html
