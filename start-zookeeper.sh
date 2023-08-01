docker pull wurstmeister/zookeeper:latest
docker run -d \
  -p 2181:2181 \
  --name zookeeper \
  wurstmeister/zookeeper:latest