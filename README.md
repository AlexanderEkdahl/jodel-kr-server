# Klottr Server

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
                \"x\": 127.024612,
                \"y\": 37.532600
             }" \
             'http://localhost:8080/post'

## Get messages within a 10000 m area of X and Y

        curl "http://localhost:8080/get?x=127.0&y=37.5"

## Insert comment

        curl --include \
             --request POST \
             --header "Content-Type: application/json" \
             --data-binary "{
                \"content\": \"Amazing post\",
                \"message_id\": 54
             }" \
             'http://localhost:8080/post_comment'

## Build and deploy the project

        GOOS=linux GOARCH=amd64 go build -o bin/klottr-linux-amd64
        cd client
        npm run build
        cd ../infrastructure
        terraform taint aws_instance.web
        terraform apply

SSH in and run the following using `screen`

        DATABASE_URL=`cat DATABASE_URL` ADDR=':8080' STATIC='./www/' ./klottr
        caddy

## SSH into the running server

        ssh ec2-user@`terraform output ip` -i id_rsa
