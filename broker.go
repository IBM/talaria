package main

import (
	"fmt"
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

	// regex used to validate hostname setting. Accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	hostnameRegexStringRFC1123 = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?$`
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
}

// NewBroker returns a new instance of Broker.
// For now OpenTalaria does not support rack awareness, but this will change in the future.
func NewBroker() (Broker, error) {
	broker := Broker{}

	advertisedListeners := utils.GetEnvVar("advertised.listeners", "")
	listeners := utils.GetEnvVar("listeners", "")

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

	advertisedListenersArr := strings.Split(advertisedListeners, ",")
	for _, l := range advertisedListenersArr {
		listener, err := parseListener(l)
		if err != nil {
			return Broker{}, err
		}

		broker.AdvertisedListeners = append(broker.Listeners, listener)
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

// TODO: implement the Kafka custom schema.
// TODO: advertised.listeners cannot bind to 0.0.0.0, add this validation during parsing. See https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#advertised-listeners
func parseListener(l string) (Listener, error) {
	listener, err := url.Parse(l)
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
		Host: host,
		Port: int32(parsedPort),
	}, nil
}
