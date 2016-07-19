package lordsofvegas

import (
	"strings"

	"github.com/aymerick/raymond"
	"github.com/brdgme/render"
)

var tpl = raymond.MustParse(`
Bacon {{#fg "blue"}}This {{#fg "red"}}should{{/fg}} be blue{{/fg}} cheese.

And {{#bg "green"}}here {{#bg "yellow"}}are{{/bg}} some{{/bg}} backgrounds.

Nested {{#b}}bold {{#b}}is{{/b}} cool{{/b}}, really.

{{#fg "red"}}Red{{/fg}}
{{#fg "pink"}}Pink{{/fg}}
{{#fg "purple"}}Purple{{/fg}}
{{#fg "deeppurple"}}Deep purple{{/fg}}
{{#fg "indigo"}}Indigo{{/fg}}
{{#fg "blue"}}Blue{{/fg}}
{{#fg "lightblue"}}Light blue{{/fg}}
{{#fg "cyan"}}Cyan{{/fg}}
{{#fg "teal"}}Teal{{/fg}}
{{#fg "green"}}Green{{/fg}}
{{#fg "lightgreen"}}Light green{{/fg}}
{{#fg "lime"}}Lime{{/fg}}
{{#fg "yellow"}}Yellow{{/fg}}
{{#fg "amber"}}Amber{{/fg}}
{{#fg "orange"}}Orange{{/fg}}
{{#fg "deeporange"}}Deep orange{{/fg}}
{{#fg "brown"}}Brown{{/fg}}
{{#fg "grey"}}Grey{{/fg}}
{{#fg "bluegrey"}}Blue grey{{/fg}}
{{#fg "white"}}White{{/fg}}
{{#fg "black"}}Black{{/fg}}
`)

var CasinoColors = map[int]string{
	CasinoAlbion:  render.Purple,
	CasinoSphynx:  render.Yellow,
	CasinoVega:    render.Green,
	CasinoTivoli:  render.Grey,
	CasinoPioneer: render.Red,
}

var renderStripLine = `  {{bg "black"}}  {{_bg}}  `
var renderStrip = strings.Join([]string{
	renderStripLine,
	renderStripLine,
	renderStripLine,
	renderStripLine,
}, "\n")

func (g *Game) Template() *raymond.Template {
	return tpl
}

// func (g *Game) RenderForPlayer(player int) (string, error) {
// 	cells := [][]interface{}{}
// 	for _, layoutRow := range BoardLayout {
// 		cellsRow := []interface{}{}
// 		for _, layoutCell := range layoutRow {
// 			var cell interface{}
// 			if layoutCell == "ST" {
// 				cell = renderStrip
// 			} else if bs, ok := BoardSpaceByLocation[layoutCell]; ok {
// 				cell = RenderSpace(bs, g.Board[layoutCell])
// 			} else {
// 				cell = "\n\n\n"
// 			}
// 			cellsRow = append(cellsRow, cell)
// 		}
// 		cells = append(cells, cellsRow)
// 	}
// 	return render.Table(cells, 0, 0), nil
// }
//
// var (
// 	renderSpaceBorderStr    = " "
// 	renderSpaceBorderWidth  = 2
// 	renderSpaceContentWidth = 5
// )
//
// func RenderCasinoBg(input string, casino int) string {
// 	if col, ok := CasinoColors[casino]; ok {
// 		return render.Bg(input, col)
// 	}
// 	return input
// }
//
// func RenderSpace(bs BoardSpace, bsState BoardSpaceState) string {
// 	edge := RenderCasinoBg(strings.Repeat(renderSpaceBorderStr, renderSpaceBorderWidth), bsState.Casino)
// 	header := []interface{}{
// 		edge,
// 		RenderCasinoBg(strings.Repeat(renderSpaceBorderStr, renderSpaceContentWidth), bsState.Casino),
// 		edge,
// 	}
// 	locText := bs.Location
// 	if bsState.Owned {
// 		locText = render.Markup(locText, render.PlayerColour(bsState.Owner), true)
// 	}
// 	cells := [][]interface{}{
// 		header,
// 		{edge, render.Centred(locText), edge},
// 	}
// 	contextualRow := ""
// 	if bsState.Casino == CasinoNone {
// 		contextualRow = fmt.Sprintf(
// 			"%s %d",
// 			RenderPrice(bs.BuildPrice),
// 			bs.Dice,
// 		)
// 	} else if bsState.Owned && bsState.Dice > 0 {
// 		contextualRow = render.Markup(
// 			fmt.Sprintf("%d", bsState.Dice),
// 			render.PlayerColour(bsState.Owner),
// 			true,
// 		)
// 	}
// 	cells = append(cells, []interface{}{
// 		edge,
// 		render.Centred(contextualRow),
// 		edge,
// 	})
// 	cells = append(cells, header)
// 	return render.Table(cells, 0, 0)
// }
//
// func RenderPrice(price int) string {
// 	return render.Markup(fmt.Sprintf("$%d", price), render.Yellow, true)
// }
