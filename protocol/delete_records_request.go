// protocol has been generated from message format json - DO NOT EDIT
package protocol

// DeleteRecordsPartition contains each partition that we want to delete records from.
type DeleteRecordsPartition struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// Offset contains the deletion offset.
	Offset int64
}

func (p *DeleteRecordsPartition) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt64(p.Offset)

	if p.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *DeleteRecordsPartition) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.Offset, err = pd.getInt64(); err != nil {
		return err
	}

	if p.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// DeleteRecordsTopic contains each topic that we want to delete records from.
type DeleteRecordsTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains each partition that we want to delete records from.
	Partitions []DeleteRecordsPartition
}

func (t *DeleteRecordsTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(t.Partitions)); err != nil {
		return err
	}
	for _, block := range t.Partitions {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	if t.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *DeleteRecordsTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.Partitions = make([]DeleteRecordsPartition, numPartitions)
	for i := 0; i < numPartitions; i++ {
		var block DeleteRecordsPartition
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.Partitions[i] = block
	}

	if t.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRecordsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topics contains each topic that we want to delete records from.
	Topics []DeleteRecordsTopic
	// TimeoutMs contains a How long to wait for the deletion to complete, in milliseconds.
	TimeoutMs int32
}

func (r *DeleteRecordsRequest) encode(pe packetEncoder) (err error) {
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

	pe.putInt32(r.TimeoutMs)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeleteRecordsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Topics = make([]DeleteRecordsTopic, numTopics)
	for i := 0; i < numTopics; i++ {
		var block DeleteRecordsTopic
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Topics[i] = block
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

func (r *DeleteRecordsRequest) GetKey() int16 {
	return 21
}

func (r *DeleteRecordsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DeleteRecordsRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *DeleteRecordsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *DeleteRecordsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
