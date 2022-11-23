package unitpay

var UnitpayAllowedIPs = map[string]interface{}{"31.186.100.49": nil, "52.29.152.23": nil, "52.19.56.234": nil}

func CheckIPAllowed(IPAddress string) bool {
	if _, allowed := UnitpayAllowedIPs[IPAddress]; allowed {
		return allowed
	}
	return false
}
