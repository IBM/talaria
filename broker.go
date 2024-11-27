package main

import (
	"fmt"
	"net"
	"net/netip"
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

	advertisedListeners := strings.Split(strings.ReplaceAll(utils.GetEnvVar("advertised.listeners", ""), " ", ""), ",")
	listeners := strings.Split(strings.ReplaceAll(utils.GetEnvVar("listeners", ""), " ", ""), ",")

	if len(advertisedListeners) == 0 {
		advertisedListeners = listeners
	}

	listenersArray, err := parseListeners(listeners)
	if err != nil {
		return Broker{}, err
	}

	broker.Listeners = append(broker.Listeners, listenersArray...)

	err = broker.validateListeners()
	if err != nil {
		return Broker{}, err
	}

	advertisedListenersArr, err := parseListeners(advertisedListeners)
	if err != nil {
		return Broker{}, err
	}
	broker.AdvertisedListeners = append(broker.AdvertisedListeners, advertisedListenersArr...)

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

func parseListeners(listeners []string) ([]Listener, error) {
	result := []Listener{}

	for _, l := range listeners {
		if l == "" {
			continue
		}

		listener, err := parseListener(l)
		if err != nil {
			return []Listener{}, err
		}

		result = append(result, listener)
	}

	return result, nil
}

func parseListener(l string) (Listener, error) {
	listener, err := url.Parse(l)
	if err != nil {
		return Listener{}, err
	}

	// parse the security protocol from the url scheme.
	// If the protocol is unknown treat the scheme as broker name and check the listener.security.protocol.map
	listenerName, securityProtocol, err := getBrokerNameComponents(listener.Scheme)
	if err != nil {
		return Listener{}, err
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

// getBrokerNameComponents checks if the broker name, inferred from the URL schema is a valid security protocol.
// If not, it checks the listener.security.protocol.map for mapping for custom broker names and returns the broker name/security protocol pair.
// If no mapping is found in the case of custom broker name, the function returns an error.
func getBrokerNameComponents(s string) (string, utils.SecurityProtocol, error) {
	securityProtocol, ok := utils.ParseSecurityProtocol(s)

	if ok {
		return s, securityProtocol, nil
	} else {
		// the listener schema is not a known security protocol, treat is as broker name
		// and extract the security protocol from listener.security.protocol.map
		spm := strings.Split(strings.ReplaceAll(utils.GetEnvVar("listener.security.protocol.map", ""), " ", ""), ",")

		for _, sp := range spm {
			components := strings.Split(sp, ":")

			if strings.EqualFold(s, components[0]) {
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

// validateListeners performs common checks on the listeners as per Kafka specification https://kafka.apache.org/documentation/#brokerconfigs_listeners.
// Broker name and port pairs have to be unique. The exception is if the host for two entries is IPv4 and IPv6 respectively.
func (b *Broker) validateListeners() error {
	checkNamePortPair := make([]map[string]string, 0)
	for _, listener := range b.Listeners {
		npPair := fmt.Sprintf("%s:%d", listener.ListenerName, listener.Port)
		for _, npp := range checkNamePortPair {
			if host, ok := npp[npPair]; ok {
				// the broker name/port pair is duplicated. Check if one is IPv4 and the other IPv6, otherwise return an error.
				addr1, _ := netip.ParseAddr(host) // ignore errors from ParseAddr, which will be thrown if a hostname is provided, we care only about IP addresses.
				existingAddrIPVer := addr1.Is4()

				addr2, _ := netip.ParseAddr(listener.Host)
				newAddrIPVer := addr2.Is4()

				if existingAddrIPVer == newAddrIPVer {
					return fmt.Errorf("listener name and port are not unique for listener %s", listener.ListenerName)
				}
			}
		}
		checkNamePortPair = append(checkNamePortPair, map[string]string{npPair: listener.Host})
	}
	return nil
}
