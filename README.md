# project-optican
Personal API project for tracking health stats

# Getting Started
You'll need to create a .env file for local development options. The variables needed are:
```
DB_USER=optican
DB_PASSWORD=optican
DB_HOST=localhost
APP_PORT=8080
```
Set as appropriate, then run the project

# Running the Project
Running docker compose will create two containers, one for the development database and one for the api
The api is available on port `3000` and the database is exposed for development purposes on port `5432`
Run `docker-compose up`

Once the docker containers are running, it will be possible to use regular go debugging techniques
while connected to the database. This includes using `go run`, but it will be available on whatever `APP_PORT`
you've chosen
