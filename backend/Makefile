start-server:
	docker compose up -d
stop-server:
	docker compose down
restart-vocabs:
	docker-compose -p ai_learning_vocabs build my_vocabs
	docker-compose -p ai_learning_vocabs up my_vocabs -d

# Using curl to test API on local
add-word:
	curl -X POST http://localhost:8080/api/v1/words/add -d '{"word": "$w"}'
list-word:
	curl http://localhost:8080/api/v1/words/review
mark-complete:	
	curl -X POST http://localhost:8080/api/v1/words/complete -d '{"word": "$w"}'

