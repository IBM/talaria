// protocol has been generated from message format json - DO NOT EDIT
package protocol

// UserName contains the users to describe, or null/empty to describe all users.
type UserName struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the user name.
	Name string
}

func (u *UserName) encode(pe packetEncoder, version int16) (err error) {
	u.Version = version
	if err := pe.putString(u.Name); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (u *UserName) decode(pd packetDecoder, version int16) (err error) {
	u.Version = version
	if u.Name, err = pd.getString(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type DescribeUserScramCredentialsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Users contains the users to describe, or null/empty to describe all users.
	Users []UserName
}

func (r *DescribeUserScramCredentialsRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putArrayLength(len(r.Users)); err != nil {
		return err
	}
	for _, block := range r.Users {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	pe.putUVarint(0)
	return nil
}

func (r *DescribeUserScramCredentialsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	var numUsers int
	if numUsers, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numUsers > 0 {
		r.Users = make([]UserName, numUsers)
		for i := 0; i < numUsers; i++ {
			var block UserName
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Users[i] = block
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *DescribeUserScramCredentialsRequest) GetKey() int16 {
	return 50
}

func (r *DescribeUserScramCredentialsRequest) GetVersion() int16 {
	return r.Version
}

func (r *DescribeUserScramCredentialsRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *DescribeUserScramCredentialsRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *DescribeUserScramCredentialsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
