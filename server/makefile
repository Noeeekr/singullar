DATABASE_ENV_PATH=./secrets/postgres.env

NETWORK=internal
PORT=80

.PHONY: run-test-container
run-test-container:
	-docker network create internal
	docker run \
	--net=$(NETWORK) \
	-it --rm \
	-v $(PWD):/go/singullar \
	-p $(PORT):80 \
	test-container:1.0.0

.PHONY: build-test-container 
build-test-container:
	$(eval include $(DATABASE_ENV_PATH))
	docker build -t test-container:1.0.0 .

.PHONY: run-database-container
run-database-container: 
	$(eval include $(DATABASE_ENV_PATH))
	@echo Executing: Running postgres container
	@echo With User: $(POSTGRES_USER)
	@echo With Password: $(POSTGRES_PASSWORD)
	@echo With Container Name: $(PG_CONTAINER_NAME)
	@echo With volume: $(PG_CONTAINER_VOLUME)
	@echo On Image: Postgres:17.5
	-docker network create internal
	docker volume create "$(PG_CONTAINER_VOLUME)"
	docker run \
	--name "$(PG_CONTAINER_NAME)" \
	-v $(PG_CONTAINER_VOLUME):/var/lib/postgresql/data \
	-p 5432:5432 \
	--network internal \
	-e POSTGRES_USER="$(POSTGRES_USER)" \
	-e POSTGRES_PASSWORD="$(POSTGRES_PASSWORD)" \
	-d postgres:17.5

.PHONY: stop-database-container 
stop-database-container: 
	$(eval include $(DATABASE_ENV_PATH))
	@echo Stopping container $(PG_CONTAINER_NAME)
	@docker container stop $(PG_CONTAINER_NAME)

.PHONY: start-database-container 
start-database-container: 
	$(eval include $(DATABASE_ENV_PATH))
	@docker container start $(PG_CONTAINER_NAME)
	
.PHONY: delete-database-container
delete-database-container: docker-stop-database
	$(eval include $(DATABASE_ENV_PATH))
	docker network rm internal
	@echo Deleting container $(PG_CONTAINER_NAME)
	@docker container rm $(PG_CONTAINER_NAME)