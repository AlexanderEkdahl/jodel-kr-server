# JodelKR Server

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
