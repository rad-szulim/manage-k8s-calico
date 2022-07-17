FROM amr-registry.caas.intel.com/sed-registry/library/golang:1.17 As build
WORKDIR /workspace
# Copy go.mod, go.sum files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy sources to the working directory
COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
  go build -ldflags="-s -w -extldflags=-static" -o build/manage-calico ./cmd

FROM scratch
COPY --from=build /workspace/build/manage-calico /app
ENTRYPOINT [ "/app" ]