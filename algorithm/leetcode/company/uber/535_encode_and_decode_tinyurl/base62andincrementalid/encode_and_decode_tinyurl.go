package base62andincrementalid

import (
	"math/rand"
	"strings"
	"time"
)

var set = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var pattern = "www.url.com/"

type Codec struct {
	shortToLong map[string]string
	counter     int
	r           *rand.Rand
	// We can also add map from long to short to avoid duplication
}

func NewCodec() *Codec {
	return &Codec{map[string]string{}, 1, rand.New(rand.NewSource(time.Now().Unix()))}
}

func (c *Codec) getCodeFixLength() string {
	var url strings.Builder
	for i := 0; i < 6; i++ {
		url.WriteString(string(set[c.r.Intn(62)]))
	}
	return url.String()
}

func (c *Codec) getCodeByID(counter int) string {
	var url strings.Builder
	for counter > 0 {
		counter--
		url.WriteString(string(set[counter%len(set)]))
		counter /= 62
	}
	return url.String()
}

func (c *Codec) Encode(longUrl string) string {
	url := c.getCodeFixLength()
	//c.counter++
	c.shortToLong[url] = longUrl
	return pattern + url
}

func (c *Codec) Decode(shortUrl string) string {
	url := strings.Replace(shortUrl, pattern, "", -1)
	return c.shortToLong[url]
}
