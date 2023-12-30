FROM	golang:1.21.5

MAINTAINER	Jorge Araya Navarro <jorge@esavara.cr>
WORKDIR	/app
COPY	.	.
RUN	go build -o main ./cmd/reviewer
CMD	["./main"]
