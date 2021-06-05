package pkg

import (
	"regexp"
)

//ParseHereDocs function parses a byte array for a bash heredoc structures and returns a map
//of token to heredoc contents
func ParseHereDocs(content []byte) map[string][]byte {
	hdStart := regexp.MustCompile(`<<(-?)(\w+)\n`)
	indexes := hdStart.FindAllSubmatchIndex(content, -1)
	matches := make(map[string][]byte)
	for _, m := range indexes {
		if len(m) != 6 {
			continue
		}
		indentedHereDoc := m[3]-m[2] == 1 && content[m[2]:m[3]][0] == '-'
		token := string(content[m[4]:m[5]])
		rePattern := "(?s)" + token + `\n(.*)\n`
		if indentedHereDoc {
			rePattern += `\s*`
		}
		rePattern += token
		hdRe := regexp.MustCompile(rePattern)
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

//ParseHereDoc parses content and return contained data of the first heredoc statement
func ParseHereDoc(content []byte) []byte {
	hereDocs := ParseHereDocs(content)
	for _, v := range hereDocs {
		return v
	}
	return nil
}

//ParseHereDocString just like ParseHereDoc parses string content and return contained string of the first heredoc statement
func ParseHereDocString(content string) string {
	hereDoc := ParseHereDoc([]byte(content))
	if hereDoc == nil {
		return ""
	}
	return string(hereDoc)
}
