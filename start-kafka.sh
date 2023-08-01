docker pull wurstmeister/kafka:latest

KAFKA_BROKER_ID="001"
KAFKA_CREATE_TOPICS="test0:1:3,test1:1:1:compact"
## test0 will have 1 partition and 3 replicas
## test1 will have 1 partition, 1 replica and a cleanup.policy set to compact.


docker run -d -p 9092:9092 \
  -e KAFKA_BROKER_ID="${KAFKA_BROKER_ID}" \
  -e KAFKA_CREATE_TOPICS="${KAFKA_CREATE_TOPICS}" \
  -e HOSTNAME_COMMAND="docker info | grep ^Name: | cut -d' ' -f 2" \
  -e KAFKA_ZOOKEEPER_CONNECT="zookeeper:2181" \
  -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP="INSIDE:PLAINTEXT,OUTSIDE:PLAINTEXT" \
  -e KAFKA_ADVERTISED_LISTENERS="INSIDE://:9091,OUTSIDE://localhost:9092" \
  -e KAFKA_LISTENERS="INSIDE://:9091,OUTSIDE://0.0.0.0:9092" \
  -e KAFKA_INTER_BROKER_LISTENER_NAME="INSIDE" \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --link zookeeper:zookeeper \
  --name kafka \
  wurstmeister/kafka:latest