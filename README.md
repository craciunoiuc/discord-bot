# Discord Bot with Markov

## Build setup

 0. `touch .env`
 1. `docker-compose up registry`
 2. `docker build . -t localhost:5000/markov-discord-go/markov-discord-go:latest`
 3. `docker push --disable-content-trust localhost:5000/markov-discord-go/markov-discord-go:latest`
 4. `make devenv`
 5. `make build`

## Run setup

 1. Copy your `json` data to `data/`
 2. `./dist/discord-bot start`