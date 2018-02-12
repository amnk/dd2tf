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

func getAllDashboards(client datadog.Client) []int {
    var ids []int
    dashboards, err := client.GetDashboards()
    if err != nil {
        log.Fatal(err)
    }
    for _, elem := range dashboards {
        ids = append(ids, *elem.Id)
    }
    return ids
}

func getAllMonitors(client datadog.Client) []int {
    var ids []int
    monitors, err := client.GetMonitors()
    if err != nil {
        log.Fatal(err)
    }
    for _, elem := range monitors {
        ids = append(ids, *elem.Id)
    }
    return ids
}

func main() {
    var dashboards = flag.String("dashboards", "-1",
        "IDs of dashboards, separated by comma. If nothing is given, all dashboards are parsed")
    var monitors = flag.String("monitors", "-1",
        "IDs of monitors, separated by comma. If nothing is given, all monitors are exported")
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

    client := datadog.NewClient(datadog_api_key, datadog_app_key)

    //Keeps a list of dashboard IDs as ints
    var dash_ids []int
    if ! (*dashboards == "-1") {
        if len(*dashboards) == 0 {
            dash_ids = getAllDashboards(*client)
        } else {
            for _, element := range strings.Split(*dashboards, ",") {
                dash, _ := strconv.Atoi(element)
                dash_ids = append(dash_ids, dash)
            }
        }
    }

    var mon_ids []int
    if ! (*monitors == "-1") {
        if len(*monitors) == 0 {
            mon_ids = getAllMonitors(*client)
        } else {
            for _, element := range strings.Split(*monitors, ",") {
                mon, _ := strconv.Atoi(element)
                mon_ids = append(mon_ids, mon)
            }
        }
    }

    for _, element := range dash_ids {
        dash, err := client.GetDashboard(element)
        if err != nil {
            log.Fatal(err)
        }

        b, _ := Asset("tmpl/timeboard.tmpl")
        t, _ := template.New("").Parse(string(b))

        if *files {
            file := fmt.Sprintf("dash-%v.tf", element)
            f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
            if err != nil {
		log.Fatal(err)
	    }
            out := bufio.NewWriter(f)
            t.Execute(out, *dash)
            out.Flush()
            if err := f.Close(); err != nil {
                log.Fatal(err)
            }
        } else {
            t.Execute(os.Stdout, *dash)
        }
    }

    //TODO: those two loops should probably be refactored into some common loop.
    for _, element := range mon_ids {
        dash, err := client.GetMonitor(element)
        if err != nil {
            log.Fatal(err)
        }

        b, _ := Asset("tmpl/monitor.tmpl")
        t, _ := template.New("").Parse(string(b))

        if *files {
            file := fmt.Sprintf("mon-%v.tf", element)
            f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
            if err != nil {
		log.Fatal(err)
	    }
            out := bufio.NewWriter(f)
            t.Execute(out, *dash)
            out.Flush()
            if err := f.Close(); err != nil {
                log.Fatal(err)
            }
        } else {
            t.Execute(os.Stdout, *dash)
        }
    }

}
