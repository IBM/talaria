// protocol has been generated from message format json - DO NOT EDIT
package protocol

type ListGroupsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// StatesFilter contains the states of the groups we want to list. If empty all groups are returned with their state.
	StatesFilter []string
}

func (r *ListGroupsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 4 {
		if err := pe.putStringArray(r.StatesFilter); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ListGroupsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 4 {
		if r.StatesFilter, err = pd.getStringArray(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ListGroupsRequest) GetKey() int16 {
	return 16
}

func (r *ListGroupsRequest) GetVersion() int16 {
	return r.Version
}

func (r *ListGroupsRequest) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 2
	}
	return 1
}

func (r *ListGroupsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *ListGroupsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
