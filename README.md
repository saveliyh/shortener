# To run the project

1. Create a db.env file in the root directory of the project with variables:
    * POSTGRES_HOST
    * POSTGRES_PORT
    * POSTGRES_USER
    * POSTGRES_PASSWORD
    * POSTGRES_DB

2. run `docker compose up` to start the service

# Settings

Optional environment variable `STORAGE_TYPE` allows to choose storage for service. Possible values: `postgres` or `memmory`. Default value is `memmory`.
