FROM public.ecr.aws/docker/library/golang:alpine AS build

WORKDIR /app
COPY * .
ENV CGO_ENABLED=0
RUN go build -o health .

FROM public.ecr.aws/docker/library/alpine:latest
COPY --from=build /app/health /usr/local/bin

ENTRYPOINT ["health"]
