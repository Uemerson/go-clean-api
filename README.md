# About

Clean Architecture with Go

Technologies used:

- Go
- Google UUID
- Crypto
- PostgreSQL
- Gorm

# Bash scripts

- The script [up.dev.v2.sh](./up.dev.v2.sh) contains bash instructions to bring up the application using `Docker v2` and stop the containers by pressing `Ctrl+C`.

# Instructions

This project uses [Docker and Docker Compose, so it's necessary to install them.](https://docs.docker.com/compose/install/)

After installation, rename the files `.env.dev.example` to `.env.dev`. 

- `.env.dev.example` containing environment variables with default values.
- `.env.dev`: It's used to set the environment variables used by the application.

Afterwards, you can bring up the containers using the bash script below:

```
$ bash up.dev.v2.sh
```

or if you don't want to use the bash script:

```
$ docker compose -f docker-compose.dev.yml --env-file .env.dev up -d --build --remove-orphans
```

After bringing up the application with the default values from .env.example, the API should be running at [localhost:8080](http://localhost:8080/). 

# Note

If you are using **Docker version 1 and Docker Compose version 1**, please replace `docker compose` with `docker-compose` in the bash scripts.

If you are on localhost, upon opening the API in the browser, you should see the **Django Debug Toolbar** on the side. It is used to display information about request times, executed SQL commands, and other useful information.