# Reviewer

## How to start
if you run this program for the first time, you need to seed the MongoDB database.

First, start the Docker compose stack with

``` shell
docker compose up --build
```

By default, the backend program `app` will run at port `8000`, the MongoDB instance at port `27017`.

Now that the docker compose stack is up and running, call the Makefile target `seed`, like this

``` shell
MONGODB_URL=mongodb://localhost:27017/app make seed
```

and you are ready to go :)

## How to use
Use any RESTFul client of your choice. These are the available API endpoints:

- `GET /product` lists all the products, without parameters `page` and `amount` the backend will return the first page of the product list with 10 items.
- `POST /product/{id}/review` will add a new review, this only accepts a JSON body:

``` json
{
    "name": "User name",
    "text": "my experience with this product...",
    "rating": 5
}
```

the `"rating"` property is anything between `> 0 and <= 5`.

## How to run unit tests
You do that with `make test`
