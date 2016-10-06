#

## Install Go and PostgreSQL with PostGIS        
        
        go get
        go build
        DATABASE_URL="<postgres>" ./jodelkr

## Insert message

        curl --data "message=Alex was here&x=37.532600y=127.024612" http://localhost:8080/post

## Get messages within a 1000000 m area of X and Y

        curl http://localhost:8080/get?x=37.532600&y=127.024612
