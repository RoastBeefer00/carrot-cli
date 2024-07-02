package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/lipgloss"
	"github.com/RoastBeefer00/recipes-cli/list"
    fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
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
    MarginLeft(1).
    Align(lipgloss.Center)
    
var timeStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color(text)).
    Align(lipgloss.Center).MarginLeft(1)

var listStyle = lipgloss.NewStyle().
    Bold(true).
    MarginTop(2).
    PaddingLeft(1).
    PaddingRight(1).
    MarginLeft(1).
    MarginBottom(1).
    Foreground(lipgloss.Color(crust)).
    Background(lipgloss.Color(green)).
    Align(lipgloss.Center)

var itemStyle = lipgloss.NewStyle().
    Width(80).PaddingTop(1).MarginLeft(2)

var numStyle = lipgloss.NewStyle().
    PaddingTop(1).MarginLeft(2)

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

    idx, err := fuzzyfinder.FindMulti(
        recipes,
        func(i int) string {
            return fmt.Sprintf("%s (%s)", recipes[i].Name, recipes[i].Time)
        },
        fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
            if i == -1 {
                return ""
            }
            ing := list.New(recipes[i].Ingredients).Enumerator(list.Bullet)
            s := list.New(recipes[i].Steps).Enumerator(list.Arabic).ItemStyle(itemStyle).EnumeratorStyle(numStyle)
            return fmt.Sprintf("%s \n%s\n\n%s\n%s\n\n%s\n%s",
                titleStyle.Render(recipes[i].Name),
                timeStyle.Render(recipes[i].Time),
                listStyle.Render("Ingredients"),
                ing.EnumeratorStyle(lipgloss.NewStyle().MarginLeft(1)).ItemStyle(lipgloss.NewStyle().MarginLeft(1)),
                listStyle.Render("Steps"),
                s,
                )
        }))
    if err != nil {
        panic(err)
    }

    for _, i := range idx {
        recipe := recipes[i]
        ing := list.New(recipe.Ingredients).Enumerator(list.Bullet)
        s := list.New(recipe.Steps).Enumerator(list.Arabic).ItemStyle(itemStyle).EnumeratorStyle(numStyle)

        fmt.Println(titleStyle.Render(recipe.Name))
        fmt.Println(timeStyle.Render(recipe.Time))
        fmt.Println(listStyle.Render("Ingredients"))
        fmt.Println(ing.EnumeratorStyle(lipgloss.NewStyle().MarginLeft(1)).ItemStyle(lipgloss.NewStyle().MarginLeft(1)))
        fmt.Println(listStyle.Render("Steps"))
        fmt.Println(s)

    }
}
