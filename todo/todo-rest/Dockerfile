FROM golang:1.15.6-alpine3.12 AS BUILD

WORKDIR /app

#cache modules
ADD /go.mod /app
ADD /go.sum /app
RUN go mod download

#now build source code
ADD / /app
RUN go build -x -o /todo-service



FROM golang:1.15.6-alpine3.12
EXPOSE 3000
ENV BASE_LOCATION_URL       ''
COPY --from=BUILD /todo-service /
ADD startup.sh /
CMD [ "/startup.sh" ]
