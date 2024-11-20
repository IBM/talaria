package main

import (
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"opentalaria/utils"
	"strconv"
	"strings"
)

var (
	// by default in KRaft mode, generated broker IDs start from reserved.broker.max.id + 1,
	// where reserved.broker.max.id=1000 if the property is not set.
	// KRaft mode is the default Kafka mode, since Kafka v3.3.1, so OpenTalaria will implement default settings in KRaft mode.
	RESERVED_BROKER_MAX_ID = 1000
)

type Broker struct {
	BrokerID int32
	Rack     *string
	// https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#listeners
	Listeners []Listener
	// https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#advertised-listeners
	AdvertisedListeners []Listener
}

type Listener struct {
	Host string
	Port int32
	// If the listener name is a security protocol, like PLAINTEXT,SSL,SASL_PLAINTEXT,SASL_SSL,
	// the name will be set as SecurityProtocol. Otherwise the name should be mapped in listener.security.protocol.map.
	// see https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#listener-security-protocol-map.
	SecurityProtocol utils.SecurityProtocol
	ListenerName     string
}

// NewBroker returns a new instance of Broker.
// For now OpenTalaria does not support rack awareness, but this will change in the future.
func NewBroker() (Broker, error) {
	broker := Broker{}

	advertisedListeners := strings.ReplaceAll(utils.GetEnvVar("advertised.listeners", ""), " ", "")
	listeners := strings.ReplaceAll(utils.GetEnvVar("listeners", ""), " ", "")

	if advertisedListeners == "" {
		advertisedListeners = listeners
	}

	listernersArr := strings.Split(listeners, ",")
	for _, l := range listernersArr {
		listener, err := parseListener(l)
		if err != nil {
			return Broker{}, err
		}

		broker.Listeners = append(broker.Listeners, listener)
	}

	err := broker.validateListeners()
	if err != nil {
		return Broker{}, err
	}

	advertisedListenersArr := strings.Split(advertisedListeners, ",")
	for _, l := range advertisedListenersArr {
		listener, err := parseListener(l)
		if err != nil {
			return Broker{}, err
		}

		broker.AdvertisedListeners = append(broker.AdvertisedListeners, listener)
	}

	brokerIdSetting := utils.GetEnvVar("broker.id", "-1")

	brokerId, err := strconv.Atoi(brokerIdSetting)
	if err != nil {
		return broker, fmt.Errorf("error parsing broker.id: %s", err)
	}

	// validate Broker ID
	if brokerId > RESERVED_BROKER_MAX_ID {
		return broker, fmt.Errorf("the configured node ID is greater than `reserved.broker.max.id`. Please adjust the `reserved.broker.max.id` setting. [%d > %d]",
			brokerId,
			RESERVED_BROKER_MAX_ID)
	}

	if brokerId == -1 {
		brokerId = RESERVED_BROKER_MAX_ID + 1
	}

	broker.BrokerID = int32(brokerId)

	return broker, nil
}

func parseListener(l string) (Listener, error) {
	listener, err := url.Parse(l)
	if err != nil {
		return Listener{}, err
	}

	// parse the security protocol from the url scheme.
	// If the protocol is unknown treat the scheme as broker name and check the listener.security.protocol.map
	securityProtocol, ok := utils.ParseSecurityProtocol(listener.Scheme)
	listenerName := listener.Scheme
	if !ok {
		// the url scheme
		slog.Debug("check listener.security.protocol.map")
	}

	host, port, err := net.SplitHostPort(listener.Host)
	if err != nil {
		return Listener{}, err
	}

	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		return Listener{}, err
	}

	return Listener{
		Host:             host,
		Port:             int32(parsedPort),
		SecurityProtocol: securityProtocol,
		ListenerName:     listenerName,
	}, nil
}

func parseBrokerName(s string) (string, utils.SecurityProtocol, error) {
	securityProtocol, ok := utils.ParseSecurityProtocol(s)

	if ok {
		return s, securityProtocol, nil
	} else {
		// the listener schema is not a known security protocol, treat is as broker name
		// and extract the security protocol from listener.security.protocol.map
		spm := strings.ReplaceAll(utils.GetEnvVar("listener.security.protocol.map", ""), " ", "")
		spMapArray := strings.Split(spm, ",")

		for _, sp := range spMapArray {
			components := strings.Split(sp, ":")

			if s == components[0] {
				securityProtocol, ok := utils.ParseSecurityProtocol(components[1])
				if !ok {
					return "", utils.UNDEFINED_SECURITY_PROTOCOL, fmt.Errorf("unknown security protocol for listener %s", components[0])
				}

				return s, securityProtocol, nil
			}
		}
	}

	return "", utils.UNDEFINED_SECURITY_PROTOCOL, fmt.Errorf("broker %s not found in listener.security.protocol.map", s)
}

func (b *Broker) validateListeners() error {
	return nil
}
