docker-build-local: ## Build the docker image locally
	@docker build -t go-test-image .

docker-run-local: ## Runs the locally build docker image (forwards port 8080)
	@docker run --rm -p="8080:8080" --name go-test-container go-test-image

docker-stop-local: 
	@docker stop go-test-container