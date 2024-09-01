obu:
	@go build -o bin/obu.exe ./obu
	@./bin/obu

reciver:
	@go build -o bin/receiver.exe ./data-receiver
	@./bin/receiver

.PHONY: obu
