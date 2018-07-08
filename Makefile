run:
	docker run -d \
		--name linebot \
		-p 5566:5566 \
		-v /PATH/TO/SSL:/go/src/github.com/jiazhen-lin/linebot/ssl \
		linebot:$(VERSION)

run-production:
	docker-compose up -d

run-develop:
	docker-compose up -d

down-production:
	docker-compose down

down-develop:
	docker-compose down

build:
	docker build -t linebot:$(VERSION) .
