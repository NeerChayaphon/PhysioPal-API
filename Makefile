include .env
BINARY_NAME=physiopal

build:
	docker build -t physiopal \
	--build-arg MONGODB_LOCAL_URI="${MONGODB_LOCAL_URI}" \
	--build-arg MONGODB_STAGING_URI="${MONGODB_STAGING_URI}" \
	--build-arg REDIS_LOCAL_URI="${REDIS_LOCAL_URI}" \
	--build-arg REDIS_LOCAL_PASSWORD="${REDIS_LOCAL_PASSWORD}" .
	

build-p:
	docker build -t physiopal \
	--build-arg MONGODB_STAGING_URI="${MONGODB_STAGING_URI}" \
	--build-arg MONGODB_PRODUCTION_URI="${MONGODB_PRODUCTION_URI}" \
	--build-arg REDIS_LOCAL_URI="${REDIS_LOCAL_URI}" \
	--build-arg REDIS_LOCAL_PASSWORD="${REDIS_LOCAL_PASSWORD}" .

build-s:
	docker build -t physiopal \
	--build-arg MONGODB_STAGING_URI="${MONGODB_STAGING_URI}" \
	--build-arg REDIS_LOCAL_URI="${REDIS_LOCAL_URI}" \
	--build-arg REDIS_LOCAL_PASSWORD="${REDIS_LOCAL_PASSWORD}" .

run:
	docker run -p 8080:8080 physiopal

run-s:
	env MONGODB_STAGING_URI="${MONGODB_STAGING_URI}" go run main.go
