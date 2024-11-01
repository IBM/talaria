// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// TransactionState_ListTransactionsResponse contains a
type TransactionState_ListTransactionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TransactionalID contains a
	TransactionalID string
	// ProducerID contains a
	ProducerID int64
	// TransactionState contains the current transaction state of the producer
	TransactionState string
}

func (t *TransactionState_ListTransactionsResponse) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.TransactionalID); err != nil {
		return err
	}

	pe.putInt64(t.ProducerID)

	if err := pe.putString(t.TransactionState); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (t *TransactionState_ListTransactionsResponse) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.TransactionalID, err = pd.getString(); err != nil {
		return err
	}

	if t.ProducerID, err = pd.getInt64(); err != nil {
		return err
	}

	if t.TransactionState, err = pd.getString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type ListTransactionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains a
	ErrorCode int16
	// UnknownStateFilters contains a Set of state filters provided in the request which were unknown to the transaction coordinator
	UnknownStateFilters []string
	// TransactionStates contains a
	TransactionStates []TransactionState_ListTransactionsResponse
}

func (r *ListTransactionsResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.ThrottleTimeMs)

	pe.putInt16(r.ErrorCode)

	if err := pe.putStringArray(r.UnknownStateFilters); err != nil {
		return err
	}

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

func (r *ListTransactionsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.UnknownStateFilters, err = pd.getStringArray(); err != nil {
		return err
	}

	var numTransactionStates int
	if numTransactionStates, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTransactionStates > 0 {
		r.TransactionStates = make([]TransactionState_ListTransactionsResponse, numTransactionStates)
		for i := 0; i < numTransactionStates; i++ {
			var block TransactionState_ListTransactionsResponse
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

func (r *ListTransactionsResponse) GetKey() int16 {
	return 66
}

func (r *ListTransactionsResponse) GetVersion() int16 {
	return r.Version
}

func (r *ListTransactionsResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *ListTransactionsResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ListTransactionsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ListTransactionsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
