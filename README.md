# Back Citate Generator

This project is a backend service for generating citations. It provides an API to create, retrieve, and manage citations programmatically.

## Features

- Generate citations in various formats.
- RESTful API for easy integration.
- Lightweight and fast.

## RUNTIME
### Dev run

1. Clone the repository:
    ```bash
    git clone https://github.com/nrf24l01/back_citate_gen.git
    ```
2. Navigate to the project directory:
    ```bash
    cd back_citate_gen
    ```
3. Install dependencies:
    ```bash
    go mod download
    ```
4. Run in dev mode
    ```bash
    go run main.go
    ```

### Prod
#### Just docker
```bash
docker pull ghcr.io/nrf24l01/back_citate_gen/backend:latest
```
#### Docker compose
```bash
docker compose up --build
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.