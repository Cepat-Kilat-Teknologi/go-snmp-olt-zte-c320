# Monitoring OLT ZTE C320 with SNMP
Service for integration into the C320 OLT with the Go programming language

#### 👨‍💻 Full list what has been used:
* [Go](https://go.dev/) - Programming language
* [Chi](https://github.com/go-chi/chi/) - HTTP Server
* [GoSNMP](https://github.com/gosnmp/gosnmp) - SNMP library for Go
* [Redis](https://github.com/redis/go-redis/v9) - Redis client for Go
* [Zerolog](https://github.com/rs/zerolog) - Logger
* [Viper](https://github.com/spf13/viper) - Configuration management
* [Docker](https://www.docker.com/) - Containerization
* [Task](https://github.com/go-task/task) - Task runner
* [Air](https://github.com/cosmtrek/air) - Live reload for Go apps

#### 👨‍💻Recommendation for local development most comfortable usage:

``` shell
task dev
```

#### Docker development usage:
```shell
task up
```

```shell
docker-compose -f docker-compose.local.yaml up -d && air -c .air.toml
```


#### Available tasks for this project:

| Syntax             | Description                                                     |
|--------------------|-----------------------------------------------------------------|
| init               | Initialize the environment                                      |
| dev                | Start the local development                                     |
| app-build          | Build the app binary                                            |
| build-image        | Build docker image with tag latest                              |
| push-image         | push docker image with tag latest                               |
| pull-image         | pull docker image with tag latest                               |
| docker-run         | Run the docker container image with tag latest                  |
| docker-stop        | Stop the docker container                                       |
| docker-remove      | Remove the docker container                                     |
| up                 | Start the docker containers in the background                   |
| up-rebuild         | Rebuild the docker containers                                   |
| down               | Stop and remove the docker containers                           |
| restart            | Restart the docker containers                                   |
| rebuild            | Rebuild the docker image and up with detached mode              |
| tidy               | Clean up dependencies                                           |

#### Test with curl GET method Board 2 Pon 7
``` shell
curl -sS localhost:8081/api/v1/board/2/pon/7 | jq
```
#### Output Result
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "board": 2,
      "pon": 7,
      "onu_id": 3,
      "name": "Siti Khotimah",
      "onu_type": "F670LV7.1",
      "serial_number": "ZTEGCE3E0FFF",
      "rx_power": "-22.22",
      "status": "Online"
    },
    {
      "board": 2,
      "pon": 7,
      "onu_id": 4,
      "name": "Isroh",
      "onu_type": "F670LV7.1",
      "serial_number": "ZTEGCEEA1119",
      "rx_power": "-21.08",
      "status": "Online"
    },
    {
      "board": 2,
      "pon": 7,
      "onu_id": 5,
      "name": "Hadi Susilo",
      "onu_type": "F670LV7.1",
      "serial_number": "ZTEGCEC3033C",
      "rx_power": "-19.956",
      "status": "Online"
    }
  ]
}
```

#### Test with curl GET method Board 2 Pon 7 Onu 4
```shell
 curl -sS localhost:8081/api/v1/board/2/pon/7/onu/4 | jq
```

#### Output Result
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "board": 2,
    "pon": 7,
    "onu_id": 4,
    "name": "Isroh",
    "description": "Bale Agung",
    "onu_type": "F670LV7.1",
    "serial_number": "ZTEGCEEA1119",
    "rx_power": "-21.08",
    "tx_power": "2.5340000000000007",
    "status": "Online",
    "ip_address": "10.90.1.214"
  }
}
```

#### Test with curl GET method Get Empty ONU_ID in Board 2 Pon 5
```shell
curl -sS localhost:8081/api/v1/board/2/pon/5/onu_id/empty | jq
```

#### Output Result
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "board": 2,
      "pon": 5,
      "onu_id": 123
    },
    {
      "board": 2,
      "pon": 5,
      "onu_id": 124
    },
    {
      "board": 2,
      "pon": 5,
      "onu_id": 125
    },
    {
      "board": 2,
      "pon": 5,
      "onu_id": 126
    }
  ]
}
```

#### Test with curl GET method Get Empty ONU_ID After Add ONU in Board 2 Pon 5
```shell
curl -sS localhost:8081/api/v1/board/2/pon/5/onu_id/update | jq
```

```json
{
  "code": 200,
  "status": "OK",
  "data": "Success Update Empty ONU_ID"
}
```
