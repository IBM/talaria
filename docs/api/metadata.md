# Overview
https://cwiki.apache.org/confluence/display/KAFKA/A+Guide+To+The+Kafka+Protocol#AGuideToTheKafkaProtocol-MetadataAPI

# Caveats
* For now a rack ID cannot be assigned to the broker, since OpenTalaria will not support distributed clusters in the initial release. This will be added at later stages. To satisfy the protocol, the API will always return the rack as null.