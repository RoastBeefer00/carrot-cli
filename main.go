package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

type Recipe struct {
	Name        string   `json:"Name"`
	Time        string   `json:"Time"`
	Ingredients []string `json:"Ingredients"`
	Steps       []string `json:"Steps"`
}

const rosewater string = "#f5e0dc"
const flamingo string = "#f2cdcd"
const pink string = "#f5c2e7"
const mauve string = "#cba6f7"
const red string = "#f38ba8"
const maroon string = "#eba0ac"
const peach string = "#fab387"
const yellow string = "#f9e2af"
const green string = "#a6e3a1"
const teal string = "#94e2d5"
const sky string = "#89dceb"
const sapphire string = "#74c7ec"
const blue string = "#89b4fa"
const lavender string = "#b4befe"
const text string = "#cdd6f4"
const subtext1 string = "#bac2de"
const subtext0 string = "#a6adc8"
const overlay2 string = "#9399b2"
const overlay1 string = "#7f849c"
const overlay0 string = "#6c7086"
const surface2 string = "#585b70"
const surface1 string = "#45475a"
const surface0 string = "#313244"
const base string = "#1e1e2e"
const mantle string = "#181825"
const crust string = "#11111b"

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color(crust)).
	Background(lipgloss.Color(sky)).
    PaddingLeft(2).
    PaddingRight(2).
    PaddingTop(1).
    PaddingBottom(1).
    Width(22).
    Align(lipgloss.Center)
    
var timeStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color(text)).
    Align(lipgloss.Center)

var recipes []Recipe

func main() {
	client := http.Client{}
	url := "https://r-j-magenta-carrot-42069.uc.r.appspot.com/recipes/all"

	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &recipes)
	if err != nil {
		panic(err)
	}

    recipe := recipes[0]
    l := list.New(recipe.Ingredients).Enumerator(list.Bullet)

    fmt.Println(titleStyle.Render(recipe.Name))
    fmt.Println(timeStyle.Render(recipe.Time))
    fmt.Println(l)
}
