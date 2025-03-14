# Description: Makefile for my-vocabs project
start-server:
	docker compose up -d
stop-server:
	docker compose down
restart-vocabs:
	docker-compose -p ai_learning_vocabs build my_vocabs
	docker-compose -p ai_learning_vocabs up my_vocabs -d

# Using curl to test API on local
add-vocab:
	curl -X POST http://localhost:8080/api/v1/vocabs/add -d '{"word": "$w"}'
review-vocabs:
	curl http://localhost:8080/api/v1/vocabs/review
mark-complete:	
	curl -X POST http://localhost:8080/api/v1/vocabs/complete -d '{"word": "$w"}'

# Build & deploy -------------------------------------------------------------------------------------------------
build-prod:
	docker buildx build \
		--platform linux/arm64 \
		-t gcr.io/lla-app-395103/my-vocabs:latest \
		-f ./cmd/my_vocabs/Dockerfile \
		--build-arg ENV=prod \
		--push .
deploy-prod:
	gcloud run deploy my-vocabs \
	--image gcr.io/lla-app-395103/my-vocabs:latest \
	--platform managed \
	--region us-central1 \
	--allow-unauthenticated \
	--update-secrets DB_SOURCE_SECRET=my-vocabs-db-source:latest \
	--update-secrets GSA_KEY_SECRET=my-vocabs-gsa-key:latest \


# Database -------------------------------------------------------------------------------------------------
drop-db:
	docker exec -it postgres_db dropdb -U nhinlee my_vocabs

init-db:
	docker exec -it postgres_db createdb -U nhinlee my_vocabs

migrate-up:
	migrate -path ./internal/my_vocabs/db/migrations -database "postgresql://nhinlee:123456@localhost:5432/my_vocabs?sslmode=disable" -verbose up

migrate-up-prod:
	migrate -path ./internal/my_vocabs/db/migrations -database $$DB_PATH -verbose up

access-db:
	docker exec -it postgres_db psql -U nhinlee -d my_vocabs

access-db-prod:
	psql $$DB_PATH

reset-db:
	docker exec -it postgres_db psql -U nhinlee -d my_vocabs -c "DROP TABLE IF EXISTS schema_migrations;"
	make migrate-up

reset-db-prod:
	psql $$DB_PATH -c "DROP TABLE IF EXISTS schema_migrations;"
	make migrate-up-prod
