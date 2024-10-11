// protocol has been generated from message format json - DO NOT EDIT
package protocol

type DescribeTransactionsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// TransactionalIds contains a Array of transactionalIds to include in describe results. If empty, then no results will be returned.
	TransactionalIds []string
}

func (r *DescribeTransactionsRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putStringArray(r.TransactionalIds); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeTransactionsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.TransactionalIds, err = pd.getStringArray(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeTransactionsRequest) GetKey() int16 {
	return 65
}

func (r *DescribeTransactionsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeTransactionsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *DescribeTransactionsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeTransactionsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
