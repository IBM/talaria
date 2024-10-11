// protocol has been generated from message format json - DO NOT EDIT
package protocol

type DescribeGroupsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Groups contains the names of the groups to describe
	Groups []string
	// IncludeAuthorizedOperations contains a Whether to include authorized operations.
	IncludeAuthorizedOperations bool
}

func (r *DescribeGroupsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 5 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putStringArray(r.Groups); err != nil {
		return err
	}

	if r.Version >= 3 {
		pe.putBool(r.IncludeAuthorizedOperations)
	}

	if r.Version >= 5 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DescribeGroupsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 5 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Groups, err = pd.getStringArray(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if r.IncludeAuthorizedOperations, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.Version >= 5 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DescribeGroupsRequest) GetKey() int16 {
	return 15
}

func (r *DescribeGroupsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeGroupsRequest) GetHeaderVersion() int16 {
	if r.Version >= 5 {
		return 2
	}
	return 1
}

func (r *DescribeGroupsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 5
}

func (r *DescribeGroupsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
