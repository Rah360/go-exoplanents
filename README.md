# Exoplanet Microservice

## Description
This microservice provides functionality to add, list, update, delete exoplanets, and estimate fuel costs for trips to exoplanets.

## Features
- Add an Exoplanet
- List Exoplanets
- Get Exoplanet by ID
- Update Exoplanet
- Delete Exoplanet
- Estimate Fuel Cost for a trip to an Exoplanet

## Requirements
- Go 1.22 or later
- Docker

## Setup

### Local Development
1. **Clone the repository:**
   ```git clone https://github.com/yourusername/exoplanet_microservice.git
   cd exoplanet_microservice
Build the Go application:

```go build -o main```

Run the application:
```./main```

Docker
Build the Docker image:
```docker build -t exoplanet_microservice .```

Run the Docker container:

```docker run -p 8080:8080 exoplanet_microservice```


API Endpoints

```Add an Exoplanet
URL: /exoplanets
Method: POST
Request Body:
{
    "name": "Planet Name",
    "description": "Planet Description",
    "distance": 50,
    "radius": 1.2,
    "mass": 5.5,  // Optional for GasGiant
    "type": "Terrestrial"
}
Response:
{
    "id": "generated-uuid",
    "name": "Planet Name",
    "description": "Planet Description",
    "distance": 50,
    "radius": 1.2,
    "mass": 5.5,
    "type": "Terrestrial"
}

List Exoplanets
URL: /exoplanets
Method: GET
Response:
[
    {
        "id": "uuid",
        "name": "Planet Name",
        "description": "Planet Description",
        "distance": 50,
        "radius": 1.2,
        "mass": 5.5,
        "type": "Terrestrial"
    },
    ...
]

Get Exoplanet by ID
URL: /exoplanets/{id}
Method: GET
Response:
{
    "id": "uuid",
    "name": "Planet Name",
    "description": "Planet Description",
    "distance": 50,
    "radius": 1.2,
    "mass": 5.5,
    "type": "Terrestrial"
}

Update Exoplanet
URL: /exoplanets/{id}
Method: PATCH
Request Body: (partial update)
{
    "name": "Updated Planet Name"
}
Response:
{
    "id": "uuid",
    "name": "Updated Planet Name",
    "description": "Planet Description",
    "distance": 50,
    "radius": 1.2,
    "mass": 5.5,
    "type": "Terrestrial"
}

Delete Exoplanet
URL: /exoplanets/{id}
Method: DELETE
Response: 204 No Content
Estimate Fuel Cost
URL: /exoplanets/{id}/fuel?crew={crew}
Method: GET
Response:
{
    "fuel_cost": 123.45
}
```
Error Handling
All endpoints return appropriate HTTP status codes for errors, along with a JSON response describing the error.
