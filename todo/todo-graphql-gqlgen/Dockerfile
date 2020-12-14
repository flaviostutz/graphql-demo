FROM golang:1.15.6-alpine3.12 AS BUILD

WORKDIR /app

#cache modules
ADD /go.mod /app
ADD /go.sum /app
RUN go mod download

#now build source code
ADD / /app
RUN go build -x -o /todo-gqlgen



FROM golang:1.15.6-alpine3.12
EXPOSE 3000
ENV TODO_SERVICE_URL       ''
COPY --from=BUILD /todo-gqlgen /
CMD [ "/todo-gqlgen" ]
