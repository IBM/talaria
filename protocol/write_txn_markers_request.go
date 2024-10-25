// protocol has been generated from message format json - DO NOT EDIT
package protocol

// WritableTxnMarkerTopic contains each topic that we want to write transaction marker(s) for.
type WritableTxnMarkerTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// PartitionIndexes contains the indexes of the partitions to write transaction markers for.
	PartitionIndexes []int32
}

func (t *WritableTxnMarkerTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.PartitionIndexes); err != nil {
		return err
	}

	if t.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *WritableTxnMarkerTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	if t.PartitionIndexes, err = pd.getInt32Array(); err != nil {
		return err
	}

	if t.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// WritableTxnMarker contains the transaction markers to be written.
type WritableTxnMarker struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ProducerID contains the current producer ID.
	ProducerID int64
	// ProducerEpoch contains the current epoch associated with the producer ID.
	ProducerEpoch int16
	// TransactionResult contains the result of the transaction to write to the partitions (false = ABORT, true = COMMIT).
	TransactionResult bool
	// Topics contains each topic that we want to write transaction marker(s) for.
	Topics []WritableTxnMarkerTopic
	// CoordinatorEpoch contains a Epoch associated with the transaction state partition hosted by this transaction coordinator
	CoordinatorEpoch int32
}

func (m *WritableTxnMarker) encode(pe packetEncoder, version int16) (err error) {
	m.Version = version
	pe.putInt64(m.ProducerID)

	pe.putInt16(m.ProducerEpoch)

	pe.putBool(m.TransactionResult)

	if err := pe.putArrayLength(len(m.Topics)); err != nil {
		return err
	}
	for _, block := range m.Topics {
		if err := block.encode(pe, m.Version); err != nil {
			return err
		}
	}

	pe.putInt32(m.CoordinatorEpoch)

	if m.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (m *WritableTxnMarker) decode(pd packetDecoder, version int16) (err error) {
	m.Version = version
	if m.ProducerID, err = pd.getInt64(); err != nil {
		return err
	}

	if m.ProducerEpoch, err = pd.getInt16(); err != nil {
		return err
	}

	if m.TransactionResult, err = pd.getBool(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		m.Topics = make([]WritableTxnMarkerTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block WritableTxnMarkerTopic
			if err := block.decode(pd, m.Version); err != nil {
				return err
			}
			m.Topics[i] = block
		}
	}

	if m.CoordinatorEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if m.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type WriteTxnMarkersRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Markers contains the transaction markers to be written.
	Markers []WritableTxnMarker
}

func (r *WriteTxnMarkersRequest) encode(pe packetEncoder) (err error) {
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

func (r *WriteTxnMarkersRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numMarkers int
	if numMarkers, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numMarkers > 0 {
		r.Markers = make([]WritableTxnMarker, numMarkers)
		for i := 0; i < numMarkers; i++ {
			var block WritableTxnMarker
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

func (r *WriteTxnMarkersRequest) GetKey() int16 {
	return 27
}

func (r *WriteTxnMarkersRequest) GetVersion() int16 {
	return r.Version
}

func (r *WriteTxnMarkersRequest) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 2
	}
	return 1
}

func (r *WriteTxnMarkersRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *WriteTxnMarkersRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
