package api

import (
	"opentalaria/protocol"
	"opentalaria/utils"
	"time"
)

type MetadataAPI struct {
	Request Request
}

func (m MetadataAPI) Name() string {
	return "Metadata"
}

func (m MetadataAPI) GetRequest() Request {
	return m.Request
}

func (m MetadataAPI) GetHeaderVersion(requestVersion int16) int16 {
	return (&protocol.MetadataResponse{Version: requestVersion}).GetHeaderVersion()
}

func (m MetadataAPI) GeneratePayload() ([]byte, error) {
	req := protocol.MetadataRequest{}
	_, err := protocol.VersionedDecode(m.GetRequest().Message, &req, m.GetRequest().Header.RequestApiVersion)
	if err != nil {
		return nil, err
	}

	response := GenerateMetadataResponse(m.GetRequest().Header.RequestApiVersion)
	return protocol.Encode(response)
}

func GenerateMetadataResponse(version int16) *protocol.MetadataResponse {
	// For now the returned data is mock, just so we can continue developing the rest of the APIs.
	// Once we have a more robust project architecture, this struct will be populated with the real
	// cluster metadata.
	response := protocol.MetadataResponse{}

	response.Version = version
	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	response.Brokers = append(response.Brokers, protocol.MetadataResponseBroker{
		NodeID: 1,
		Host:   "localhost",
		Port:   9092,
		Rack:   nil, // for now OpenTalaria does not support rack awareness.
	})

	clusterId := "aaabbbfff"
	response.ClusterID = &clusterId
	response.ControllerID = 1
	topicName := "test-topic"

	response.Topics = append(response.Topics, protocol.MetadataResponseTopic{
		ErrorCode:  int16(utils.ErrNoError),
		Name:       &topicName,
		IsInternal: false,
		Partitions: []protocol.MetadataResponsePartition{{
			ErrorCode:       int16(utils.ErrNoError),
			PartitionIndex:  0,
			LeaderID:        1,
			LeaderEpoch:     int32(time.Now().Unix()),
			ReplicaNodes:    []int32{0},
			IsrNodes:        []int32{0},
			OfflineReplicas: []int32{0},
		}},
		TopicAuthorizedOperations: 0,
	})
	response.ClusterAuthorizedOperations = 0

	return &response
}
