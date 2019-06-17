run_web:
		go run cmd/consumer/main.go

run_container:
		docker run -p 8080:8080 consumer

run_database:
		docker run -p 5432:5432 db


run:
		docker-compose up

down:
		docker-compose down

clean:
		docker rm db consumer
build:	
		go build -o ./cmd/consumer/web-consumer -i ./cmd/consumer
		docker build -t web-consumer ./cmd/consumer
		rm ./cmd/consumer/web-consumer

test:
	go test -v --cover