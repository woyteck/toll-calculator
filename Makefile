obu:
	@go build -C obu -o ../bin/obu .
	@./bin/obu

receiver:
	@go build -C data_receiver -o ../bin/receiver .
	@./bin/receiver

.PHONY: obu
