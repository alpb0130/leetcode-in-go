package loop

import "strings"

// Process string without using lib function
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
	var newEmail strings.Builder
	localName := []byte{}
	for i := 0; i < len(names[0]); i++ {
		if names[0][i] == '+' {
			break
		}
		if names[0][i] == '.' {
			continue
		}
		localName = append(localName, names[0][i])
	}
	newEmail.WriteString(string(localName))
	newEmail.WriteString("@")
	newEmail.WriteString(names[1])
	return newEmail.String()
}
