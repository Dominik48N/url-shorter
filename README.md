# URL Shorter

With this project, you have the possibility to shorten your URLs. In addition, the project offers the possibility of statistics. Environment variables can be used to specify which statistics are to be saved.

It is recommended to deploy the project in a Kubernetes cluster, corresponding manifest files can be found [here](manifests/).

The URLs then look like this: `example.com/aJhIleoaUc` *(The ID can be between 3 and 12 characters long.)*

## Requirements
* A [PostgreSQL](https://www.postgresql.org/) database
* A [Redis](https://redis.io/) cluster
* *(Optional)* A Kubernetes cluster

## Enviroment variables

Each service uses environment variables to e.g. to get the connection data of the database. If you use the Kubernetes Manifest files, then the required environment variables with example values ​​are already available. Then you only have to change the values. However, all environment variables for the individual services are listed here again if you do not use Kubernetes or do not use the provided manifest files.

### Query Service

* `POSTGRES_HOST` The hostname with port for the PostgreSQL database (e.g., `127.0.0.1:5432`)
* `POSTGRES_USERNAME` The username for the PostgreSQL database authentication
* `POSTGRES_PASSWORD` The password for the PostgreSQL database authentication
* `POSTGRES_DATABASE` The database name where the URLs are stored.
* `FALLBACK_URL` If no short URL is found, the user will be redirected to the website specified here.
  * If this environment variable is not set or is empty, you will not be redirected to a website, but only receive a message that no URL was found.
