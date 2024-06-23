package browser_history

type BrowserHistory struct {
	history []string
	current int
}

func NewBrowserHistory() *BrowserHistory {
	return &BrowserHistory{
		history: []string{},
		current: -1,
	}
}

func (b *BrowserHistory) GoToLink(link string) {
	b.history = append(b.history[:b.current+1], link)
	b.current = len(b.history) - 1
}

func (b *BrowserHistory) Forward() bool {
	if b.current+1 >= len(b.history) {
		return false
	}
	b.current++
	return true
}

func (b *BrowserHistory) JumpForward(n int) bool {
	if b.current+n >= len(b.history) {
		return false
	}
	b.current += n
	return true
}

func (b *BrowserHistory) Back() bool {
	if b.current-1 < 0 {
		return false
	}
	b.current--
	return true
}

func (b *BrowserHistory) JumpBack(n int) bool {
	if b.current-n < 0 {
		return false
	}
	b.current -= n
	return true
}

func (b *BrowserHistory) Current() (string, bool) {
	if b.current >= len(b.history) || b.current < 0 {
		return "", false
	}
	return b.history[b.current], true
}
