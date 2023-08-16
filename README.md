# RSS Feed Aggregator for Blogs - Backend API

Welcome to the README for the RSS Feed Aggregator for Blogs backend API written in Go. This project allows you to create an RSS feed aggregator that collects and manages updates from various blogs and websites. You can follow your favorite blogs, automatically fetch their latest posts, and stay up-to-date with their content.

## Project Details

This project aims to build an RSS feed aggregator using Go, enabling users to keep track of their preferred blogs and websites that provide RSS feeds in XML format. Future plans include extending the aggregator to support news sites, podcasts, and various feed formats in which we can fetch posts, including JSON and creating a frontend for it.

## Features

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds added by other users
- Regularly fetch the latest posts from followed RSS feeds and save them in the database

## Getting Started

Follow these steps to get your development environment up and running.

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.
- [Git](https://git-scm.com/downloads) installed.
- [goose](https://github.com/pressly/goose) installed (database migration tool)
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) installed (generates sql query to type safe code)
- [Postgres](https://www.postgresql.org/download/) installed. (I used postgres but you can use any database and configure it accordingly. You may have to make certain changes in the main.go and .env file)

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/abhishekghimire40/blog-feed-aggregator.git
   cd blog-feed-aggregator
   ```

2. Install dependencies:

   ```bash
    go mod download
   ```

3. Install goose and sqlc:

   ```bash
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

4. Run database migrations:

   ```bash
   cd sql/schema/
   goose postgres postgres://username:password@host:port/databasename up
   ```

   **NOTE**: swap with your postgres database details

5. Build the project:

   ```bash
    go build
   ```

## Usage

### Configuration

Before running the API, make sure to set up the required configuration. You can either use environment variables or create a .env file in the root directory and populate it with the necessary values. An example .env file might look like this(i used postgres as the database):

```env
    PORT=8080
    DATABASE_URL=postgres://username:password@host:port/databasename?sslmode=disable
```

## API Endpoints

1. **Health Check** `/v1/healthz`

   Returns a `200` status code and a JSON response:

   ```json
   {
     "status": "ok"
   }
   ```

2. **Error Test** `GET /v1/err`

   Returns a `500` status code and a JSON response:

   ```json
   {
     "error": "Internal Server Error"
   }
   ```

3. **Create User** `POST /v1/users`

   Example request body:

   ```json
   {
     "name": "Lane"
   }
   ```

   Example response body:

   ```json
   {
     "id": "uuid",
     "created_at": "2021-09-01T00:00:00Z",
     "updated_at": "2021-09-01T00:00:00Z",
     "name": "John Doe",
     "api_key": "f0b4f72c8e9d1a6e87cd2d75409a3b64a185d6c92d8f7a22c8b7b063eef175d9"
   }
   ```

4. **Get Current User** `GET /v1/users`
   `Authentication Required`
   Request headers: `Authorization: ApiKey <key>`

   Example response body:

   ```json
   {
     "id": "7e94b29a-d87e-4c33-8a4b-1e6c55d7dabe",
     "created_at": "2022-11-15T08:32:45Z",
     "updated_at": "2023-05-19T17:20:10Z",
     "name": "Eleanor",
     "api_key": "f0b4f72c8e9d1a6e87cd2d75409a3b64a185d6c92d8f7a22c8b7b063eef175d9"
   }
   ```

5. **Create Feed** `POST /v1/feeds`
   `Authentication Required`

   When a user creates a new feed, they should automatically be following that feed. They can of course choose to unfollow it later, but it should be there by default.

   Request headers: `Authorization: ApiKey <key>`

   Example request body:

   ```json
   {
     "name": "Epic Insights",
     "url": "https://epicinsights.blog/random-feed.xml"
   }
   ```

   Example response body:

   ```json
   {
     "feed": {
       "id": "c6710e6f-7d4d-4ec1-8f59-8b97411f753b",
       "created_at": "2022-07-10T12:45:00Z",
       "updated_at": "2022-08-20T18:30:00Z",
       "name": "Epic Insights",
       "url": "https://epicinsights.blog/random-feed.xml",
       "user_id": "9b5a3ef3-8e18-42e6-916b-5a415e53b931"
     },
     "feed_follow": {
       "id": "f74a632d-1e39-42c6-a24a-834cf07a3e8f",
       "feed_id": "c6710e6f-7d4d-4ec1-8f59-8b97411f753b",
       "user_id": "582149c0-ae42-4e32-af43-15b3361c2379",
       "created_at": "2020-03-15T08:00:00Z",
       "updated_at": "2020-03-15T08:00:00Z"
     }
   }
   ```

6. **Get All Feeds** `GET /v1/feeds`

   Retrieves all feeds in the database without requiring authentication.

7. **Follow Feed** `POST /v1/feed_follows`
   `Authentication Required`
   Request headers:
   Authorization: ApiKey **Key**

   Example request body:

   ```json
   {
     "feed_id": "ba2f3e4d-6c5b-4a1a-bd17-89a6c745b7c9"
   }
   ```

   Example response body:

   ```json
   {
     "id": "f8d2e1c0-93a7-4b2e-b637-5f6d8c9e0a4f",
     "feed_id": "ba2f3e4d-6c5b-4a1a-bd17-89a6c745b7c9",
     "user_id": "7f1b8a9e-5d4c-3a2b-6e1f-8c0d9b4a5f6e",
     "created_at": "2023-08-15T15:45:00Z",
     "updated_at": "2023-08-15T15:45:00Z"
   }
   ```

8. **Unfollow Feed** `DELETE /v1/feed_follows/`
   Endpoint to unfollow a feed to not recieve blogs related to that feed
9. **Get User's Feed Follows** `GET /v1/feed_follows`

   `Authentication Required`
   Request headers: `Authorization: ApiKey <key>`

   Example response:

   ```json
   [
     {
       "id": "e9a5b18f-3dfc-42c9-8cf7-6e6a2c1e12a1",
       "feed_id": "7b0f4826-80e3-4da5-9e4b-85ca1e7f2f84",
       "user_id": "34e61a89-1db9-48c5-9a0f-d8e7b2c6a3f2",
       "created_at": "2022-07-15T08:30:00Z",
       "updated_at": "2022-07-15T08:30:00Z"
     },
     {
       "id": "b1a8c87d-4eb3-49e7-9e90-60c4a952e3d2",
       "feed_id": "c6b3a8f2-56d4-4f5e-bc89-1d2e3a4c5b6f",
       "user_id": "19e7f6d5-c4b3-2a1e-9f8d-7c6b5a4e3d2f",
       "created_at": "2022-07-15T08:30:00Z",
       "updated_at": "2022-07-15T08:30:00Z"
     }
   ]
   ```

10. **Get all posts from users followed feeds** `GET /v1/posts`
    `Authentication Required`
    Request headers: `Authorization: ApiKey <key>`
    Query params: `limit: integer value e.g 10`

    Example Reponse Body:

    ```json
    [
      {
        "id": "9b24a563-ec1c-4e4a-a37c-1a8d590fbd43",
        "created_at": "2022-05-12T08:20:37.449812Z",
        "updated_at": "2022-05-12T08:20:37.449813Z",
        "title": "Exploring the Unknown",
        "url": "https://example.com/exploration/",
        "description": "Embark on a journey of exploration and discovery, delving into the mysteries of the universe. From ancient civilizations to cutting-edge science, this blog 	invites you to expand your horizons.",
        "published_at": "2022-05-01T10:30:00Z",
        "feed_id": "a1b2c3d4-e5f6-7a8b-9c0d-e1f2a3b4c5d6"
      },
      {
        "id": "d2e3f4a5-b6c7-8d9e-0f1a-2b3c4d5e6f7",
        "created_at": "2022-07-08T14:15:20.678901Z",
        "updated_at": "2022-07-08T14:15:20.678902Z",
        "title": "Unraveling the Enigma",
        "url": "https://example.com/unraveling-enigma/",
        "description": "Dive into the depths of complex puzzles and unravel enigmas that challenge the mind. Explore the art of critical thinking and problem-solving in this intriguing journey.",
        "published_at": "2022-06-15T18:45:00Z",
        "feed_id": "x0y1z2a3-b4c5-6d7e-8f9a-b1c2d3e4f5"
      }
    ]
    ```

## Contributing

I welcome contributions from the community! To contribute to this project, follow these steps:

    1. Fork this repository.
    2. Create a new branch for your feature/bugfix: git checkout -b feature/your-feature-name.
    3. Make your changes and commit them: git commit -m "Add your changes".
    4. Push the changes to your fork: git push origin feature/your-feature-name.
    5. Open a pull request describing your changes.

> **NOTE**:
> You can also fork this repository and make it your own project by adding new features and show it on your github.
