package ipranger

import (
	"github.com/khulnasoft-labs/mapcidr"
)

// Ips of a cidr
func Ips(cidr string) ([]string, error) {
	return mapcidr.IPAddresses(cidr)
}
