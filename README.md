# URL Shorter

With this project, you have the possibility to shorten your URLs. In addition, the project offers the possibility of statistics. Environment variables can be used to specify which statistics are to be saved.

The URLs then look like this: `example.com/aJhIleoaUc` *(The ID can be between 3 and 12 characters long.)*

## Requirements
* A [PostgreSQL](https://www.postgresql.org/) database
* A [Redis](https://redis.io/) cluster
* *(Optional)* A Kubernetes cluster

You need at least a PostgreSQL database, but you can also use multiple PostgreSQL databases. We have SQL files that you need to run once when setting up your database. You are free to run each file in an extra database or in one database.
Of course, we always recommend the variant of working with several PostgreSQL databases, for reasons of security and scalability. You can find the SQL files [here](db/).

## Enviroment variables

Each service uses environment variables to e.g. to get the connection data of the database. If you use the Kubernetes Manifest files, then the required environment variables with example values ​​are already available. Then you only have to change the values. However, all environment variables for the individual services are listed here again if you do not use Kubernetes or do not use the provided manifest files.

### Query

* `POSTGRES_HOST` The hostname with port for the PostgreSQL database (e.g., `127.0.0.1:5432`)
* `POSTGRES_USERNAME` The username for the PostgreSQL database authentication
  * Only a read-only user is required!
* `POSTGRES_PASSWORD` The password for the PostgreSQL database authentication
* `POSTGRES_DATABASE` The database name where the URLs are stored.
* `REDIS_HOSTS` The hosts of the Redis Cluster nodes separated by commas. (e.g., `127.0.0.1:6379,127.0.0.1:6370`)
* `REDIS_PASSWORD` The Redis Cluster password for authentication
* `URL_CACHING_TIME` The duration in seconds that URLs should be cached in Redis.
  * If the environment variable is not set, the time is set to 180 seconds.
* `FALLBACK_URL` If no short URL is found, the user will be redirected to the website specified here.
  * If this environment variable is not set or is empty, you will not be redirected to a website, but only receive a message that no URL was found.
* `HTTP_PORT` This is the port on which the HTTP server is running.
  * If this environment variable is not passed, port `3000` is used.

### Creator
* `POSTGRES_HOST` The hostname with port for the PostgreSQL database (e.g., `127.0.0.1:5432`)
* `POSTGRES_USERNAME` The username for the PostgreSQL database authentication
  * Write access is required!
* `POSTGRES_PASSWORD` The password for the PostgreSQL database authentication
* `POSTGRES_DATABASE` The database name where the URLs are stored.
* `HTTP_PORT` This is the port on which the HTTP server is running.
  * If this environment variable is not passed, port `3000` is used.

### Users

* `POSTGRES_HOST` The hostname with port for the PostgreSQL database (e.g., `127.0.0.1:5432`)
* `POSTGRES_USERNAME` The username for the PostgreSQL database authentication
  * Write access is required!
* `POSTGRES_PASSWORD` The password for the PostgreSQL database authentication
* `POSTGRES_DATABASE` The database name where the users are stored.
* `HTTP_PORT` This is the port on which the HTTP server is running.
  * If this environment variable is not passed, port `3000` is used.
* `SESSION_TIME` The time in minutes how long a JWT token should be valid.
  * By default, the time is 60 minutes.
* `JWT_SECRET` The secret key for JWT.
  * This should be secure and must be identical in all URL-Shorter services.

## Kubernetes

If you want to run URL Shorter in a Kubernetes cluster, you can use the provided [manifest files](manifests/).

Depending on how many users you expect, you have to adapt the manifest files.
* The replicas are only example values
* You have to adjust the values ​​of the environment variables
* An extra load balancer is used for many services. If you don't expect so many users, a load balancer for all services can make more sense.
