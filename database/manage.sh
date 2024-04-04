#!/bin/bash

# Build the Go program
build() {
    go build -o myprogram main.go
}

# Run the Go program
run() {
    ./myprogram
}

# Clean up generated files
clean() {
    rm myprogram
}

# Main function to handle command line arguments
main() {
    case "$1" in
        build)
            build
            ;;
        run)
            run
            ;;
        clean)
            clean
            ;;
        *)
            echo "Usage: $0 {build|run|clean}"
            exit 1
            ;;
    esac
}

main "$@"
