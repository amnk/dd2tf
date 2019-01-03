//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"github.com/zorkian/go-datadog-api"
)

var exampleScreenboard = datadog.Screenboard{
	Id:     datadog.Int(1),
	Title:  datadog.String("Test"),
	Shared: datadog.Bool(false),
	Widgets: []datadog.Widget{
		{
			Type:       datadog.String("query_value"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			Legend:     datadog.Bool(true),
			LegendSize: datadog.String("16"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("1d"),
			},
			TileDef: &datadog.TileDef{
				Viz: datadog.String("query_value"),
				Requests: []datadog.TileDefRequest{{
					Query: datadog.String("avg:system.cpu.user{*}"),
					Type:  datadog.String("line"),
					Style: &datadog.TileDefRequestStyle{
						Palette: datadog.String("purple"),
						Type:    datadog.String("dashed"),
						Width:   datadog.String("thin"),
					},
					ConditionalFormats: []datadog.ConditionalFormat{
						{
							Comparator: datadog.String(">="),
							Value:      datadog.String("1"),
							Palette:    datadog.String("white_on_red"),
						}},
					Aggregator: datadog.String("max"),
				}},
				CustomUnit: datadog.String("%"),
				Autoscale:  datadog.Bool(false),
				TextAlign:  datadog.String("right"),
			},
			Logset: datadog.String("test"),
		},
	},
}

func Example() {
	config := LocalConfig{files: false}
	item := Item{id: 1, d: ScreenBoard{}}
	item.renderElement(exampleScreenboard, config)

	// Unordered output:
	// resource "datadog_screenboard" "1" {
	//   title = "Test"
	//   shared = false
	//   widget {
	//     type = "query_value"
	//     x = "1"
	//     y = "1"
	//     title = "true"
	//     title_text = "Test title"
	//     title_size = "16"
	//     height = "5"
	//     width = "5"
	//     title_align = "right"
	//     time {
	//       live_span = "1d"
	//     }
	//     tile_def {
	//       viz = "query_value"
	//       custom_unit = "%"
	//       autoscale = "false"
	//       text_align = "right"
	//       request {
	//         q = "avg:system.cpu.user{*}"
	//         type = "line"
	//         aggregator = "max"
	//         conditional_format {
	//           palette = "white_on_red"
	//           comparator = ">="
	//           value = "1"
	//         }
	//         style {
	//           palette = "purple"
	//           type = "dashed"
	//           width = "thin"
	//         } //style
	//       } //request
	//     } //tile_def
	//     legend = "true"
	//     legend_size = "16"
	//     logset = "test"
	//   } //widget
	// }
}
