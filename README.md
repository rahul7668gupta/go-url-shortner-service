# go-url-shortner-service
A simple url shortner service built in golang

# URL Shortener Service

A simple URL shortener service built in Golang.

## Features

* `URL shortening`: The service allows users to shorten long URLs into shorter, more manageable links.
* `Custom alias`: Users can specify a custom alias for their shortened URLs.
* `URL redirection`: The service redirects users from the shortened URL to the original URL.
* `URL analytics`: The service provides analytics for each shortened URL, including the number of clicks and the date and time of each click.
* `RESTful API`: The service provides a RESTful API for creating and managing shortened URLs.
* `Caching support`: The service uses a Redis cache to store shortened URLs and their corresponding original URLs.
* `Docker support`: The service can be easily deployed using Docker.

## Getting Started

To get started with the service, follow these steps:

1. Clone the repository:
    ```
    git clone https://github.com/rahul7668gupta/go-url-shortner-service.git
    ```

2. Export env variable:
    ```
    export PORT=:8080
    export REDIS_ENDPOINT=localhost:6379
    export REDIS_DB=
    export REDIS_PASSWORD=0
    export SHORT_URL_DOMAIN=http://localhost:8080/
    ```

3. Run redis locally on port `6379`

4. Run the service:
    ```
    go run main.go
    ```

5. Access the service at <http://localhost:8080>.

## Build and Run using Docker

To build the Docker image, run the following command in the root directory of the project:
```
docker build -t url-shortener .
```

This command starts the container with the `PORT` environment variable set to `8081`.  You can then access the URL shortener service at <http://localhost:8081>.
```
docker run -p 8080:8080 -e PORT=8081 url-shortener
```

The service can be configured using environment variables mentioned in #Configuration:

You can set these environment variables using the `-e` flag when running the Docker container.

## API Endpoints

* `/shorten`: `POST` Creates a new shortened URL, if asked again ask for the same URL, it returns the old short url. Payload: 
```
{
    "url":"https://amazon.in"
}
```
* `/redirect/{code}`: `GET` Redirects to the original URL for the given code.
* `/metrics`: `GET` Returns top 3 domain names that have been shortened the most
number of times

## Configuration

The service can be configured using environment variables:

* `PORT`: The port number to listen on (recommened for local: `:8080`).
* `REDIS_ENDPOINT`: The database host (recommened for local: `localhost:6379`).
* `REDIS_DB`: The database port (recommened for local: "").
* `REDIS_PASSWORD`: The database password (recommened for local: `0`).
* `SHORT_URL_DOMAIN`: The custom domain to be used for generated short url(recommened for local: `http://localhost:8080/`).

## Running the Service

To run the service, you need to have Golang and Redis installed. You also need to have Docker installed if you want to use Docker to deploy the service.

## Contributing

Contributions are welcome! If you'd like to contribute to the project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Create a pull request.

## License

This project is licensed under the MIT License.