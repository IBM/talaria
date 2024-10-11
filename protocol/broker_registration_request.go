// protocol has been generated from message format json - DO NOT EDIT
package protocol

import uuid "github.com/google/uuid"

// Listener contains the listeners of this broker
type Listener struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the name of the endpoint.
	Name string
	// Host contains the hostname.
	Host string
	// Port contains the port.
	Port uint16
	// SecurityProtocol contains the security protocol.
	SecurityProtocol int16
}

func (l *Listener) encode(pe packetEncoder, version int16) (err error) {
	l.Version = version
	if err := pe.putString(l.Name); err != nil {
		return err
	}

	if err := pe.putString(l.Host); err != nil {
		return err
	}

	pe.putUint16(l.Port)

	pe.putInt16(l.SecurityProtocol)

	pe.putUVarint(0)
	return nil
}

func (l *Listener) decode(pd packetDecoder, version int16) (err error) {
	l.Version = version
	if l.Name, err = pd.getString(); err != nil {
		return err
	}

	if l.Host, err = pd.getString(); err != nil {
		return err
	}

	if l.Port, err = pd.getUint16(); err != nil {
		return err
	}

	if l.SecurityProtocol, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

// Feature contains the features on this broker
type Feature struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the feature name.
	Name string
	// MinSupportedVersion contains the minimum supported feature level.
	MinSupportedVersion int16
	// MaxSupportedVersion contains the maximum supported feature level.
	MaxSupportedVersion int16
}

func (f *Feature) encode(pe packetEncoder, version int16) (err error) {
	f.Version = version
	if err := pe.putString(f.Name); err != nil {
		return err
	}

	pe.putInt16(f.MinSupportedVersion)

	pe.putInt16(f.MaxSupportedVersion)

	pe.putUVarint(0)
	return nil
}

func (f *Feature) decode(pd packetDecoder, version int16) (err error) {
	f.Version = version
	if f.Name, err = pd.getString(); err != nil {
		return err
	}

	if f.MinSupportedVersion, err = pd.getInt16(); err != nil {
		return err
	}

	if f.MaxSupportedVersion, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

type BrokerRegistrationRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// BrokerID contains the broker ID.
	BrokerID int32
	// ClusterID contains the cluster id of the broker process.
	ClusterID string
	// IncarnationID contains the incarnation id of the broker process.
	IncarnationID uuid.UUID
	// Listeners contains the listeners of this broker
	Listeners []Listener
	// Features contains the features on this broker
	Features []Feature
	// Rack contains the rack which this broker is in.
	Rack *string
	// IsMigratingZkBroker contains a If the required configurations for ZK migration are present, this value is set to true
	IsMigratingZkBroker bool
}

func (r *BrokerRegistrationRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	pe.putInt32(r.BrokerID)

	if err := pe.putString(r.ClusterID); err != nil {
		return err
	}

	if err := pe.putUUID(r.IncarnationID); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Listeners)); err != nil {
		return err
	}
	for _, block := range r.Listeners {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.Features)); err != nil {
		return err
	}
	for _, block := range r.Features {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if err := pe.putNullableString(r.Rack); err != nil {
		return err
	}

	if r.Version >= 1 {
		pe.putBool(r.IsMigratingZkBroker)
	}

	pe.putUVarint(0)
	return nil
}

func (r *BrokerRegistrationRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.BrokerID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.ClusterID, err = pd.getString(); err != nil {
		return err
	}

	if r.IncarnationID, err = pd.getUUID(); err != nil {
		return err
	}

	var numListeners int
	if numListeners, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Listeners = make([]Listener, numListeners)
	for i := 0; i < numListeners; i++ {
		var block Listener
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Listeners[i] = block
	}

	var numFeatures int
	if numFeatures, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Features = make([]Feature, numFeatures)
	for i := 0; i < numFeatures; i++ {
		var block Feature
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Features[i] = block
	}

	if r.Rack, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.IsMigratingZkBroker, err = pd.getBool(); err != nil {
			return err
		}
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *BrokerRegistrationRequest) GetKey() int16 {
	return 62
}

func (r *BrokerRegistrationRequest) GetVersion() int16 {
	return r.Version
}

func (r *BrokerRegistrationRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *BrokerRegistrationRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *BrokerRegistrationRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
