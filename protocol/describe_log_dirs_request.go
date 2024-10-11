// protocol has been generated from message format json - DO NOT EDIT
package protocol

// DescribableLogDirTopic contains each topic that we want to describe log directories for, or null for all topics.
type DescribableLogDirTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains the topic name
	Topic string
	// Partitions contains the partition indexes.
	Partitions []int32
}

func (t *DescribableLogDirTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Topic); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.Partitions); err != nil {
		return err
	}

	if t.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *DescribableLogDirTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Topic, err = pd.getString(); err != nil {
		return err
	}

	if t.Partitions, err = pd.getInt32Array(); err != nil {
		return err
	}

	if t.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogDirsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topics contains each topic that we want to describe log directories for, or null for all topics.
	Topics []DescribableLogDirTopic
}

func (r *DescribeLogDirsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Topics)); err != nil {
		return err
	}
	for _, block := range r.Topics {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeLogDirsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]DescribableLogDirTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block DescribableLogDirTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeLogDirsRequest) GetKey() int16 {
	return 35
}

func (r *DescribeLogDirsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeLogDirsRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *DescribeLogDirsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *DescribeLogDirsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
