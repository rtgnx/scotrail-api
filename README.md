# Scotrail API

[![Build Status](https://travis-ci.com/rtgnx/scotrail-api.svg?token=FDwqTidyvdfhSyx6QupG&branch=master)](https://travis-ci.com/rtgnx/scotrail-api)

## How to build ?

Simplest way to build this service is to use provided `Makfile` which will build
docker image and run it.

`$: go test`     - Run tests
`$: make build`  - Build Docker image
`$: make run`    - Build docker image and start container
`$: make update` - Build new image and restart docker container


## API Documentation

### Live board

Endpoint: `GET https://scotrail.pw/live/:station_code:`

Example request:
```
$: curl -X GET https://scotrail.pw/live/GLC

```

Example response:

```JSON
{
  "Station": {
  "name": "Glasgow Central",
  "postcode": "G1 3SA",
  "code": "GLC"
  },
  "Services": [
    {
      "id": "aQjBGRoe%2BIXqfw0PxLyCgg%3D%3D",
      "platform": 0,
      "destination": "Milngavie",
      "departs": "2018-05-27T12:04:00Z",
      "arrives": "2018-05-27T12:00:00Z",
      "expected": "Cancelled",
      "origin": "Rutherglen",
      "operator": "ScotRail"
    },
    {
      "id": "hBBH%23C%2BZoU3l%23R68lwnWTA%3D%3D",
      "platform": 0,
      "destination": "Motherwell",
      "departs": "2018-05-27T12:08:00Z",
      "arrives": "0001-01-01T00:00:00Z",
      "expected": "On time",
      "origin": "Glasgow Central",
      "operator": "ScotRail"
    }
  ]
}
```



### Service details

Endpoint: `GET https://scotrail.pw/service/:service_id:`

Example request:
```
$: curl -X GET https://scotrail.pw/service/lq3VvvsPFW1qAqJQ5qeb9Q%3D%3D

```

Example response:

```JSON
[
  {
    "Station": "East Kilbride",
    "Status": "Departed on time",
    "Connection": null,
    "Time": "2018-05-27T11:26:00Z"
  },
  {
    "Station": "Hairmyres",
    "Status": "Departed no report",
    "Connection": null,
    "Time": "2018-05-27T11:30:00Z"
  }
	...
]
```

### Station details

Endpoint: `GET https://scotrail.pw/station/:station_code:`

Example request:
```
$: curl -X GET https://scotrail.pw/station/GLC

```

Example response:

```JSON
{
  "name": "Glasgow Central",
  "postcode": "G1 3SA",
  "code": "GLC"
}
```

### Station search

Endpoint: `GET https://scotrail.pw/search/:name:`

Example request:
```
$: curl -X GET https://scotrail.pw/search/Glasgow

```

Example response:

```JSON
[
  {
    "name": "Queen's Park (Glasgow)",
    "postcode": "G42 8PH",
    "code": "QPK"
  },
  {
    "name": "Exhibition Centre (Glasgow)",
    "postcode": "G3 8LE",
    "code": "EXG"
  },
  {
    "name": "Port Glasgow",
    "postcode": "PA14 5JN",
    "code": "PTG"
  },
  {
    "name": "Charing Cross (Glasgow)",
    "postcode": "G2 4PR",
    "code": "CHC"
  },
  {
    "name": "Glasgow Central",
    "postcode": "G1 3SA",
    "code": "GLC"
  },
  {
    "name": "High Street Glasgow",
    "postcode": "G1 1QF",
    "code": "HST"
  },
  {
    "name": "Glasgow Queen Street",
    "postcode": "G1 2AF",
    "code": "GLQ"
  }
]
```
