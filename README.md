# Klottr Server

## TODO

* Better error handling
* Support for comments
* Support for get a users messages

## Dependencies

* Go
* PostgreSQL with PostGIS

## Run

        go get
        go build
        DATABASE_URL="<postgres>" ./jodelkr

## Insert message

        curl --include \
             --request POST \
             --header "Content-Type: application/json" \
             --data-binary "{
                \"message\": \"Alex was here\",
                \"x\": 37.532600,
                \"y\": 127.024612
             }" \
             'http://localhost:8080/post'

## Get messages within a 10000 m area of X and Y

        curl "http://localhost:8080/get?x=37.532600&y=127.024612"

## Build and deploy the project

        GOOS=linux GOARCH=amd64 go build -o bin/klottr-linux-amd64
        cd infrastructure
        terraform taint aws_instance.web
        terraform apply

SSH in and run

        DATABASE_URL=`cat DATABASE_URL` ADDR=':80' ./klottr >klottr.log 2>&1 &

## SSH into the running server

        ssh ec2-user@`terraform output ip` -i id_rsa
