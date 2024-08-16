# Monitoring OLT ZTE C320 with SNMP
[![build](https://github.com/megadata-dev/go-snmp-olt-zte-c320/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/megadata-dev/go-snmp-olt-zte-c320/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/megadata-dev/go-snmp-olt-zte-c320)](https://goreportcard.com/report/github.com/megadata-dev/go-snmp-olt-zte-c320)
[![codecov](https://codecov.io/gh/megadata-dev/go-snmp-olt-zte-c320/graph/badge.svg?token=NB3N7GMUX3)](https://codecov.io/gh/megadata-dev/go-snmp-olt-zte-c320)

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


#### Note : This service is still in development ⚠️👨‍💻👩‍💻

## Getting Started 🚀

### 👨‍💻Recommendation for local development most comfortable usage:

``` shell
task dev
```

### Docker development usage:
```shell
task up
```

```shell
docker-compose -f docker-compose.local.yaml up -d && air -c .air.toml
```

### Production usage with internal redis in docker:
```shell
task docker-run
```
```shell
docker network create local-dev && \
docker run -d --name redis-container \
--network local-dev -p 6379:6379 redis:7.2 && \
docker run -d -p 8081:8081 --name go-snmp-olt-zte-c320 \
--network local-dev -e REDIS_HOST=redis-container \
-e REDIS_PORT=6379 -e REDIS_DB=0 \
-e REDIS_MIN_IDLE_CONNECTIONS=200 -e REDIS_POOL_SIZE=12000 \
-e REDIS_POOL_TIMEOUT=240 -e SNMP_HOST=x.x.x.x \
-e SNMP_PORT=161 -e SNMP_COMMUNITY=xxxx \
sumitroajiprabowo/go-snmp-olt-zte-c320:latest
```

### Production usage without external redis:
```shell
docker run -d -p 8081:8081 --name go-snmp-olt-zte-c320 \
-e REDIS_HOST=redis_host \
-e REDIS_PORT=redis_port \
-e REDIS_DB=redis_db \
-e REDIS_MIN_IDLE_CONNECTIONS=redis_min_idle_connection \
-e REDIS_POOL_SIZE=redis_pool_size \
-e REDIS_POOL_TIMEOUT=redis_pool_timeout \
-e SNMP_HOST=snmp_host \
-e SNMP_PORT=snmp_port \
-e SNMP_COMMUNITY=snmp_community \
sumitroajiprabowo/go-snmp-olt-zte-c320:latest
```


### Available tasks for this project:

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

### Test with curl GET method Board 2 Pon 7
``` shell
curl -sS localhost:8081/api/v1/board/2/pon/7 | jq
```
### Result
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

### Test with curl GET method Board 2 Pon 7 Onu 4
```shell
 curl -sS localhost:8081/api/v1/board/2/pon/7/onu/4 | jq
```

### Result
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
    "rx_power": "-20.71",
    "tx_power": "2.57",
    "status": "Online",
    "ip_address": "10.90.1.214",
    "last_online": "2024-08-11 10:09:37",
    "last_offline": "2024-08-11 10:08:35",
    "uptime": "5 days 13 hours 10 minutes 50 seconds",
    "last_down_time_duration": "0 days 0 hours 1 minutes 2 seconds",
    "offline_reason": "PowerOff",
    "gpon_optical_distance": "6701"
  }
}
```

### Test with curl GET method Get Empty ONU_ID in Board 2 Pon 5
```shell
curl -sS localhost:8081/api/v1/board/2/pon/5/onu_id/empty | jq
```

### Result
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

### Test with curl GET method Get Empty ONU_ID After Add ONU in Board 2 Pon 5
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

### Test with curl GET method Get Onu Information in Board 2 Pon 8 with paginate
```shell
curl -sS 'http://localhost:8081/api/v1/paginate/board/2/pon/8?limit=3&page=2' | jq
```
### Result
```json
{
  "code": 200,
  "status": "OK",
  "page": 2,
  "limit": 3,
  "page_count": 23,
  "total_rows": 69,
  "data": [
    {
      "board": 2,
      "pon": 8,
      "onu_id": 4,
      "name": "Arif Irwan Setiawan",
      "onu_type": "F670LV7.1",
      "serial_number": "ZTEGC5A27AE1",
      "rx_power": "-19.17",
      "status": "Online"
    },
    {
      "board": 2,
      "pon": 8,
      "onu_id": 5,
      "name": "Putra Chandra Agusta",
      "onu_type": "F660V6.0",
      "serial_number": "ZTEGD00E4BCC",
      "rx_power": "-19.54",
      "status": "Online"
    },
    {
      "board": 2,
      "pon": 8,
      "onu_id": 6,
      "name": "Tarjito",
      "onu_type": "F670LV7.1",
      "serial_number": "ZTEGC5A062E0",
      "rx_power": "-21.81",
      "status": "Online"
    }
  ]
}
```

### Description of Paginate
| Syntax             | Description                                                     |
|--------------------|-----------------------------------------------------------------|
| page               | Page number                                                     |
| limit              | Limit data per page                                             |
| page_count         | Total page                                                      |
| total_rows         | Total rows                                                      |
| data               | Data of onu                                                     |

#### Default paginate
``` go
var (
	DefaultPageSize = 10 // default page size
	MaxPageSize     = 100 // max page size
	PageVar         = "page"
	PageSizeVar     = "limit"
)
```


### LICENSE
[MIT License](https://github.com/megadata-dev/go-snmp-olt-zte-c320/blob/main/LICENSE)
