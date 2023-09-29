restart:
	@echo "Restarting the server..."
	@docker-compose down
	@docker rmi -f goctf-app
	@docker-compose up -d --build

log:
	@echo "Showing logs..."
	@docker logs -f goctf-app-1