# GoURL - URL Shortener

GoURL is a URL shortener built with Golang, PostgreSQL, Redis, and Svelte. It provides a simple and efficient way to shorten long URLs into short, memorable ones.

## Features

- Shorten long URLs into short, easy-to-remember ones
- Redirect users to the original URL when they access the shortened URL
- Track the number of clicks and basic analytics for each shortened URL
- Dockerized application for easy deployment

## Technologies Used

- **Golang**: Backend server and API development
- **PostgreSQL**: Database management system
- **Redis**: In-memory data store for caching and analytics
- **Svelte**: Frontend framework for building the user interface

## Prerequisites

Before running the application, ensure that you have the following dependencies installed on your system:

- Go (1.16 or later)
- PostgreSQL
- Redis

## Getting Started

These instructions will help you set up the project and get it running on your local machine.

1. Clone the repository:
```bash
git clone https://github.com/your-username/your-repo.git
```

2. Navigate to the project directory:
```bash
cd your-repo
```

3. Start the Docker containers using Docker Compose:
```bash
docker-compose up -d
```



## Usage

1. Open the GoURL application in your web browser.

2. Enter a long URL that you want to shorten in the input field.

3. Click the "Shorten" button to generate a short URL.

4. Copy the shortened URL and share it with others.

5. When users access the shortened URL, they will be redirected to the original long URL.

6. To view analytics for a shortened URL, append `/stats` to the URL. For example:


5. Access the application in your web browser at `http://localhost:8000`.

## Dockerization

This project includes a `Dockerfile` and `docker-compose.yml` file to enable easy Dockerization and deployment.

- The `Dockerfile` contains the instructions to build a Docker image for the GoURL application.

- The `docker-compose.yml` file defines the services required for running the application, including the GoURL server, PostgreSQL database, and Redis cache.

To deploy the application using Docker, follow the steps mentioned in the "Getting Started" section above.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [Svelte](https://svelte.dev/)

## Contributing

Contributions are welcome! If you find any issues or want to enhance the application, please submit a pull request.
