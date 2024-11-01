package protocol

import "fmt"

func (r *RequestHeader) String() string {
	return fmt.Sprintf("API Key: %d\nAPI Version: %d\nCorrelation ID: %d\nClient Name: %s",
		r.RequestApiKey,
		r.RequestApiVersion,
		r.CorrelationID,
		*r.ClientID)
}
