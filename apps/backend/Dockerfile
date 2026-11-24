FROM golang:1.26-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /backend .

FROM scratch
COPY --from=build /backend /backend
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/backend"]
