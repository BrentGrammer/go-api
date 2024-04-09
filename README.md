# Golang API

- A simple API using GoLang for reference
- Notes and examples from [this tutorial](https://www.youtube.com/watch?v=8uiZC0l4Ajw)

## Usage with Docker

- Build the image with: `docker build -t goapi:latest .`
- Run the docker container with: `docker run goapi:latest`
- Access API on exposed port 8000

### Docker compose

- `docker compose up --build`

## Using the API

- Requests can be sent to `http://localhost:8000`
  - `/health`
  - `/account/coins?username=alex)`
    - Set `Authorization` header to "123ABC"
