package protocol

import (
	"time"
)

type CompressionType int8

const (
	CompressionNone CompressionType = iota
	CompressionGzip
	CompressionSnappy
	CompressionLz4
	CompressionZstd
)

type TimestampType int8

const (
	CreateTime TimestampType = iota
	LogAppendTime

	compressionCodecBit   = 0x07 // bits 0-2 in the attributes field signify compression
	timestampTypeBit      = 0x08 // attr bit 3
	isTransactionalBit    = 0x10 // attr bit 4
	isControlBit          = 0x20 // attr bit 5
	hasDeleteHorizonMsBit = 0x40 // attr bit 6
)

// RecordBatch is the struct representation of the Kafka primitive type records
// https://kafka.apache.org/documentation/#recordbatch
type RecordBatch struct {
	size       int32
	attributes int16

	Version              int16
	BaseOffset           int64
	BatchLength          int32
	PartitionLeaderEpoch int32
	Magic                int8
	CRC                  uint32
	CompressionType      CompressionType
	TimestampType        TimestampType // https://cwiki.apache.org/confluence/display/KAFKA/KIP-32+-+Add+timestamps+to+Kafka+message
	IsTransactional      bool
	IsControlBatch       bool
	HasDeleteHorizonMs   bool
	LastOffsetDelta      int32
	BaseTimestamp        time.Time
	MaxTimestamp         time.Time
	ProducerId           int64
	ProducerEpoch        int16
	BaseSequence         int32
	RecordsLen           int
	Records              []byte
}

func (r *RecordBatch) encode(pe packetEncoder, version int16) (err error) {
	pe.putInt64(r.BaseOffset)
	pe.putInt32(r.BatchLength)
	pe.putInt32(r.PartitionLeaderEpoch)
	pe.putInt8(r.Magic)

	// TODO: This has to be dynamically calculated
	pe.putUint32(r.CRC)
	// TODO: This has to be dynamically calculated as well
	pe.putInt16(r.attributes)

	pe.putInt32(r.LastOffsetDelta)
	pe.putInt64(getMillisFromTime(r.BaseTimestamp))
	pe.putInt64(getMillisFromTime(r.MaxTimestamp))
	pe.putInt64(r.ProducerId)
	pe.putInt16(r.ProducerEpoch)
	pe.putInt32(r.BaseSequence)
	pe.putBytes(r.Records)

	return err
}

func (r *RecordBatch) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version

	if r.size, err = pd.getInt32(); err != nil {
		return err
	}

	if r.BaseOffset, err = pd.getInt64(); err != nil {
		return err
	}

	if r.BatchLength, err = pd.getInt32(); err != nil {
		return err
	}

	if r.PartitionLeaderEpoch, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Magic, err = pd.getInt8(); err != nil {
		return err
	}

	if r.CRC, err = pd.getUint32(); err != nil {
		return err
	}

	if r.attributes, err = pd.getInt16(); err != nil {
		return err
	}

	// parse attributes and append corresponding flags in the struct
	r.CompressionType = CompressionType(int8(r.attributes) & compressionCodecBit)
	// If the timestamp type is LogAppendType, the broker should set the record batch timestamp,
	// overriding the timestamp set by the client. The bindings won't set the timestamp automatically,
	// this will be left to the client software that consumes the library.
	r.TimestampType = TimestampType(int8(r.attributes) & timestampTypeBit)
	r.IsTransactional = r.attributes&isTransactionalBit == isTransactionalBit
	r.IsControlBatch = r.attributes&isControlBit == isControlBit

	if r.LastOffsetDelta, err = pd.getInt32(); err != nil {
		return err
	}

	baseTime, err := pd.getInt64()
	if err != nil {
		return err
	}
	r.BaseTimestamp = getTimeFromMillis(baseTime)

	maxTime, err := pd.getInt64()
	if err != nil {
		return err
	}
	r.MaxTimestamp = getTimeFromMillis(maxTime)

	if r.ProducerId, err = pd.getInt64(); err != nil {
		return err
	}

	if r.ProducerEpoch, err = pd.getInt16(); err != nil {
		return err
	}

	if r.BaseSequence, err = pd.getInt32(); err != nil {
		return err
	}

	if r.RecordsLen, err = pd.getArrayLength(); err != nil {
		return err
	}

	remainingBytes, err := pd.getRawBytes(pd.remaining())
	if err != nil {
		return err
	}

	r.Records = remainingBytes

	return err
}
