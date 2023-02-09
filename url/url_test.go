package url

import (
	"errors"
	"testing"
)

func TestIsSchemeValid(t *testing.T) {

	cases := []struct {
		V string
		R bool
	}{
		{V: "http", R: true},
		{V: "https", R: true},
		{V: "gopher", R: true},
		{V: "iris.xpc", R: true},
		{V: "ms-help", R: true},
		{V: "HTTP", R: false},
		{V: "http s", R: false},
		{V: "http/s", R: false},
	}

	for i := range cases {

		if r := IsValidScheme(cases[i].V); r != cases[i].R {
			t.Fatalf("FAIL: %s is %v", cases[i].V, r)
		}
	}
}

func BenchmarkIsSchemeValid(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsValidScheme("https")
	}
}

func TestParseUserInfo(t *testing.T) {

	cases := []struct {
		V    string
		Err  error
		User string
		Pass string
	}{
		{V: "", Err: ErrUserInfoEmpty, User: "", Pass: ""},
		{V: ":", Err: ErrUserInfoNameEmpty, User: "", Pass: ""},
		{V: "user", Err: nil, User: "user", Pass: ""},
		{V: "user:", Err: nil, User: "user", Pass: ""},
	}

	for i := range cases {

		ui, err := parseUserInfo(cases[i].V)

		if !errors.Is(err, cases[i].Err) {
			t.Fatalf("Invalid error for %s: want=\"%s\", got=\"%s\"", cases[i].V, cases[i].Err, err)
		}

		if err != nil {
			continue
		}

		if ui.Username != cases[i].User {
			t.Fatalf("Invalid Username for %s: want=\"%s\", got=\"%s\"", cases[i].V, cases[i].User, ui.Username)
		}
		if ui.Password != cases[i].Pass {
			t.Fatalf("Invalid Password for %s: want=\"%s\", got=\"%s\"", cases[i].V, cases[i].Pass, ui.Password)
		}
	}
}

func BenchmarkParseUserInfo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		parseUserInfo("user:pass")
	}

}

func TestParseAuthority(t *testing.T) {

	cases := []struct {
		V   string
		Err error
		UI  string
		H   string // Host
		P   int    // Port
	}{
		{V: "test", Err: nil, H: "test", P: -1},
		{V: "example.com", Err: nil, H: "example.com", P: -1},
		{V: "example.com:80", Err: nil, H: "example.com", P: 80},
		{V: "user@example.com:80", Err: nil, UI: "user:", H: "example.com", P: 80},
		{V: "user:@example.com:80", Err: nil, UI: "user:", H: "example.com", P: 80},
		{V: "user:pass@example.com:80", Err: nil, UI: "user:pass", H: "example.com", P: 80},

		{V: "", Err: ErrAuthortyMissing},
		{V: "@", Err: ErrUserInfoEmpty},
		{V: "test@", Err: ErrHostEmpty},
		{V: "@example.com:80", Err: ErrUserInfoEmpty, H: "", P: 0},
		{V: ":pass@example.com:80", Err: ErrUserInfoNameEmpty, H: "", P: 0},
		{V: "example.com:", Err: ErrPortEmpty, H: "", P: 0},
		{V: "example.com:655350", Err: ErrPortInvalid, H: "", P: 0},
	}

	for i := range cases {

		a, err := parseAuthority(cases[i].V)
		if !errors.Is(err, cases[i].Err) {
			t.Fatalf("FAIL: Invalid error for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].Err, err)
		}

		if err != nil {
			continue
		}

		if cases[i].UI == "" && a.UserInfo != nil {
			t.Fatalf("FAIL: UserInfo for %s should be nil", cases[i].V)
		}
		if a.UserInfo == nil && cases[i].UI != "" {
			t.Fatalf("FAIL: UserInfo for %s should NOT be nil", cases[i].V)

		}

		if a.UserInfo != nil {
			if a.UserInfo.String() != cases[i].UI {
				t.Fatalf("FAIL: Invalid UserInfo for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].UI, a.UserInfo.String())
			}
		}

		if a.Host != cases[i].H {
			t.Fatalf("FAIL: Invalid Host for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].H, a.Host)
		}

		if a.Port != cases[i].P {
			t.Fatalf("FAIL: Invalid Port for %s: want=%d get=%d", cases[i].V, cases[i].P, a.Port)
		}
	}
}

func BenchmarkParseAuthority(b *testing.B) {

	for i := 0; i < b.N; i++ {
		parseAuthority("user:pass@example.com:80")
	}
}

func TestParseQueries(t *testing.T) {

	cases := []struct {
		V   string // Test string
		Err error
		N   int    // Num of queries
		F0  string // Queries[0].Field
		V0  string // Queries[0].Value
		F1  string
		V1  string
		F2  string
		V2  string
		P   int
	}{
		{V: "f0=v0", Err: nil, N: 1, F0: "f0", V0: "v0"},
		{V: "f0=v0&f1=v1", Err: nil, N: 2, F0: "f0", V0: "v0", F1: "f1", V1: "v1"},
		{V: "f0=v0&f1=v1&f2=v2", Err: nil, N: 3, F0: "f0", V0: "v0", F1: "f1", V1: "v1", F2: "f2", V2: "v2"},

		{V: "f0=", Err: nil, N: 1, F0: "f0", V0: ""},
		{V: "=v0", Err: ErrQueryFieldEmpty},
		{V: "f0", Err: ErrQueryInvalid},
	}

	for i := range cases {

		queries, err := parseQueries(cases[i].V)
		if !errors.Is(err, cases[i].Err) {
			t.Fatalf("FAIL: Invalid error for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].Err, err)
		}

		if err != nil {
			continue
		}

		if len(queries) != cases[i].N {
			t.Fatalf("FAIL: Invalid length for %s: want=%d get=%d", cases[i].V, cases[i].N, len(queries))
		}

		// First element
		if queries[0].Field != cases[i].F0 {
			t.Fatalf("FAIL: Invalid Field[0] for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].F0, queries[0].Field)
		}
		if queries[0].Value != cases[i].V0 {
			t.Fatalf("FAIL: Invalid Value[0] for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].V0, queries[0].Value)
		}

		// Second element
		if len(queries) < 2 {
			continue
		}

		if queries[1].Field != cases[i].F1 {
			t.Fatalf("FAIL: Invalid Field[1] for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].F1, queries[1].Field)
		}
		if queries[1].Value != cases[i].V1 {
			t.Fatalf("FAIL: Invalid Value[1] for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].V1, queries[1].Value)
		}

		// Third element
		if len(queries) < 3 {
			continue
		}

		if queries[2].Field != cases[i].F2 {
			t.Fatalf("FAIL: Invalid Field[2] for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].F1, queries[2].Field)
		}
		if queries[2].Value != cases[i].V2 {
			t.Fatalf("FAIL: Invalid Value[2] for %s: want=\"%s\" get=\"%s\"", cases[i].V, cases[i].V1, queries[2].Value)
		}

	}
}

func BenchmarkParseQueries(b *testing.B) {

	for i := 0; i < b.N; i++ {
		parseQueries("f0=v0&f1=v1")
	}
}

func TestParse(t *testing.T) {

	cases := []struct {
		V   string
		Err error
		R   string
	}{
		{V: "http://example.com", R: "http://example.com/"},
		{V: "http://example.com:80", R: "http://example.com:80/"},
		{V: "http://user@example.com:80", R: "http://user:@example.com:80/"},
		{V: "http://user:pass@example.com:80", R: "http://user:pass@example.com:80/"},
		{V: "http://user:pass@example.com:80/path/to/file", R: "http://user:pass@example.com:80/path/to/file"},
		{V: "http://user:pass@example.com:80/path/to/file?field0=value0&field1=value1", R: "http://user:pass@example.com:80/path/to/file?field0=value0&field1=value1"},
		{V: "http://user:pass@example.com:80/path/to/file?field0=value0&filed1=value1#header", R: "http://user:pass@example.com:80/path/to/file?field0=value0&filed1=value1#header"},

		{V: "", Err: ErrURLEmpty},
		{V: "example.com", Err: ErrSchemeMissing},
		{V: "http/s://example.com", Err: ErrSchemeInvalid},
		{V: "https:example.com", Err: ErrURLNotAbsolute},
		{V: "https://example.com/?", Err: ErrQueryEmpty},
		{V: "https://example.com/?field", Err: ErrQueryInvalid},
		{V: "https://example.com/?=value", Err: ErrQueryFieldEmpty},
		{V: "https:///?field=value", Err: ErrAuthortyMissing},
		{V: "https://@example.com/", Err: ErrUserInfoEmpty},
		{V: "https://:@example.com/", Err: ErrUserInfoNameEmpty},
		{V: "https://example.com:/", Err: ErrPortEmpty},
		{V: "https://example.com:str/", Err: ErrPortInvalid},
		{V: "https://example.com:655350/", Err: ErrPortInvalid},
	}

	for i := range cases {
		comp, err := Parse(cases[i].V)
		if !errors.Is(err, cases[i].Err) {
			t.Fatalf("FAIL: Invalid error for %s: want=\"%s\", got=\"%s\"", cases[i].V, cases[i].Err, err)
		}

		if err != nil {
			continue
		}

		if comp.String() != cases[i].R {
			t.Fatalf("FAIL: Invalid result for %s: want=\"%s\", got=\"%s\"", cases[i].V, cases[i].R, comp)
		}
	}
}

func BenchmarkParse(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Parse("http://user:pass@example.com:80/path/to/file?field0=value0&filed1=value1#header")
	}
}
