package lib

import "strings"

// Use some lib func
// Time: O(m*n) - m is the number of emails. n is the length of email.
// Space: O(m*n)
func numUniqueEmails(emails []string) int {
	emailMap := map[string]bool{}
	for _, email := range emails {
		processedEmail := processEmail(email)
		emailMap[processedEmail] = true
	}
	return len(emailMap)
}

func processEmail(email string) string {
	names := strings.Split(email, "@")
	idx := strings.Index(names[0], "+")
	if idx != -1 {
		names[0] = names[0][:idx]
	}
	names[0] = strings.Replace(names[0], ".", "", -1)
	return names[0] + "@" + names[1]
}
