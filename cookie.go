package cookiejar

import (
	"encoding/json"
	"strings"
)

type CookieSet struct {
	Site string
	Set  map[string]*Entry
}

func NewCookieSet(site string) *CookieSet {
	ret := &CookieSet{
		Site: site,
		Set:  make(map[string]*Entry, 0),
	}
	return ret
}

func (self *CookieSet) AddCookie(cookie *Entry) {
	if strings.HasPrefix(cookie.Domain, ".") {
		cookie.Domain = cookie.Domain[1:len(cookie.Domain)]
	}
	key := cookie.Domain + ";" + cookie.Path + ";" + cookie.Name
	self.Set[key] = cookie
}

func (self *CookieSet) Marshal() string {
	header := make(map[string]interface{}, 0)
	header[self.Site] = self.Set
	body, _ := json.Marshal(header)
	return string(body)
}
