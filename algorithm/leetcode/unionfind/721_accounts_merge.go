package unionfind

import (
	"sort"
)

// Weighted union find with path compression
// Time: ~O(A alpha(A) + sum(log a_i))) = O(A alpha(A)). A is the number of emails for input account.
// a_i is the length of account[i]. The second part is for sorting.
// Space: O(B) B is number of unique emails
func accountsMerge(accounts [][]string) [][]string {
	if accounts == nil {
		return nil
	}
	parents := map[string]string{}
	size := map[string]int{}
	emailNameMap := map[string]string{}
	for _, account := range accounts {
		name := account[0]
		for i, email := range account {
			if i == 0 {
				continue
			}
			if _, ok := parents[email]; !ok {
				parents[email] = email
				size[email] = 1
				emailNameMap[email] = name
			}
			if i > 1 {
				union_721(parents, size, email, account[i-1])
			}
		}
	}

	preProcessedAccounts := map[string][]string{}
	for email, _ := range emailNameMap {
		rootEmail := find_721(parents, email)
		emails, ok := preProcessedAccounts[rootEmail]
		if !ok {
			preProcessedAccounts[rootEmail] = []string{email}
		} else {
			preProcessedAccounts[rootEmail] = append(emails, email)
		}
	}

	results := [][]string{}
	for rootEmail, emails := range preProcessedAccounts {
		name := emailNameMap[rootEmail]
		account := []string{name}
		sort.Strings(emails)
		results = append(results, append(account, emails...))
	}
	return results
}

func union_721(parents map[string]string, size map[string]int, email1, email2 string) {
	root1 := find_721(parents, email1)
	root2 := find_721(parents, email2)
	if root1 == root2 {
		return
	}
	size1 := size[root1]
	size2 := size[root2]
	if size1 < size2 {
		parents[root2] = root1
		size[root1] = size1 + size2
	} else {
		parents[root1] = root2
		size[root2] = size1 + size2
	}
	return
}

func find_721(parents map[string]string, email string) string {
	if parents[email] != email {
		parents[email] = find_721(parents, parents[email])
	}
	return parents[email]
}

// DFS
// Time: O(sum(a_i) + #_email + sum(log b_i))
// Space: O (sum (a_i))
func accountsMerge1(accounts [][]string) [][]string {
	if accounts == nil {
		return nil
	}
	// Build graph and record email-name map
	emailNameMap := map[string]string{}
	emailGraph := map[string][]string{}
	isVisited := map[string]bool{}
	for _, nameAndEmails := range accounts {
		name := nameAndEmails[0]
		emails := nameAndEmails[1:]
		for i, email := range emails {
			emailNameMap[email] = name
			isVisited[email] = false
			if _, ok := emailGraph[email]; !ok {
				emailGraph[email] = []string{}
			}
			// build graph
			if i > 0 {
				connect(emailGraph, email, emails[i-1])
				connect(emailGraph, emails[i-1], email)
			}
		}
	}

	// DFS. Construct new acount map.
	emailGroups := [][]string{}
	for emailFrom, _ := range emailGraph {
		if !isVisited[emailFrom] {
			emailGroup := dfs_721(emailGraph, isVisited, emailFrom)
			emailGroups = append(emailGroups, emailGroup)
		}
	}

	processedAccounts := [][]string{}
	for _, emailGroup := range emailGroups {
		account := []string{emailNameMap[emailGroup[0]]}
		sort.Strings(emailGroup)
		account = append(account, emailGroup...)
		processedAccounts = append(processedAccounts, account)
	}
	return processedAccounts
}

func connect(graph map[string][]string, emailFrom, emailTo string) {
	toList := graph[emailFrom]
	graph[emailFrom] = append(toList, emailTo)
}

func dfs_721(graph map[string][]string, isVisited map[string]bool, emailFrom string) []string {
	emailGroup := []string{emailFrom}
	isVisited[emailFrom] = true
	for _, emailTo := range graph[emailFrom] {
		if !isVisited[emailTo] {
			emails := dfs_721(graph, isVisited, emailTo)
			emailGroup = append(emailGroup, emails...)
		}
	}
	return emailGroup
}
