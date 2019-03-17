package mysql

import (
	"github.com/c-bata/go-prompt"
)

func Completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "SELECT"},
		{Text: "FROM"},
		{Text: "WHERE"},
		{Text: "GROUP BY"},
		{Text: "ORDER BY"},
		{Text: "COUNT(*)"},
		{Text: "MAX"},
		{Text: "SUM"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}
