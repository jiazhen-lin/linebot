VERSION := 0.0.1

run:
	docker run -d \
		--name linebot \
		-p 8088:8088 \
		-v /home/linebot/ssl:/go/src/github.com/jiazhen-lin/linebot/ssl \
		linebot:$(VERSION)

build:
	docker build -t linebot:$(VERSION) .