# URL-shortener
Ozon fintech internship task.

URL-shortener is a project that aims to create a stateless service for creating shortened 
links and getting their original reference.

## Table of Contents
- [Architecture](#architecture)
- [Database overview](#database-overview)
- [Deployment](#deployment)
- [Documentation](#documentation)
- [Demo](#demo)

## Architecture
The project can be divided into 3 main components:
- **gRPC server** - main component of the service. It is a web application that 
provides a gRPC API for shortening links and getting their original references.
- **Database** - there are two options: in-memory storage and on disk. In-memory storage is implemented using Redis 
when as disk storage is implemented using PostgreSQL
- **HTTP proxy server** - gRPC Gateway implementation. It is provides a REST API for interacting with the gRPC server

![Architecture of the system](/assets/architecture.png)

## Database overview
### PostgreSQL
In PostreSQL database there is only one table - *urls*.

![PostgreSQL database diagram](/assets/pg_database.png)

### Redis
In Redis, for persistence and uniqueness service is using two Redis databases.

![Redis database diagram](/assets/redis_db_diagram.jpeg)

## Deployment
All components are dockerized and can be deployed using make commands.

**First deploy with PostgreSQL database**
```
make compile postgres
```

**First deploy with Redis database**
```
make compile redis
```

>In case you want to change database - stop running container and run the following command:
>
>**For Redis**
>```
>make redis
>```
>**For PostgreSQL**
>```
>make postgres
>```

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
shortUrl - is a given later shortened URL

**For posting original URL and getting a shor one:**

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

## Demo
### Only gRPC

**PostUrl procedure**

![PostUrl gRPC](/assets/post_url_grpc.png)

**GetUrl procedure**

![GetUrl gRPC](/assets/get_url_grpc.png)

### With gRPC Gateway

**PostUrl procedure**

![PostUrl HTTP](/assets/post_url_http.png)

**GetUrl procedure**

![GetUrl HTTP](/assets/get_url_http.png)
