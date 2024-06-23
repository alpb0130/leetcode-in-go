package _map

import (
	"fmt"
	"strconv"
	"strings"
)

// Time: O(m*n)
// Space: O(n)
func subdomainVisits(cpdomains []string) []string {
	domainCnt := map[string]int{}
	for _, cpdomain := range cpdomains {
		cp := strings.Split(cpdomain, " ")
		cnt, _ := strconv.Atoi(cp[0])
		domain := cp[1]
		domainCnt[domain] += cnt
		for i := 0; i < len(domain); i++ {
			if domain[i] == '.' {
				domainCnt[domain[i+1:]] += cnt
			}
		}
	}
	res := []string{}
	for domain, cnt := range domainCnt {
		res = append(res, fmt.Sprintf("%d %s", cnt, domain))
	}
	return res
}
