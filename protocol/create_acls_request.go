// protocol has been generated from message format json - DO NOT EDIT
package protocol

// AclCreation contains the ACLs that we want to create.
type AclCreation struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ResourceType contains the type of the resource.
	ResourceType int8
	// ResourceName contains the resource name for the ACL.
	ResourceName string
	// ResourcePatternType contains the pattern type for the ACL.
	ResourcePatternType int8
	// Principal contains the principal for the ACL.
	Principal string
	// Host contains the host for the ACL.
	Host string
	// Operation contains the operation type for the ACL (read, write, etc.).
	Operation int8
	// PermissionType contains the permission type for the ACL (allow, deny, etc.).
	PermissionType int8
}

func (c *AclCreation) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	pe.putInt8(c.ResourceType)

	if err := pe.putString(c.ResourceName); err != nil {
		return err
	}

	if c.Version >= 1 {
		pe.putInt8(c.ResourcePatternType)
	}

	if err := pe.putString(c.Principal); err != nil {
		return err
	}

	if err := pe.putString(c.Host); err != nil {
		return err
	}

	pe.putInt8(c.Operation)

	pe.putInt8(c.PermissionType)

	if c.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (c *AclCreation) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.ResourceType, err = pd.getInt8(); err != nil {
		return err
	}

	if c.ResourceName, err = pd.getString(); err != nil {
		return err
	}

	if c.Version >= 1 {
		if c.ResourcePatternType, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if c.Principal, err = pd.getString(); err != nil {
		return err
	}

	if c.Host, err = pd.getString(); err != nil {
		return err
	}

	if c.Operation, err = pd.getInt8(); err != nil {
		return err
	}

	if c.PermissionType, err = pd.getInt8(); err != nil {
		return err
	}

	if c.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAclsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Creations contains the ACLs that we want to create.
	Creations []AclCreation
}

func (r *CreateAclsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Creations)); err != nil {
		return err
	}
	for _, block := range r.Creations {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *CreateAclsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numCreations int
	if numCreations, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numCreations > 0 {
		r.Creations = make([]AclCreation, numCreations)
		for i := 0; i < numCreations; i++ {
			var block AclCreation
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Creations[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *CreateAclsRequest) GetKey() int16 {
	return 30
}

func (r *CreateAclsRequest) GetVersion() int16 {
	return r.Version
}

func (r *CreateAclsRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *CreateAclsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *CreateAclsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
