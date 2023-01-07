BINARY_NAME=physiopal

build:
	docker build -t physiopal \
	--build-arg MONGODB_STAGING_URI="mongodb+srv://neer:neer1234@physiopalstaging.g5mywjg.mongodb.net/?retryWrites=true&w=majority" \
	--build-arg REDIS_LOCAL_URI="rediss://physiopal.redis.cache.windows.net:6380" \
	--build-arg REDIS_LOCAL_PASSWORD="yAL0wvf5h3F3Fw7O2CBBjYvNZk9ANJqqpAzCaN8NIBo=" .

run:
	docker run -p 8080:8080 physiopal
