generate:
	docker build --target gen -t shop-proto-gen .
	docker run --name shop-proto-gen shop-proto-gen
	docker cp shop-proto-gen:/proto/gen/go ./gen
	docker rm shop-proto-gen

lint:
	docker build --target lint -t shop-proto-gen .