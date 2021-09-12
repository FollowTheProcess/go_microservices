_default:
    @just --list

# Runs the swagger tool to generate API docs
swagger:
    swagger generate spec -o ./swagger.yaml --scan-models

# Run the unit tests
test:
    go test -race ./...


# Start the server
run:
    go run main.go