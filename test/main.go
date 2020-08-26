package main

import (
	"fmt"
	"log"
	"os/exec"
)

const shellCMD = `
container_name='mysql-dev'
container_id=$(docker ps -aq -f name=${container_name})
docker rm -f "${container_id}"
docker run -d -p 3306:3306 --privileged=true \
    -v /home/docker/mysql/conf/my.conf:/etc/mysql/my.conf \
    -v /home/docker/mysql/data:/var/lib/mysql \
    -e MYSQL_USER="root" \
    -e MYSQL_PASSWORD="test123" \
    -e MYSQL_ROOT_PASSWORD="test123" \
    --name ${container_name} \
    mysql:5.7
docker ps -a -f name=${container_name}
`

func main() {
	cmd := exec.Command("/bin/bash", "-c", shellCMD)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("command exec err: %v\n", err)
	}
	fmt.Printf("%s\n", output)
}
