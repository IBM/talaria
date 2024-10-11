// protocol has been generated from message format json - DO NOT EDIT
package protocol

type ListTransactionsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// StateFilters contains the transaction states to filter by: if empty, all transactions are returned; if non-empty, then only transactions matching one of the filtered states will be returned
	StateFilters []string
	// ProducerIdFilters contains the producerIds to filter by: if empty, all transactions will be returned; if non-empty, only transactions which match one of the filtered producerIds will be returned
	ProducerIdFilters []int64
}

func (r *ListTransactionsRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putStringArray(r.StateFilters); err != nil {
		return err
	}

	if err := pe.putInt64Array(r.ProducerIdFilters); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *ListTransactionsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.StateFilters, err = pd.getStringArray(); err != nil {
		return err
	}

	if r.ProducerIdFilters, err = pd.getInt64Array(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *ListTransactionsRequest) GetKey() int16 {
	return 66
}

func (r *ListTransactionsRequest) GetVersion() int16 {
	return r.Version
}

func (r *ListTransactionsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *ListTransactionsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *ListTransactionsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
