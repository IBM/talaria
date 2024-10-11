// protocol has been generated from message format json - DO NOT EDIT
package protocol

// StopReplicaPartitionError contains the responses for each partition.
type StopReplicaPartitionError struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the topic name.
	TopicName string
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the partition error code, or 0 if there was no partition error.
	ErrorCode int16
}

func (p *StopReplicaPartitionError) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	if err := pe.putString(p.TopicName); err != nil {
		return err
	}

	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if p.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *StopReplicaPartitionError) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.TopicName, err = pd.getString(); err != nil {
		return err
	}

	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type StopReplicaResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the top-level error code, or 0 if there was no top-level error.
	ErrorCode int16
	// PartitionErrors contains the responses for each partition.
	PartitionErrors []StopReplicaPartitionError
}

func (r *StopReplicaResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	if err := pe.putArrayLength(len(r.PartitionErrors)); err != nil {
		return err
	}
	for _, block := range r.PartitionErrors {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *StopReplicaResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numPartitionErrors int
	if numPartitionErrors, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.PartitionErrors = make([]StopReplicaPartitionError, numPartitionErrors)
	for i := 0; i < numPartitionErrors; i++ {
		var block StopReplicaPartitionError
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.PartitionErrors[i] = block
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *StopReplicaResponse) GetKey() int16 {
	return 5
}

func (r *StopReplicaResponse) GetVersion() int16 {
	return r.Version
}

func (r *StopReplicaResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *StopReplicaResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *StopReplicaResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
