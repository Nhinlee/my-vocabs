start-server:
	docker compose up -d
stop-server:
	docker compose down
restart-vocabs:
	docker-compose -p ai_learning_vocabs build my_vocabs
	docker-compose -p ai_learning_vocabs up my_vocabs -d

# Using curl to test API on local
local-add-word:
	curl -X POST http://localhost:8080/api/v1/words/add -d '{"word": "$w"}'
local-list-word:
	curl http://localhost:8080/api/v1/words/review
local-mark-complete:	
	curl -X POST http://localhost:8080/api/v1/words/complete -d '{"word": "$w"}'

# Using curl to test API on local
add-word:
	curl -X POST ${BASE_URL}/api/v1/words/add -d '{"word": "$w"}'
list-word:
	curl ${BASE_URL}/api/v1/words/review
mark-complete:	
	curl -X POST ${BASE_URL}/api/v1/words/complete -d '{"word": "$w"}'

# BUild & deploy
build:
	docker buildx build --platform linux/amd64 -t gcr.io/lla-app-395103/my-vocabs:latest -f ./cmd/my_vocabs/Dockerfile --push .
deploy:
	gcloud run deploy my-vocabs \
	--image gcr.io/lla-app-395103/my-vocabs:latest \
	--platform managed \
	--region us-central1 \
	--allow-unauthenticated
