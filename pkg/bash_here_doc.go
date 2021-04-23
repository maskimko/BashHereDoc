package pkg

import (
	"regexp"
)

//ParseHereDocs function parses a byte array for a bash heredoc structures and returns a map
//of token to heredoc contents
func ParseHereDocs(content []byte) map[string][]byte {
	hdStart := regexp.MustCompile(`<<(\w+)\n`)
	indexes := hdStart.FindAllSubmatchIndex(content, -1)
	matches := make(map[string][]byte)
	for _, m := range indexes {
		if len(m) != 4 {
			continue
		}
		token := string(content[m[2]:m[3]])
		hdRe := regexp.MustCompile("(?s)" + token + `\n(.*)\n` + token)
		hdIndexes := hdRe.FindSubmatchIndex(content)
		if len(hdIndexes) != 4 {
			continue
		}
		matches[token] = content[hdIndexes[2]:hdIndexes[3]]
	}
	if len(matches) == 0 {
		return nil
	}
	return matches
}
