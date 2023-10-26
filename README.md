# Cakeplabs Exam

## Prerequisite

- [Go 1.21.3](https://golang.org/dl/)
- [GORM](https://gorm.io/index.html)
- [Echo Framework](https://echo.labstack.com/)

## Running this Repository Locally

1. Clone this repository

   ```shell
    git clone https://github.com/titan2903/cakeplabs-technical-test.git
   ```

2. Change to the `cakeplabs-technical-test` directory

    ```shell
    cd cakeplabs-technical-test
    ```

3. Run command `go mod tidy` to clean up and optimize the dependencies of a Go module

    ```shell
    go mod tidy
    ```

4. Copy the file `.env.example` to a file named `.env`

   ```shell
   cp .env.example .env
   ```

5. Configure environment variables to match your local database

6. Run this repository locally

   ```shell
   go run main.go
   ```

7. This repository will be running on port `:8080`

8. Check the healthcheck endpoint at `http://localhost:8080/v1/healthcheck` to ensure this application is running properly

9. Sedeer data Menus and Additions with this endpoint

    ```shell
    curl --location --request POST 'http://localhost:8080/v1/seeder'
    ```

## Postman API Sample

- Can view the Postman documentation in the file `postman_collection.json` within this repository or in this [link postman collection](https://ik.imagekit.io/ckb21lc9cd/postman_collection_18WzZQ4tF.json?updatedAt=1698279382114)

## Dockerhub Pull Image

- [Dockerhub pull image](https://hub.docker.com/r/titan29/cakeplabs-technical-test)

## Relational Database

- [Database modeling](https://dbdiagram.io/d/cakeplabs-tehcnical-test-65386b2affbf5169f062112e)

## Github Action

- [Github action pipeline](https://github.com/titan2903/cakeplabs-technical-test/actions)
