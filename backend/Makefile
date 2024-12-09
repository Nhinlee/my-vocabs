start-server:
	docker compose up -d
stop-server:
	docker compose down
restart-vocabs:
	docker-compose -p ai_learning_vocabs build my_vocabs
	docker-compose -p ai_learning_vocabs up my_vocabs -d

# Using curl to test API on local
local-add-word:
	curl -X POST http://localhost:8080/api/v1/vocabs/add -d '{"word": "$w"}'
local-list-word:
	curl http://localhost:8080/api/v1/vocabs/review
local-mark-complete:	
	curl -X POST http://localhost:8080/api/v1/vocabs/complete -d '{"word": "$w"}'

# Using curl to test API on local
add-word:
	curl -X POST ${BASE_URL}/api/v1/vocabs/add -d '{"word": "$w"}'
list-word:
	curl ${BASE_URL}/api/v1/vocabs/review
mark-complete:	
	curl -X POST ${BASE_URL}/api/v1/vocabs/complete -d '{"word": "$w"}'

# BUild & deploy
build:
	docker buildx build --platform linux/amd64 -t gcr.io/lla-app-395103/my-vocabs:latest -f ./cmd/my_vocabs/Dockerfile --push .
deploy:
	gcloud run deploy my-vocabs \
	--image gcr.io/lla-app-395103/my-vocabs:latest \
	--platform managed \
	--region us-central1 \
	--allow-unauthenticated \
	--update-secrets DB_SOURCE_PROD=my-vocabs-db-source:latest

# drop-db:
# 	docker exec -it postgres_db dropdb -U nhinlee my_vocabs

migrate-up:
	migrate -path ./internal/my_vocabs/db/migrations -database "postgresql://nhinlee:123456@localhost:5432/my_vocabs?sslmode=disable" -verbose up

migrate-up-prod:
	migrate -path ./internal/my_vocabs/db/migrations -database $$DB_PATH -verbose up

access-db:
	docker exec -it postgres_db psql -U nhinlee -d postgres

access-db-prod:
	psql $$DB_PATH