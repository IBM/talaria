// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// DescribeClusterBroker contains each broker in the response.
type DescribeClusterBroker struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the broker ID.
	BrokerID int32
	// Host contains the broker hostname.
	Host string
	// Port contains the broker port.
	Port int32
	// Rack contains the rack of the broker, or null if it has not been assigned to a rack.
	Rack *string
}

func (b *DescribeClusterBroker) encode(pe packetEncoder, version int16) (err error) {
	b.Version = version
	pe.putInt32(b.BrokerID)

	if err := pe.putString(b.Host); err != nil {
		return err
	}

	pe.putInt32(b.Port)

	if err := pe.putNullableString(b.Rack); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (b *DescribeClusterBroker) decode(pd packetDecoder, version int16) (err error) {
	b.Version = version
	if b.BrokerID, err = pd.getInt32(); err != nil {
		return err
	}

	if b.Host, err = pd.getString(); err != nil {
		return err
	}

	if b.Port, err = pd.getInt32(); err != nil {
		return err
	}

	if b.Rack, err = pd.getNullableString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeClusterResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the top-level error code, or 0 if there was no error
	ErrorCode int16
	// ErrorMessage contains the top-level error message, or null if there was no error.
	ErrorMessage *string
	// ClusterID contains the cluster ID that responding broker belongs to.
	ClusterID string
	// ControllerID contains the ID of the controller broker.
	ControllerID int32
	// Brokers contains each broker in the response.
	Brokers []DescribeClusterBroker
	// ClusterAuthorizedOperations contains a 32-bit bitfield to represent authorized operations for this cluster.
	ClusterAuthorizedOperations int32
}

func (r *DescribeClusterResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putString(r.ClusterID); err != nil {
		return err
	}

	pe.putInt32(r.ControllerID)

	if err := pe.putArrayLength(len(r.Brokers)); err != nil {
		return err
	}
	for _, block := range r.Brokers {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putInt32(r.ClusterAuthorizedOperations)

	pe.putUVarint(0)
	return nil
}

func (r *DescribeClusterResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.ClusterID, err = pd.getString(); err != nil {
		return err
	}

	if r.ControllerID, err = pd.getInt32(); err != nil {
		return err
	}

	var numBrokers int
	if numBrokers, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Brokers = make([]DescribeClusterBroker, numBrokers)
	for i := 0; i < numBrokers; i++ {
		var block DescribeClusterBroker
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Brokers[i] = block
	}

	if r.ClusterAuthorizedOperations, err = pd.getInt32(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeClusterResponse) GetKey() int16 {
	return 60
}

func (r *DescribeClusterResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeClusterResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *DescribeClusterResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeClusterResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeClusterResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
