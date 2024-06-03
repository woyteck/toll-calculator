obu:
	@go build -C obu -o ../bin/obu .
	@./bin/obu

receiver:
	@go build -C data_receiver -o ../bin/receiver .
	@./bin/receiver

calculator:
	@go build -C distance_calculator -o ../bin/calculator .
	@./bin/calculator

agg:
	@go build -C aggregator -o ../bin/agg .
	@./bin/agg

.PHONY: obu
