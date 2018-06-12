//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"

	flag "github.com/spf13/pflag"
	"gopkg.in/zorkian/go-datadog-api.v2"
)

func getAllDashboards(client datadog.Client) []Item {
	var ids []Item
	dashboards, err := client.GetDashboards()
	if err != nil {
		log.Fatal(err)
	}
	for _, elem := range dashboards {
		ids = append(ids, Item{id: *elem.Id, d: Dashboard{}})
	}
	return ids
}

func getAllMonitors(client datadog.Client) []Item {
	var ids []Item
	monitors, err := client.GetMonitors()
	if err != nil {
		log.Fatal(err)
	}
	for _, elem := range monitors {
		ids = append(ids, Item{id: *elem.Id, d: Monitor{}})
	}
	return ids
}

type LocalConfig struct {
	client datadog.Client
	items  []Item
	files  bool
}

type DatadogElement interface {
	getElement(client datadog.Client, i int) (interface{}, error)
	getAsset() string
	getName() string
}

type Item struct {
	id int
	d  DatadogElement
}

type Dashboard struct {
}

func (d Dashboard) getElement(client datadog.Client, id int) (interface{}, error) {
	dash, err := client.GetDashboard(id)
	return dash, err
}

func (d Dashboard) getAsset() string {
	return "tmpl/timeboard.tmpl"
}

func (d Dashboard) getName() string {
	return "dashboard"
}

type Monitor struct {
}

func (m Monitor) getElement(client datadog.Client, id int) (interface{}, error) {
	mon, err := client.GetMonitor(id)
	return mon, err
}

func (m Monitor) getAsset() string {
	return "tmpl/monitor.tmpl"
}

func (m Monitor) getName() string {
	return "monitor"
}

type RenderableElement interface {
	renderElement(config LocalConfig)
}

func (i *Item) renderElement(config LocalConfig) {
	item, err := i.d.getElement(config.client, i.id)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := Asset(i.d.getAsset())
	t, _ := template.New("").Funcs(template.FuncMap{
		"escapeCharacters": escapeCharacters,
	}).Parse(string(b))

	if config.files {
		file := fmt.Sprintf("%v-%v.tf", i.d.getName(), i.id)
		f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(err)
		}
		out := bufio.NewWriter(f)
		t.Execute(out, item)
		out.Flush()
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	} else {
		t.Execute(os.Stdout, item)
	}
}

// Replace escaped quote with apostrophe
func escapeCharacters(line string) string {
	return strconv.Quote(line)
}

func main() {
	var dashboards = flag.String("dashboards", "",
		"IDs of dashboards, separated by comma. If no IDs are given, all dashboards are exported.")
	flag.Lookup("dashboards").NoOptDefVal = "-1"
	var monitors = flag.String("monitors", "",
		"IDs of monitors, separated by comma. If no IDs are given, all monitors are exported.")
	flag.Lookup("monitors").NoOptDefVal = "-1"
	var files = flag.Bool("files", false, "Create file for each entity instead of stdout dump")

	flag.Parse()

	datadog_api_key, ok := os.LookupEnv("DATADOG_API_KEY")

	if !ok {
		log.Fatalf("Datadog API key not found, please make sure that DATADOG_API_KEY env variable is set")
	}

	datadog_app_key, ok := os.LookupEnv("DATADOG_APP_KEY")
	if !ok {
		log.Fatalf("Datadog APP key not found, please make sure that DATADOG_APP_KEY env variable is set")
	}

	config := LocalConfig{
		client: *datadog.NewClient(datadog_api_key, datadog_app_key),
		files:  *files,
	}

	if *dashboards == "-1" {
		config.items = append(config.items, getAllDashboards(config.client)...)
	} else if !(*dashboards == "") {
		for _, element := range strings.Split(*dashboards, ",") {
			dash, _ := strconv.Atoi(element)
			config.items = append(config.items, Item{id: dash, d: Dashboard{}})
		}

	}

	if *monitors == "-1" {
		config.items = append(config.items, getAllMonitors(config.client)...)
	} else if !(*monitors == "") {
		for _, element := range strings.Split(*monitors, ",") {
			mon, _ := strconv.Atoi(element)
			config.items = append(config.items, Item{id: mon, d: Monitor{}})
		}
	}

	for _, element := range config.items {
		element.renderElement(config)
	}
}
