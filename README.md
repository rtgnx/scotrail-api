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
$: curl -X GET https://scotrail.pw/station/EXG

```

Example response:

```JSON
{
  "code": "EXG",
  "latitude": 55.860619,
  "longitude": -4.283207,
  "name": "Exhibition Centre (Glasgow)"
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
    "code": "EXG",
    "latitude": 55.860619,
    "longitude": -4.283207,
    "name": "Exhibition Centre (Glasgow)"
  },
  {
    "code": "QPK",
    "latitude": 55.835747,
    "longitude": -4.267368,
    "name": "Queens Park (Glasgow)"
  },
  {
    "code": "GGT",
    "latitude": 55.865803,
    "longitude": -4.432114,
    "name": "Glasgow Airport"
  },
  {
    "code": "PTG",
    "latitude": 55.9342,
    "longitude": -4.689432,
    "name": "Port Glasgow"
  },
  {
    "code": "CHC",
    "latitude": 55.865345,
    "longitude": -4.270689,
    "name": "Charing Cross (Glasgow)"
  },
  {
    "code": "HST",
    "latitude": 55.858719,
    "longitude": -4.239953,
    "name": "High Street (Glasgow)"
  },
  {
    "code": "GLC",
    "latitude": 55.860157,
    "longitude": -4.259208,
    "name": "Glasgow Central"
  },
  {
    "code": "GLQ",
    "latitude": 55.863003,
    "longitude": -4.251385,
    "name": "Glasgow Queen Street"
  },
  {
    "code": "81",
    "latitude": 0,
    "longitude": 0,
    "name": "Glasgow (Any)"
  }
]
```

### Nearest station

Endpoint: `GET https://scotrail.pw/nearest/:lat:/:lon:`

Example request:
```
$: curl -X GET https://scotrail.pw/nearest/55.860150/-4.259308

```

Example response:

```JSON

{
    "distance": 6, // Meters
    "station": {
        "code": "GLC",
        "latitude": 55.860157,
        "longitude": -4.259208,
        "name": "Glasgow Central"
    }
}
```

### All stations

Endpoint: `GET https://scotrail.pw/stations`

Example request:
```
$: curl -X GET https://scotrail.pw/stations

```

Example response:

```JSON
[
  {
    "code": "EXG",
    "latitude": 55.860619,
    "longitude": -4.283207,
    "name": "Exhibition Centre (Glasgow)"
  },
  {
    "code": "QPK",
    "latitude": 55.835747,
    "longitude": -4.267368,
    "name": "Queens Park (Glasgow)"
  },

  ...
]
```
