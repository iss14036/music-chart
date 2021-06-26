# Music Chart

`Music Chart` is List of music based on user preference.

## Prerequisite

To run this program, you will need

### App Dependencies

```$xslt
- Golang 1.10+
- Go mod enabled
- docker
```

## How to Run

- Copy environment file from `config.env.example` to be `config.env` or use `ENVIRONMENT VARIABLE` directly
- Verify and download dependencies `make dep`
- Run the docker to serve database `make docker-up`

## How to Test

`make test`


### API List

You can import API to your vendor by the json file `music_chart.postman_collection.json`
