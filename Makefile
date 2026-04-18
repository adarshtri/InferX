# InferX Docker-First Orchestration

.PHONY: run stop logs clean help

# Default target: build and start the server
run:
	SCENARIO=$(SCENARIO) docker-compose up --build

# Stop the server and clean up containers
stop:
	docker-compose down

# Tail the server logs
logs:
	docker-compose logs -f

# Nuclear clean: Remove containers, volumes, and images
clean:
	docker-compose down -v --rmi all
	rm -rf lib/

# Help menu
help:
	@echo "InferX Management Commands:"
	@echo "  make run         - Build and start the hybrid server in Docker"
	@echo "  make stop        - Stop the server and containers"
	@echo "  make logs        - Tail the server logs"
	@echo "  make clean       - Remove all Docker artifacts and local /lib"
