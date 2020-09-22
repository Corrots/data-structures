#!/bin/zsh
container_name='redis-dev'
container_id=$(docker ps -aq -f name=${container_name})
docker rm -f "${container_id}"
docker run -d --privileged=true -p 6379:6379 \
  -v /home/docker/redis/conf/redis.conf:/etc/redis/redis.conf \
  -v /home/docker/redis/data:/data \
  --name myredis \
  redis:6.0-alpine \
  redis-server /etc/redis/redis.conf --appendonly yes

docker ps -a -f name=${container_name}
