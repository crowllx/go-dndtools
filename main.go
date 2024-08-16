package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func config() ([]byte, string) {
	if len(os.Args) < 3 {
		fmt.Println(`usage:
            <command> <characters.json> <log output directory`)
		os.Exit(1)
	}
	configPath := os.Args[1]
	outputDir := os.Args[2]
	characterData, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(outputDir)
	if err != nil {
		panic(err)
	}
	return characterData, outputDir
}
func main() {
	characterData, _ := config()
	clist := &clist{}
	if err := json.Unmarshal(characterData, clist); err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", clist)

	fmt.Printf("%d\n", len(clist.Characters))

	p := tea.NewProgram(initialModel(clist.Characters), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("some error: %v", err)
		os.Exit(1)
	}
}
