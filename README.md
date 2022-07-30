# Discord Bot with Markov

## Linux Steps

### First-time setup
 1. `touch .env`
 2. `docker-compose up registry`
 3. `docker build . -t localhost:5000/markov-discord-go/markov-discord-go:latest`
 4. `docker push --disable-content-trust localhost:5000/markov-discord-go/markov-discord-go:latest`

### Build setup
 1. `make devenv`
 2. `make discord-bot`

### Run setup
 1. Copy your `json` data to `data/`. Make sure it's a `.zip` file.
 2. `./dist/discord-bot start`
