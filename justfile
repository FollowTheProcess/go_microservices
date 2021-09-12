_default:
    @just --list

# Runs the swagger tool to generate API docs
swagger:
    @echo Ensure you have the swagger CLI or this command will fail.
    @echo You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger
    @echo ....

    swagger generate spec -o ./swagger.yaml --scan-models

# Run the unit tests
test:
    go test -race ./...
