// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse contains a
type AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// ResultsByPartition contains the results for each partition
	ResultsByPartition []AddPartitionsToTxnPartitionResult
}

func (a *AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	if err := pe.putString(a.Name); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(a.ResultsByPartition)); err != nil {
		return err
	}
	for _, block := range a.ResultsByPartition {
		if err := block.encode(pe, a.Version); err != nil {
			return err
		}
	}

	if a.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (a *AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.Name, err = pd.getString(); err != nil {
		return err
	}

	var numResultsByPartition int
	if numResultsByPartition, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numResultsByPartition > 0 {
		a.ResultsByPartition = make([]AddPartitionsToTxnPartitionResult, numResultsByPartition)
		for i := 0; i < numResultsByPartition; i++ {
			var block AddPartitionsToTxnPartitionResult
			if err := block.decode(pd, a.Version); err != nil {
				return err
			}
			a.ResultsByPartition[i] = block
		}
	}

	if a.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// AddPartitionsToTxnPartitionResult contains a
type AddPartitionsToTxnPartitionResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// PartitionIndex contains the partition indexes.
	PartitionIndex int32
	// PartitionErrorCode contains the response error code.
	PartitionErrorCode int16
}

func (a *AddPartitionsToTxnPartitionResult) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	pe.putInt32(a.PartitionIndex)

	pe.putInt16(a.PartitionErrorCode)

	if a.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (a *AddPartitionsToTxnPartitionResult) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.PartitionIndex, err = pd.getInt32(); err != nil {
		return err
	}

	if a.PartitionErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if a.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// AddPartitionsToTxnResult contains a Results categorized by transactional ID.
type AddPartitionsToTxnResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TransactionalID contains the transactional id corresponding to the transaction.
	TransactionalID string
	// TopicResults contains the results for each topic.
	TopicResults []AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse
}

func (r *AddPartitionsToTxnResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		if err := pe.putString(r.TransactionalID); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		if err := pe.putArrayLength(len(r.TopicResults)); err != nil {
			return err
		}
		for _, block := range r.TopicResults {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AddPartitionsToTxnResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		if r.TransactionalID, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		var numTopicResults int
		if numTopicResults, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numTopicResults > 0 {
			r.TopicResults = make([]AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse, numTopicResults)
			for i := 0; i < numTopicResults; i++ {
				var block AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.TopicResults[i] = block
			}
		}
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type AddPartitionsToTxnResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains a Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the response top level error code.
	ErrorCode int16
	// ResultsByTransaction contains a Results categorized by transactional ID.
	ResultsByTransaction []AddPartitionsToTxnResult
	// ResultsByTopicV3AndBelow contains the results for each topic.
	ResultsByTopicV3AndBelow []AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse
}

func (r *AddPartitionsToTxnResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	if r.Version >= 4 {
		pe.putInt16(r.ErrorCode)
	}

	if r.Version >= 4 {
		if err := pe.putArrayLength(len(r.ResultsByTransaction)); err != nil {
			return err
		}
		for _, block := range r.ResultsByTransaction {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		if err := pe.putArrayLength(len(r.ResultsByTopicV3AndBelow)); err != nil {
			return err
		}
		for _, block := range r.ResultsByTopicV3AndBelow {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AddPartitionsToTxnResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 4 {
		if r.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		var numResultsByTransaction int
		if numResultsByTransaction, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numResultsByTransaction > 0 {
			r.ResultsByTransaction = make([]AddPartitionsToTxnResult, numResultsByTransaction)
			for i := 0; i < numResultsByTransaction; i++ {
				var block AddPartitionsToTxnResult
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.ResultsByTransaction[i] = block
			}
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		var numResultsByTopicV3AndBelow int
		if numResultsByTopicV3AndBelow, err = pd.getArrayLength(); err != nil {
			return err
		}
		if numResultsByTopicV3AndBelow > 0 {
			r.ResultsByTopicV3AndBelow = make([]AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse, numResultsByTopicV3AndBelow)
			for i := 0; i < numResultsByTopicV3AndBelow; i++ {
				var block AddPartitionsToTxnTopicResult_AddPartitionsToTxnResponse
				if err := block.decode(pd, r.Version); err != nil {
					return err
				}
				r.ResultsByTopicV3AndBelow[i] = block
			}
		}
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *AddPartitionsToTxnResponse) GetKey() int16 {
	return 24
}

func (r *AddPartitionsToTxnResponse) GetVersion() int16 {
	return r.Version
}

func (r *AddPartitionsToTxnResponse) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 1
	}
	return 0
}

func (r *AddPartitionsToTxnResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *AddPartitionsToTxnResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *AddPartitionsToTxnResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
