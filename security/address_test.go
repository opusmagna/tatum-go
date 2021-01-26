package security

import (
	"fmt"
	"testing"
)

func TestCheckMaliciousAddress(t *testing.T) {
	res := CheckMaliciousAddress("rHb9CJAWyB4rj91VRWn96DkukG4bwdty")
	fmt.Println(res)
}
