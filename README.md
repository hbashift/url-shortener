# URL-shortener
Ozon fintech internship task.

URL-shortener is a service for creating shortened links and getting their original reference.

## Table of Contents
- [Architecture](#architecture)
- [URL encoding algorithm](#url-encoding-algorithm)
- [Deployment](#deployment)
- [Documentation](#documentation)

## Architecture
The project can be divided into 3 main components:
- **gRPC server** - main component of the service. It is a web application that 
provides a gRPC API for shortening links and getting their original references.
- **Database** - there are two options: in-memory storage and on disk. In-memory storage is implemented using Redis 
when as disk storage is implemented using PostgreSQL
- **HTTP proxy server** - gRPC Gateway implementation. It is provides a REST API for interacting with the gRPC server

![Architecture of the system](/assets/architecture.png)

## URL encoding algorithm
The basic principle is as follows - the transition from a number system based on 10 to a number system based on 63. 
Collisions are handled by shuffling the base alphabet in random order. The number that is encoded is a randomly 
generated value.

## Deployment
All components are dockerized and can be deployed using docker compose.

**Deploy with PostgreSQL database**
```
docker compose -f docker-compose-postgres.yaml up
```

**Deploy with Redis database**
```
docker compose -f docker-compose-redis.yaml up
```

## Documentation
### gRPC only 
You can see a protobuf file right [here](proto/shortener.proto).

### With gRPC Gateway
Here service has two endpoints:

**For getting original URL:**

Endpoint
```
/v1/url/{shortUrl}
```
shortUrl - is a given earlier shortened URL

**For posting original URL and getting a short one:**

Endpoint:
```
/v1/url
```
Request body:
```json
{
  "longUrl": "http://localhost:8080/test_url"
}
```

Also, service provides a swagger documentation. You can download it [here](assets/index.html).
