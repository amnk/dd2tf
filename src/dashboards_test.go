//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"github.com/zorkian/go-datadog-api"
)

var exampleDashboard = datadog.Dashboard{
	Id:          datadog.Int(1),
	Title:       datadog.String("Redis Timeboard (created via Terraform)"),
	Description: datadog.String("created using the Datadog provider in Terraform"),
	ReadOnly:    datadog.Bool(true),
	Graphs: []datadog.Graph{
		{
			Title: datadog.String("Redis latency (ms)"),
			Definition: &datadog.GraphDefinition{
				Viz: datadog.String("timeseries"),
				Requests: []datadog.GraphDefinitionRequest{{
					Query: datadog.String("avg:redis.info.latency_ms{$host}"),
					Type:  datadog.String("bars"),
					ConditionalFormats: []datadog.DashboardConditionalFormat{{
						Comparator: datadog.String("<"),
						Palette:    datadog.String("red_on_white"),
					}},
				}},
			},
		},
	},
}

func ExampleDashboard() {
	config := LocalConfig{files: false}
	item := Item{id: 1, d: Dashboard{}}
	item.renderElement(exampleDashboard, config)

	// Unordered output:
	// resource "datadog_timeboard" "1" {
	//   title       = "Redis Timeboard (created via Terraform)"
	//   description = "created using the Datadog provider in Terraform"
	//   read_only   = true
	//   graph {
	//     title = "Redis latency (ms)"
	//     viz   = "timeseries"
	//     request {
	//       q    = "avg:redis.info.latency_ms{$host}"
	//       type = "bars"
	//       conditional_format {
	//         palette = "red_on_white"
	//         comparator = "<"
	//       }
	//     }
	//   }
	// }
}
