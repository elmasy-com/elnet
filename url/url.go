package url

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Query struct {
	Field string
	Value string
}

type UserInfo struct {
	Username string
	Password string
}

// Authority component of the URL.
//
//	[userinfo@]host[:port]
//
// If Port is -1, then port is not faound in the string. This solves the confusion if port is 0 (eg.: localhost:0).
type Authority struct {
	UserInfo *UserInfo
	Host     string
	Port     int
}

type Components struct {
	Scheme    string
	Authority Authority
	Path      string
	Queries   []Query
	Fragment  string
}

var (
	ErrURLEmpty          = errors.New("URL is empty")
	ErrURLInvalid        = errors.New("URL invalid")
	ErrURLNotAbsolute    = errors.New("URL is not absolute")
	ErrSchemeInvalid     = errors.New("scheme invalid")
	ErrSchemeMissing     = errors.New("scheme missing")
	ErrUserInfoEmpty     = errors.New("user info is empty")
	ErrUserInfoNameEmpty = errors.New("user info username is empty")
	ErrHostEmpty         = errors.New("host is empty")
	ErrPortInvalid       = errors.New("port invalid")
	ErrPortEmpty         = errors.New("port is empty")
	ErrQueryInvalid      = errors.New("query is invalid")
	ErrQueryEmpty        = errors.New("query is empty")
	ErrQueryFieldEmpty   = errors.New("query field is empty")
	ErrAuthortyMissing   = errors.New("authority missing")
)

func (q Query) String() string {
	return fmt.Sprintf("%s=%s", q.Field, q.Value)
}

func (ui UserInfo) String() string {

	return fmt.Sprintf("%s:%s", ui.Username, ui.Password)
}

func (a Authority) String() string {

	v := ""

	if a.UserInfo != nil {
		v += fmt.Sprintf("%s@", a.UserInfo)
	}

	v += a.Host

	if a.Port != -1 {
		v += fmt.Sprintf(":%d", a.Port)
	}

	return v
}

func (c Components) String() string {

	v := fmt.Sprintf("%s://%s", c.Scheme, c.Authority)

	if c.Path == "" {
		v += "/"
	} else {
		v += c.Path
	}

	if len(c.Queries) > 0 {
		v += "?"
		for i := range c.Queries {
			v += c.Queries[i].String()
			if i != len(c.Queries)-1 {
				v += "&"
			}
		}
	}

	if c.Fragment != "" {
		v += c.Fragment
	}

	return v
}

// IsValidScheme returns whether scheme is valid.
// This function just check the format, not compare this to known schemes.
func IsValidScheme(scheme string) bool {

	if len(scheme) == 0 {
		return false
	}

	for _, c := range scheme {
		switch {
		case c >= 'a' && c <= 'z':
			// a-z, lowercase characters
			continue
		case c == '-':
			continue
		case c == '.':
			continue
		case c == '+':
			continue
		default:
			return false
		}
	}

	return true
}

// Parse the UserInfo part of the URL.
func parseUserInfo(v string) (*UserInfo, error) {

	if v == "" {
		return nil, ErrUserInfoEmpty
	}

	ui := new(UserInfo)

	// 58 is Unicode ":"
	sep := strings.IndexRune(v, 58)
	if sep == -1 {
		ui.Username = v
		return ui, nil
	}

	// Username is an empty string (eg.: ":")
	if v[:sep] == "" {
		return nil, ErrUserInfoNameEmpty
	}

	ui.Username = v[:sep]

	// ":" is the last character, password is an empty string
	if len(v)-1 == sep {
		return ui, nil
	}

	ui.Password = v[sep+1:]

	return ui, nil
}

// Parse the authority section and validate the fields (eg.: not empty and port is in range (0-65535)).
// Returns *Authority and the remaining part of the URL.
// This function does not check the remaining part, only the scheme.
func parseAuthority(v string) (*Authority, error) {

	if v == "" {
		return nil, ErrAuthortyMissing
	}

	a := &Authority{}

	// Get UserInfo
	// 64 is Unicode "@"
	if sep := strings.IndexRune(v, 64); sep != -1 {

		ui, err := parseUserInfo(v[:sep])
		if err != nil {
			return nil, err
		}

		a.UserInfo = ui

		// "@" is the last char, host is missing
		if len(v)-1 == sep {
			return nil, ErrHostEmpty
		}

		v = v[sep+1:]
	}

	// Get Port
	// 58 is Unicode ":"
	if sep := strings.IndexRune(v, 58); sep != -1 {

		// ":" is the last char, host is missing
		if len(v)-1 == sep {
			return nil, ErrPortEmpty
		}

		port, err := strconv.Atoi(v[sep+1:])
		if err != nil {
			return nil, ErrPortInvalid
		}

		if port < 0 || port > 65535 {
			return nil, ErrPortInvalid
		}

		a.Port = port

		v = v[:sep]

	} else {
		a.Port = -1
	}

	a.Host = v

	return a, nil
}

// Parse the query part of the URL and return the queries.
// See more: https://en.wikipedia.org/wiki/Query_string
func parseQueries(v string) ([]Query, error) {

	queries := make([]Query, 0, 1)

	parts := strings.Split(v, "&")

	for i := range parts {

		qParts := strings.Split(parts[i], "=")

		// Format is not "f=v"
		if len(qParts) != 2 {
			return nil, ErrQueryInvalid
		}
		if qParts[0] == "" {
			return nil, ErrQueryFieldEmpty
		}

		queries = append(queries, Query{Field: qParts[0], Value: qParts[1]})
	}

	return queries, nil
}

// Parse disassemble an absolute URL (starting with the scheme).
//
//	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
func Parse(url string) (*Components, error) {

	if url == "" {
		return nil, ErrURLEmpty
	}

	comp := new(Components)

	//////////
	// Parse Scheme
	//////////

	// 58 is Unicode ':'
	sep := strings.IndexRune(url, 58)
	if sep == -1 {
		return nil, ErrSchemeMissing
	}

	if !IsValidScheme(url[:sep]) {
		return nil, ErrSchemeInvalid
	}

	comp.Scheme = url[:sep]

	// The url withous scheme should start with "://".
	// This means absolute URL.
	if !strings.HasPrefix(url[sep:], "://") {
		return nil, ErrURLNotAbsolute
	}

	url = strings.TrimPrefix(url[sep:], "://")

	//////////
	// Parse fragment
	//////////

	// 35 is Unicode '#'
	sep = strings.IndexRune(url, 35)
	if sep != -1 {

		// Include '#' in the fragment
		comp.Fragment = url[sep:]
		url = url[:sep]

	}

	//////////
	// Parse Queries
	//////////

	// 63 is Unicode '?'
	sep = strings.IndexRune(url, 63)
	if sep != -1 {

		// '?' is the last char
		if len(url)-1 == sep {
			return nil, ErrQueryEmpty
		} else {
			queries, err := parseQueries(url[sep+1:])
			if err != nil {
				return nil, err
			}
			comp.Queries = queries
			url = url[:sep]
		}
	}

	//////////
	// Parse Path
	//////////

	// 47 is Unicode '/'
	sep = strings.IndexRune(url, 47)
	if sep != -1 {
		comp.Path = url[sep:]
		url = url[:sep]
	}

	//////////
	// Parse Authority
	//////////

	// 58 is Unicode ':'

	auth, err := parseAuthority(url)
	if err != nil {
		return nil, err
	}

	comp.Authority = *auth

	return comp, nil
}
