# #  docker build -t hello:v1 .
# # docker run -p 8080:80 hello:v1



# FROM golang:1.17.2

# WORKDIR /week-6-assignment-davuttrg
# COPY go.mod .
# COPY go.sum .

# COPY . .
# RUN go mod download
# RUN GOOS=linux CGO_ENABLED=0 go build -o hello

# CMD ["./hello"]

#--------------------------------------------

FROM golang:1.17.2 as build

WORKDIR /week-6-assignment-davuttrg
COPY go.mod .
COPY go.sum .

COPY . .
RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -o /hello


FROM alpine
COPY --from=build /hello /hello

CMD ["./hello"]

#docker tag hello:v3 davuttrg/hello:latest
#docker push davuttrg/hello:latest
