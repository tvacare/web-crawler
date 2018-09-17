build: 
	docker build -t web-crawler .

start: 
	docker-compose up -d && docker logs -f web-crawler

stop: 
	docker-compose down