# Hilfling - OAuth

OAuth server for Fotogjengen


## Development environment

To bring up dev docker container with hot-reloading
`docker-compose up --build dev`


## Production environment

To bring up production environment make sure you have a local external traefik network
(docker prompts you to do this and gives you a command if you don't already have it)

To bring up the prod container (build without hot reloading)
`docker-compose up --build prod`
