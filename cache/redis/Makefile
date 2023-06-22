
CONTAINER_NAME = redis-stack

IMAGE_NAME = redis/redis-stack:latest

run:
	@docker run -d --name $(CONTAINER_NAME) -p 6379:6379 -p 8001:8001 $(IMAGE_NAME) > /dev/null 2>&1
	@echo "Redis Stack container started and running"

cli:
	@docker exec -it $(CONTAINER_NAME) redis-cli
	
stop:
	@docker stop $(CONTAINER_NAME) > /dev/null 2>&1
	@echo "Redis Stack container stopped"

delete:
	@docker rm -f $(CONTAINER_NAME) > /dev/null 2>&1
	@echo "Redis Stack container deleted"

clean: stop delete