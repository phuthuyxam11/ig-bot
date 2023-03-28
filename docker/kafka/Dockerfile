FROM wurstmeister/kafka:2.12-2.5.0

ENV KAFKA_CREATE_TOPICS = "sendmail_topic:3:1"
ENV KAFKA_ADVERTISED_HOST_NAME = "kafka"
ENV KAFKA_ADVERTISED_PORT = "9092"
ENV KAFKA_ZOOKEEPER_CONNECT = "zookeeper:2181"
ENV KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR = 1
ENV KAFKA_MESSAGE_MAX_BYTES = 104858800

RUN kafka-topics.sh --bootstrap-server localhost:9092 --topic sendmail_topic --create --partitions 3 --replication-factor 1


RUN kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic sendmail_topic --group sendmail_consumer_group
