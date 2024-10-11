package protocol

import "time"

func getTimeFromMillis(millis int64) time.Time {
	// since the Kafka protocol can return a negative value in the time field,
	// which is an invalid integer to convert to time.Time
	// we return an empty time.Time object in this case.
	timestamp := time.Time{}
	if millis >= 0 {
		timestamp = time.Unix(millis/1000, (millis%1000)*int64(time.Millisecond))
	}

	return timestamp
}

func getMillisFromTime(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
