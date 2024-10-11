// protocol has been generated from message format json - DO NOT EDIT
package protocol

type DeleteGroupsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupsNames contains the group names to delete.
	GroupsNames []string
}

func (r *DeleteGroupsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putStringArray(r.GroupsNames); err != nil {
		return err
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeleteGroupsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.GroupsNames, err = pd.getStringArray(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DeleteGroupsRequest) GetKey() int16 {
	return 42
}

func (r *DeleteGroupsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DeleteGroupsRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *DeleteGroupsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *DeleteGroupsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
