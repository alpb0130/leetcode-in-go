package withoutlibfunction

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// The same as another method but doesn't use the lib function. We use a&-a to set all digit to 0
// except the lowest 1. And we use bit manipulation to get the power.
// Time: O(logn) ??
// Space: O(1)
func ipToCIDR(ip string, n int) []string {
	ipNum := ipToNum(ip)
	res := []string{}
	for n > 0 {
		// spacePower := minInt(ipSpace(ipNum), desiredSpace(uint32(n)))
		spacePower := minInt(desiredSpace(ipNum&-ipNum), desiredSpace(uint32(n)))
		mask := 32 - spacePower
		res = append(res, numToIP(ipNum, mask))
		ipNum += 1 << spacePower
		n -= 1 << spacePower
	}
	return res
}

func ipToNum(ip string) uint32 {
	subIPs := strings.Split(ip, ".")
	var res uint32
	for _, subIP := range subIPs {
		num, _ := strconv.Atoi(subIP)
		res = res*256 + uint32(num)
	}
	return res
}

func numToIP(num uint32, mask uint32) string {
	return fmt.Sprintf(
		"%d.%d.%d.%d/%d",
		num>>24,
		num>>16%256,
		num>>8%256,
		num%256,
		mask,
	)
}

func desiredSpace(n uint32) uint32 {
	var res uint32
	for n > 0 {
		n >>= 1
		res++
	}
	return res - 1
}

func minInt(a, b uint32) uint32 {
	return uint32(math.Min(float64(a), float64(b)))
}
