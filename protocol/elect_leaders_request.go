// protocol has been generated from message format json - DO NOT EDIT
package protocol

// TopicPartitions_ElectLeadersRequest contains the topic partitions to elect leaders.
type TopicPartitions_ElectLeadersRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains the name of a topic.
	Topic string
	// Partitions contains the partitions of this topic whose leader should be elected.
	Partitions []int32
}

func (t *TopicPartitions_ElectLeadersRequest) encode(pe packetEncoder, version int16) (err error) {
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

func (t *TopicPartitions_ElectLeadersRequest) decode(pd packetDecoder, version int16) (err error) {
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

type ElectLeadersRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ElectionType contains a Type of elections to conduct for the partition. A value of '0' elects the preferred replica. A value of '1' elects the first live replica if there are no in-sync replica.
	ElectionType int8
	// TopicPartitions contains the topic partitions to elect leaders.
	TopicPartitions []TopicPartitions_ElectLeadersRequest
	// TimeoutMs contains the time in ms to wait for the election to complete.
	TimeoutMs int32
}

func (r *ElectLeadersRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 1 {
		pe.putInt8(r.ElectionType)
	}

	if err := pe.putArrayLength(len(r.TopicPartitions)); err != nil {
		return err
	}
	for _, block := range r.TopicPartitions {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putInt32(r.TimeoutMs)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ElectLeadersRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 1 {
		if r.ElectionType, err = pd.getInt8(); err != nil {
			return err
		}
	}

	var numTopicPartitions int
	if numTopicPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopicPartitions > 0 {
		r.TopicPartitions = make([]TopicPartitions_ElectLeadersRequest, numTopicPartitions)
		for i := 0; i < numTopicPartitions; i++ {
			var block TopicPartitions_ElectLeadersRequest
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.TopicPartitions[i] = block
		}
	}

	if r.TimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ElectLeadersRequest) GetKey() int16 {
	return 43
}

func (r *ElectLeadersRequest) GetVersion() int16 {
	return r.Version
}

func (r *ElectLeadersRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *ElectLeadersRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *ElectLeadersRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
