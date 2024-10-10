package api

import (
	"log/slog"
	"talaria/protocol"
	"talaria/utils"
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
	err := protocol.VersionedDecode(m.Request.Message, &req, m.GetRequest().Header.RequestApiVersion)
	if err != nil {
		return nil, err
	}

	slog.Debug("Metadata request", "req", req)

	response := GenerateMetadataResponse()
	return protocol.Encode(response)
}

func GenerateMetadataResponse() *protocol.MetadataResponse {
	response := protocol.MetadataResponse{}

	// TODO: handle throttle time
	response.ThrottleTimeMs = 0
	rack := ""
	response.Brokers = append(response.Brokers, protocol.MetadataResponseBroker{
		NodeID: 1,
		Host:   "localhost",
		Port:   9092,
		Rack:   &rack,
	})

	clusterId := ""
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
