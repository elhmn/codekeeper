package cmd

import (
	"fmt"
	"strings"

	"github.com/elhmn/ckp/internal/config"
	"github.com/elhmn/ckp/internal/store"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

//NewFindCommand display a prompt for you to enter the code or solution you are looking for
func NewFindCommand(conf config.Config) *cobra.Command {
	command := &cobra.Command{
		Use:     "find",
		Aliases: []string{"f"},
		Short:   "find your code snippets and solutions",
		Long: `find your code snippets and solutions

	example: ckp find
	Will display a prompt for you to enter the code or solution you are looking for
	`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := findCommand(cmd, args, conf); err != nil {
				fmt.Fprintf(conf.OutWriter, "Error: %s\n", err)
				return
			}
		},
	}

	return command
}

func getTemplates() *promptui.SelectTemplates {
	//if you find a hard time understand it check out golang templating format documentation
	//here https://golang.org/pkg/text/template
	return &promptui.SelectTemplates{
		Label: "{{ if .Code.Content -}} {{`code:` | bold | green}} " +
			"{{.Code.Content}} {{- else -}} {{ .Solution.Content }} {{ end }}",
		Active: "* {{ if .Code.Content -}} {{`code:` | bold | green}} {{.Code.Content | bold}} {{ else }} " +
			"{{`solution:` | bold | yellow }} {{ .Solution.Content | bold }} {{ end }}",
		Inactive: "{{ if .Code.Content -}} {{`code:` | green}} {{.Code.Content}} " +
			"{{- else -}} {{`solution:` | yellow}} {{ .Solution.Content }} {{ end }}",
		Selected: " {{ `✓` | green }} {{if .Code.Content -}} {{ .Code.Content | bold }} {{- else -}} {{ .Solution.Content | bold }} {{ end }}",
		Details: "Type: {{- if .Code.Content }} code {{ else }} solution {{- end }}" +
			"{{ if .Code.Alias }} | Alias: {{ .Code.Alias }} {{- end }}" +
			"{{ if .Comment }} | Comment: {{ .Comment }} {{- end }}",
	}
}

func doesScriptContain(script store.Script, input string) bool {
	code := strings.Replace(strings.ToLower(script.Code.Content), " ", "", -1)
	solution := strings.Replace(strings.ToLower(script.Solution.Content), " ", "", -1)
	comment := strings.Replace(strings.ToLower(script.Comment), " ", "", -1)
	alias := strings.Replace(strings.ToLower(script.Code.Alias), " ", "", -1)
	content := fmt.Sprintf("%s %s %s %s", code, solution, comment, alias)
	input = strings.Replace(strings.ToLower(input), " ", "", -1)
	return strings.Contains(content, input)
}

func findCommand(cmd *cobra.Command, args []string, conf config.Config) error {
	//get store data
	storeFile, err := config.GetStoreFilePath(conf)
	if err != nil {
		return fmt.Errorf("failed to get the store file path: %s", err)
	}

	storeData, _, err := store.LoadStore(storeFile)
	if err != nil {
		return fmt.Errorf("failed to laod store: %s", err)
	}

	scripts := storeData.Scripts
	searchScript := func(input string, index int) bool {
		s := scripts[index]
		return doesScriptContain(s, input)
	}

	prompt := promptui.Select{
		Label:             "Enter your search text",
		Items:             scripts,
		Size:              5,
		StartInSearchMode: true,
		Searcher:          searchScript,
		Templates:         getTemplates(),
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed %v", err)
	}

	fmt.Fprintf(conf.OutWriter, "\n%s", result)
	return nil
}
