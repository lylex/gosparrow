package utils

import (
	"strings"
	"time"

	"github.com/twinj/uuid"
)

// Now returns current UTC time, formatted with RFC3339
func Now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// DaysAgo returns n days ago from now in UTC time, formatted with RFC3339
func DaysAgo(n int) string {
	return time.Now().AddDate(0, 0, -n).UTC().Format(time.RFC3339)
}

// UUID generates a UUID without error
func UUID() string {
	newUUID := uuid.NewV4()
	return newUUID.String()
}

// Contains check whether a string slice contains a string
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// ContainsStr used to assert whether a string contains a substring
//   and it is case insensitive
func ContainsStr(s, substr string) bool {
	return strings.Contains(strings.ToUpper(s), strings.ToUpper(substr))
}
