package jwt

import (
	"fmt"
	"testing"
)

func TestParseJwtWithClaims(t *testing.T) {
	jwtStr := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJVdWlkIjoiMjAyMDExMjcxNDQ2MjQ1NzJ4WGtiV0J1S1kxMDAwMzEiLCJWZXJzaW9uIjoiOC4yLjAiLCJEZXZpY2VJZCI6ImQ4NDBhMjY2NWJkN2Q0OTAzMjc5N2MxZmEyODZjYTAyMTEwZCIsIk1hY0FkZHIiOiIwMDpEQjpBNzozRTo5RTowOCIsIlNka0ludFZlciI6IjI4IiwiZXhwIjoxNzEyNTc0ODQyfQ.ZmsoLXqCIUOs6c1lxKNHhGBV2Tt86nUNLMVz3sQmFHLse_eRmygegFPNq7-CplwTtANPkHhZacU9vevVoaWoi7VI0IlX1z2O7CySAVh8EfaDlw6X51EaMUX49Hhrw_HwFB_r4_3wpEqsklMgYya37qRMURAcpn78BTQYCWNnOC8"
	claims, err := ParseJwtWithClaims(jwtStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(claims)
}
