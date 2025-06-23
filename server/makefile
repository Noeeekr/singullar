.PHONY: docker-run-database

PMI_ENV_PATH=./pkg/pmi/postgres.env

docker-run-database: 
	$(eval include $(PMI_ENV_PATH))
	@echo Executing: Running postgres container
	@echo With User: $(POSTGRES_USER)
	@echo With Password: $(POSTGRES_PASSWORD)
	@echo With Container Name: $(CONTAINER_NAME)
	@echo With volume: $(CONTAINER_VOLUME)
	@echo On Image: Postgres:17.5
	docker volume create "$(CONTAINER_VOLUME)"
	docker run \
	--name "$(CONTAINER_NAME)" \
	-v $(CONTAINER_VOLUME):/var/lib/postgresql/data \
	-p 5432:5432 \
	-e POSTGRES_USER="$(POSTGRES_USER)" \
	-e POSTGRES_PASSWORD="$(POSTGRES_PASSWORD)" \
	-d postgres:17.5

docker-stop-database: 
	$(eval include $(PMI_ENV_PATH))
	@echo Stopping container $(CONTAINER_NAME)
	@docker container stop $(CONTAINER_NAME)

docker-start-database: 
	@docker container start $(CONTAINER_NAME)

docker-delete-database: docker-stop-database
	$(eval include $(PMI_ENV_PATH))
	@echo Deleting container $(CONTAINER_NAME)
	@docker container rm $(CONTAINER_NAME)