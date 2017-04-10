REPO="github.com/alkukampela/humanizer/dev"
CONTAINER_NAME="humanizer_build"
# Start container and build binary there
docker run -t --name $CONTAINER_NAME golang sh -c \
    "go get -v $REPO/... &&
    cd src/$REPO && 
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ."
# Get binary to host machine
docker cp $CONTAINER_NAME:go/src/$REPO/main .
# Remove build container
docker rm  $CONTAINER_NAME
# Build binary container
docker build --rm -t humanizer .
