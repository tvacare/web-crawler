build: 
	docker build -t web-crawler .

start: 
	docker run -it web-crawler