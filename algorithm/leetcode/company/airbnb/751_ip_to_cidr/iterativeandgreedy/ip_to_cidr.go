package iterativeandgreedy

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
	"strings"
)

// The problem requires us to convert a IP to CIDR give the number of IP address
// we need. The key point is to identify the mask length and the next base IP address
// we have. For a give IP, the size of IP range this IP can represent is
// 2^(# trailing zeros in binary) and the mask is 32-(# trailing zeros in binary).
// Let's say it m. If we want n IPs and n > m. We need find the next IP and it's mask.
// Then the next IP would be current IP+m. What if n < m? Then we cannot use all IPs
// in the range. The range we need is 2^(position of most significant 1). Because the mask
// can only represent range of 2^x and we always want the result to be as short as possible,
// we need to take care of the most significant 1 first.
// Time: O(logn) ??
// Space: O(1)
func ipToCIDR(ip string, n int) []string {
	ipNum := ipToNum(ip)
	res := []string{}
	for n > 0 {
		spacePower := minInt(ipSpace(ipNum), desiredSpace(uint32(n)))
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

func ipSpace(ip uint32) uint32 {
	return uint32(bits.TrailingZeros32(ip))
}

func desiredSpace(n uint32) uint32 {
	return 31 - uint32(bits.LeadingZeros32(n))
}

func minInt(a, b uint32) uint32 {
	return uint32(math.Min(float64(a), float64(b)))
}
