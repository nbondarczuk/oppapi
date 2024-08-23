
#!/bin/bash

REPO="quay.io"
USER=$(id -u)
GROUP=$(id -g)
cmd="docker run --rm -it --user $USER:$GROUP --mount type=bind,src=$(pwd),dst=/opt/work --workdir /opt/work -e GOCACHE=/tmp -v $HOME:$HOME -p 8090:8090 $REPO/goswagger/swagger"
echo $cmd $1 $2 $3 $4 $5 $6 $7
$cmd $1 $2 $3 $4 $5 $6 $7
