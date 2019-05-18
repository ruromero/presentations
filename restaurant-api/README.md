# Restaurant API

## Build native image

```{bash}
mvn clean package -Pnative -Dnative-image.docker-build=true
docker build -f src/main/docker/Dockerfile.distroless -t quay.io/ruben/restaurant-api:latest .
```

## Run native image

```{bash}
docker run -p 8080:8080 -i --rm --env RESTAURANT_LOCATION="Granada" --env RESTAURANT_NAME="QuarkusDeli" --env RESTAURANT_FOOD_TYPE="Andalusian" quay.io/ruben/restaurant-api:latest
```
