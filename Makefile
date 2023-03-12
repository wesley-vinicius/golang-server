start: 
	@docker-compose up -d


run: 
	@docker build . --no-cache
	@docker-compose run app