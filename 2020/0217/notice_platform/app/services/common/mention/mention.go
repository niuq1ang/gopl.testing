package mention

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"

	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/user"
	"gopkg.in/gorp.v1"
)

type Mention struct {
	Location int
	Text     string
}

func ReplaceUUIDsWithNames(src gorp.SqlExecutor, s string, mentions []*Mention) (result string, err error) {
	var uuidSet = make(map[string]bool)
	var valid []*Mention
	for _, m := range mentions {
		if len(m.Text) == 8 {
			uuidSet[m.Text] = true
			valid = append(valid, m)
		}
	}
	if len(valid) == 0 {
		result = s
		return
	}

	uuids := make([]string, len(uuidSet))
	i := 0
	for uuid, _ := range uuidSet {
		uuids[i] = uuid
		i++
	}
	names, err := user.MapNamesByUserUUIDs(src, uuids)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	var cursor int
	for _, m := range valid {
		if name, ok := names[m.Text]; ok {
			buf.WriteString(s[cursor:m.Location])
			buf.WriteString(name)
			cursor = m.Location + len(m.Text)
		}
	}
	if cursor < len(s) {
		buf.WriteString(s[cursor:])
	}
	result = buf.String()
	return
}

func Parse(s string) (mentions []*Mention) {
	reader := bufio.NewReader(strings.NewReader(s))
	var r, prev rune
	var err error
	var loc, size int
	start := -1
	for err == nil {
		r, size, err = reader.ReadRune()
		if start >= 0 {
			if err == io.EOF || unicode.IsSpace(r) {
				if loc != start+1 {
					m := &Mention{
						Location: start + 1,
						Text:     s[start+1 : loc],
					}
					mentions = append(mentions, m)
				}
				start = -1
			}
		} else if r == '@' {
			if prev <= ' ' || prev > '~' {
				start = loc
			}
		}
		prev = r
		loc += size
	}
	return
}
