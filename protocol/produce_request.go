// protocol has been generated from message format json - DO NOT EDIT
package protocol

// PartitionProduceData contains each partition to produce to.
type PartitionProduceData struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Index contains the partition index.
	Index int32
	// Records contains the record data to be produced.
	Records RecordBatch
}

func (p *PartitionProduceData) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.Index)

	if err := p.Records.encode(pe, p.Version); err != nil {
		return err
	}

	if p.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *PartitionProduceData) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Index, err = pd.getInt32(); err != nil {
		return err
	}

	tmprecords := RecordBatch{}
	if err := tmprecords.decode(pd, p.Version); err != nil {
		return err
	}
	p.Records = tmprecords

	if p.Version >= 9 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// TopicProduceData contains each topic to produce to.
type TopicProduceData struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// PartitionData contains each partition to produce to.
	PartitionData []PartitionProduceData
}

func (t *TopicProduceData) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(t.PartitionData)); err != nil {
		return err
	}
	for _, block := range t.PartitionData {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	if t.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *TopicProduceData) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitionData int
	if numPartitionData, err = pd.getArrayLength(); err != nil {
		return err
	}
	t.PartitionData = make([]PartitionProduceData, numPartitionData)
	for i := 0; i < numPartitionData; i++ {
		var block PartitionProduceData
		if err := block.decode(pd, t.Version); err != nil {
			return err
		}
		t.PartitionData[i] = block
	}

	if t.Version >= 9 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type ProduceRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TransactionalID contains the transactional ID, or null if the producer is not transactional.
	TransactionalID *string
	// Acks contains the number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments, 1 for only the leader and -1 for the full ISR.
	Acks int16
	// TimeoutMs contains the timeout to await a response in milliseconds.
	TimeoutMs int32
	// TopicData contains each topic to produce to.
	TopicData []TopicProduceData
}

func (r *ProduceRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 9 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 3 {
		if err := pe.putNullableString(r.TransactionalID); err != nil {
			return err
		}
	}

	pe.putInt16(r.Acks)

	pe.putInt32(r.TimeoutMs)

	if err := pe.putArrayLength(len(r.TopicData)); err != nil {
		return err
	}
	for _, block := range r.TopicData {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 9 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ProduceRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 9 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 3 {
		if r.TransactionalID, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Acks, err = pd.getInt16(); err != nil {
		return err
	}

	if r.TimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numTopicData int
	if numTopicData, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.TopicData = make([]TopicProduceData, numTopicData)
	for i := 0; i < numTopicData; i++ {
		var block TopicProduceData
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.TopicData[i] = block
	}

	if r.Version >= 9 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ProduceRequest) GetKey() int16 {
	return 0
}

func (r *ProduceRequest) GetVersion() int16 {
	return r.Version
}

func (r *ProduceRequest) GetHeaderVersion() int16 {
	if r.Version >= 9 {
		return 2
	}
	return 1
}

func (r *ProduceRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 9
}

func (r *ProduceRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
