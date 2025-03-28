build:
	go build -o ccwc

bytes:
	go build -o ccwc && ./ccwc -c $(file)

run:
	go run .

clean:
	go clean