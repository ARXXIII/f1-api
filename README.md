# F1 API Service

This project is a RESTful API service built with Go, designed to provide data related to Formula 1. It offers endpoints to retrieve information about drivers, teams, races, and standings.

## Features

-   Retrieve F1 driver details.
-   Get team and constructor information.
-   Access race schedules and results.
-   View championship standings.

## Prerequisites

-   Go 1.20+ installed.
-   Docker and Docker Compose installed.
-   `make` utility installed.

## Running the Project

To run the project locally, use the provided `Makefile`. Follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/f1-api.git
    cd f1-api
    ```

2. Run the application:
    ```bash
    make run
    ```

This will start the API server on the default port (e.g., `http://localhost:8080`).

## Building and Running with Docker

To build and run the application using Docker Compose, follow these steps:

1. Build the Docker container:

    ```bash
    docker-compose build
    ```

2. Start the container:
    ```bash
    docker-compose up
    ```

The API will be accessible at `http://localhost:8080`.

## Makefile Commands

-   `make run`: Run the application locally.
-   `make build`: Build the Go binary.
-   `make test`: Run tests.

## Docker Compose Commands

-   `docker-compose build`: Build the Docker image.
-   `docker-compose up`: Start the application in a container.
-   `docker-compose stop`: Stop containers.
-   `docker-compose down`: Stop and remove containers.

## License

This project is licensed under the MIT License.
