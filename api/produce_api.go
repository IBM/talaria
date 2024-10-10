package api

// import (
// 	"lightweight-broker/protocol"
// 	"lightweight-broker/protocol/models"
// 	"log/slog"
// )

// type ProduceAPI struct {
// 	Request Request
// }

// func (p ProduceAPI) Name() string {
// 	return "Produce"
// }

// func (p ProduceAPI) GetRequest() Request {
// 	return p.Request
// }

// func (p ProduceAPI) GeneratePayload(encoder *protocol.Encoder) error {
// 	req := models.ProduceRequest{}
// 	err := req.Decode(p.GetRequest().Decoder, int(p.GetRequest().Header.APIVersion))
// 	if err != nil {
// 		return err
// 	}

// 	// for _, topic := range req.TopicData {
// 	// 	for _, partition := range topic.PartitionData {
// 	// 		slog.Debug("Received records", "records", partition.Records.String())
// 	// 	}
// 	// }

// 	slog.Debug("Produce request", "req", req)

// 	return nil
// }
