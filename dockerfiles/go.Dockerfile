# https://docs.docker.com/engine/reference/builder/#understand-how-arg-and-from-interact
ARG GOLANG_VERSION=1.13.5
FROM golang:$GOLANG_VERSION

ARG GOLANG_VERSION
RUN echo $GOLANG_VERSION

WORKDIR /app

COPY ./.cache /.cache
COPY ./ /app

#RUN go mod download
#ENTRYPOINT ["./air"]

ENTRYPOINT ["/app/scripts/go_run.sh"]
