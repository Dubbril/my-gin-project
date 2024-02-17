FROM golang:1.21.6-bullseye AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /bin/app ./app

#FROM debian:buster-slim
#
#COPY --from=build /bin/app /bin
#
#EXPOSE 8080
#
#CMD [ "/bin/app" ]

# Use a minimal RHEL 8 base image
FROM registry.redhat.io/ubi8/ubi-minimal:8.9

COPY --from=build /bin/app /bin

EXPOSE 8080

CMD [ "/bin/app" ]