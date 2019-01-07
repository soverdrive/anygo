list:
	@grep '^[^#[:space:]].*:' Makefile
test:
	@echo " >> running tests"
build:
	@echo " >> building"
