obu:
	@go build -C obu -o ../bin/obu main.go
	@./bin/obu

receiver:
	@go build -C data_receiver -o ../bin/receiver main.go
	@./bin/receiver

.PHONY: obu
