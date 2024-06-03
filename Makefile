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

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto

.PHONY: obu
