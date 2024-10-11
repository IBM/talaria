// protocol has been generated from message format json - DO NOT EDIT
package protocol

// RemainingPartition contains the partitions that the broker still leads.
type RemainingPartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TopicName contains the name of the topic.
	TopicName string
	// PartitionIndex contains the index of the partition.
	PartitionIndex int32
}

func (r *RemainingPartition) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.TopicName); err != nil {
		return err
	}

	pe.putInt32(r.PartitionIndex)

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *RemainingPartition) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.TopicName, err = pd.getString(); err != nil {
		return err
	}

	if r.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type ControlledShutdownResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the top-level error code.
	ErrorCode int16
	// RemainingPartitions contains the partitions that the broker still leads.
	RemainingPartitions []RemainingPartition
}

func (r *ControlledShutdownResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	if err := pe.putArrayLength(len(r.RemainingPartitions)); err != nil {
		return err
	}
	for _, block := range r.RemainingPartitions {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ControlledShutdownResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	var numRemainingPartitions int
	if numRemainingPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.RemainingPartitions = make([]RemainingPartition, numRemainingPartitions)
	for i := 0; i < numRemainingPartitions; i++ {
		var block RemainingPartition
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.RemainingPartitions[i] = block
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ControlledShutdownResponse) GetKey() int16 {
	return 7
}

func (r *ControlledShutdownResponse) GetVersion() int16 {
	return r.Version
}

func (r *ControlledShutdownResponse) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 1
	}
	return 0
}

func (r *ControlledShutdownResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *ControlledShutdownResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
