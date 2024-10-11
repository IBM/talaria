// protocol has been generated from message format json - DO NOT EDIT
package protocol

// TopicRequest contains a
type TopicRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// PartitionIndexes contains the indexes of the partitions to list producers for.
	PartitionIndexes []int32
}

func (t *TopicRequest) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.PartitionIndexes); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *TopicRequest) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	if t.PartitionIndexes, err = pd.getInt32Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeProducersRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topics contains a
	Topics []TopicRequest
}

func (r *DescribeProducersRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeProducersRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]TopicRequest, numTopics)
	for i := 0; i < numTopics; i++ {
		var block TopicRequest
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeProducersRequest) GetKey() int16 {
	return 61
}

func (r *DescribeProducersRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeProducersRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *DescribeProducersRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeProducersRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
