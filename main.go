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
    "github.com/zorkian/go-datadog-api"
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

func main() {
    var dashboards = flag.String("dashboards", "",
        "IDs of dashboards, separated by comma. If nothing is given, all dashboards are parsed")
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
    if len(*dashboards) == 0 {
        dash_ids = getAllDashboards(*client)
    } else {
        for _, element := range strings.Split(*dashboards, ",") {
            dash, _ := strconv.Atoi(element)
            dash_ids = append(dash_ids, dash)
        }
    }

    for _, element := range dash_ids {
        dash, err := client.GetDashboard(element)
        if err != nil {
            log.Fatal(err)
        }

        //TODO: keep template a part of the binary
        b, _ := Asset("tmpl/timeboard.tmpl")
        t, _ := template.New("").Parse(string(b))

        if *files {
            file := fmt.Sprintf("%v.tf", element)
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
