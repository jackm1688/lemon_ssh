package mcmd

import "testing"

func TestQueryIP(t *testing.T) {
	ips := []string{"27.38.33.41","27.38.56.42","33.38.78.43","22.38.56.44","27.38.69.45"}
	QueryIp(3,ips...)
}
