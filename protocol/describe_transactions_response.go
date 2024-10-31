// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// TopicData_DescribeTransactionsResponse contains the set of partitions included in the current transaction (if active). When a transaction is preparing to commit or abort, this will include only partitions which do not have markers.
type TopicData_DescribeTransactionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Topic contains a
	Topic string
	// Partitions contains a
	Partitions []int32
}

func (t *TopicData_DescribeTransactionsResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Topic); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.Partitions); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *TopicData_DescribeTransactionsResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Topic, err = pd.getString(); err != nil {
		return err
	}

	if t.Partitions, err = pd.getInt32Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// TransactionState_DescribeTransactionsResponse contains a
type TransactionState_DescribeTransactionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains a
	ErrorCode int16
	// TransactionalID contains a
	TransactionalID string
	// TransactionState contains a
	TransactionState string
	// TransactionTimeoutMs contains a
	TransactionTimeoutMs int32
	// TransactionStartTimeMs contains a
	TransactionStartTimeMs int64
	// ProducerID contains a
	ProducerID int64
	// ProducerEpoch contains a
	ProducerEpoch int16
	// Topics contains the set of partitions included in the current transaction (if active). When a transaction is preparing to commit or abort, this will include only partitions which do not have markers.
	Topics []TopicData_DescribeTransactionsResponse
}

func (t *TransactionState_DescribeTransactionsResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	pe.putInt16(t.ErrorCode)

	if err := pe.putString(t.TransactionalID); err != nil {
		return err
	}

	if err := pe.putString(t.TransactionState); err != nil {
		return err
	}

	pe.putInt32(t.TransactionTimeoutMs)

	pe.putInt64(t.TransactionStartTimeMs)

	pe.putInt64(t.ProducerID)

	pe.putInt16(t.ProducerEpoch)

	if err := pe.putArrayLength(len(t.Topics)); err != nil {
		return err
	}
	for _, block := range t.Topics {
		if err := block.encode(pe, t.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (t *TransactionState_DescribeTransactionsResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if t.TransactionalID, err = pd.getString(); err != nil {
		return err
	}

	if t.TransactionState, err = pd.getString(); err != nil {
		return err
	}

	if t.TransactionTimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	if t.TransactionStartTimeMs, err = pd.getInt64(); err != nil {
		return err
	}

	if t.ProducerID, err = pd.getInt64(); err != nil {
		return err
	}

	if t.ProducerEpoch, err = pd.getInt16(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		t.Topics = make([]TopicData_DescribeTransactionsResponse, numTopics)
		for i := 0; i < numTopics; i++ {
			var block TopicData_DescribeTransactionsResponse
			if err := block.decode(pd, t.Version); err != nil {
				return err
			}
			t.Topics[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeTransactionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// TransactionStates contains a
	TransactionStates []TransactionState_DescribeTransactionsResponse
}

func (r *DescribeTransactionsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.TransactionStates)); err != nil {
		return err
	}
	for _, block := range r.TransactionStates {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeTransactionsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numTransactionStates int
	if numTransactionStates, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTransactionStates > 0 {
		r.TransactionStates = make([]TransactionState_DescribeTransactionsResponse, numTransactionStates)
		for i := 0; i < numTransactionStates; i++ {
			var block TransactionState_DescribeTransactionsResponse
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.TransactionStates[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeTransactionsResponse) GetKey() int16 {
	return 65
}

func (r *DescribeTransactionsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DescribeTransactionsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *DescribeTransactionsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeTransactionsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DescribeTransactionsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
