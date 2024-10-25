// protocol has been generated from message format json - DO NOT EDIT
package protocol

// WritableTxnMarkerPartitionResult contains the results by partition.
type WritableTxnMarkerPartitionResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition index.
	PartitionIndex int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
}

func (p *WritableTxnMarkerPartitionResult) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	pe.putInt32(p.PartitionIndex)

	pe.putInt16(p.ErrorCode)

	if p.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *WritableTxnMarkerPartitionResult) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if p.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if p.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// WritableTxnMarkerTopicResult contains the results by topic.
type WritableTxnMarkerTopicResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains the results by partition.
	Partitions []WritableTxnMarkerPartitionResult
}

func (t *WritableTxnMarkerTopicResult) encode(pe packetEncoder, version int16) (err error) {
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

	if t.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *WritableTxnMarkerTopicResult) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	var numPartitions int
	if numPartitions, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numPartitions > 0 {
		t.Partitions = make([]WritableTxnMarkerPartitionResult, numPartitions)
		for i := 0; i < numPartitions; i++ {
			var block WritableTxnMarkerPartitionResult
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Partitions[i] = block
		}
	}

	if t.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// WritableTxnMarkerResult contains the results for writing makers.
type WritableTxnMarkerResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ProducerID contains the current producer ID in use by the transactional ID.
	ProducerID int64
	// Topics contains the results by topic.
	Topics []WritableTxnMarkerTopicResult
}

func (m *WritableTxnMarkerResult) encode(pe packetEncoder, version int16) (err error) {
	m.Version = version
	pe.putInt64(m.ProducerID)

	if err := pe.putArrayLength(len(m.Topics)); err != nil {
		return err
	}
	for _, block := range m.Topics {
		if err := block.encode(pe, m.Version); err != nil {
			return err
		}
	}

	if m.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (m *WritableTxnMarkerResult) decode(pd packetDecoder, version int16) (err error) {
	m.Version = version
	if m.ProducerID, err = pd.getInt64(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		m.Topics = make([]WritableTxnMarkerTopicResult, numTopics)
		for i := 0; i < numTopics; i++ {
			var block WritableTxnMarkerTopicResult
			if err := block.decode(pd, m.Version); err != nil {
				return err
			}
			m.Topics[i] = block
		}
	}

	if m.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type WriteTxnMarkersResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Markers contains the results for writing makers.
	Markers []WritableTxnMarkerResult
}

func (r *WriteTxnMarkersResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Markers)); err != nil {
		return err
	}
	for _, block := range r.Markers {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *WriteTxnMarkersResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numMarkers int
	if numMarkers, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numMarkers > 0 {
		r.Markers = make([]WritableTxnMarkerResult, numMarkers)
		for i := 0; i < numMarkers; i++ {
			var block WritableTxnMarkerResult
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Markers[i] = block
		}
	}

	if r.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *WriteTxnMarkersResponse) GetKey() int16 {
	return 27
}

func (r *WriteTxnMarkersResponse) GetVersion() int16 {
	return r.Version
}

func (r *WriteTxnMarkersResponse) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 1
	}
	return 0
}

func (r *WriteTxnMarkersResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *WriteTxnMarkersResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
