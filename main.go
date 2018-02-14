//go:generate go-bindata -o tpl.go tmpl

package main

import (
    "text/template"
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
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
    items []Item
    files bool
}

type DatadogElement interface {
    getElement(client datadog.Client, i int) (interface{}, error)
    getAsset() string
    getName() string
}

type Item struct {
    id int
    d DatadogElement
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
    t, _ := template.New("").Parse(string(b))

    if config.files {
        file := fmt.Sprintf("%v-%v.tf", i.d.getName(), item)
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

func main() {
    var dashboards = flag.String("dashboards", "-1",
        "IDs of dashboards, separated by comma")
    var all_dashboards = flag.Bool("all_dashboards", false,
        "If set, all dashobards will be exported.")
    var monitors = flag.String("monitors", "-1",
        "IDs of monitors, separated by comma")
    var all_monitors = flag.Bool("all_monitors", false,
        "If set, all monitors will be exported.")
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
        files: *files,
    }

    if (*all_dashboards) && !(*dashboards == "-1") {
        log.Fatalf("Either dashboards or all_dashboards can be used.")
    } else if (*all_dashboards) {
        config.items = append(config.items, getAllDashboards(config.client)...)
    } else if ! (*dashboards == "-1") {
        for _, element := range strings.Split(*dashboards, ",") {
            dash, _ := strconv.Atoi(element)
            config.items = append(config.items, Item{id: dash, d: Dashboard{}})
        }

    }

    if (*all_monitors) && !(*monitors == "-1") {
        log.Fatalf("Either monitors or all_monitors can be used.")
    } else if (*all_monitors) {
        config.items = append(config.items, getAllMonitors(config.client)...)
    } else if ! (*monitors == "-1") {
        for _, element := range strings.Split(*monitors, ",") {
            mon, _ := strconv.Atoi(element)
            config.items = append(config.items, Item{id: mon, d: Monitor{}})
        }
    }

    for _, element := range config.items {
        element.renderElement(config)
    }
}
