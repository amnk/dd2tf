package main

import (
    "text/template"
    "flag"
    "log"
    "os"
    "strconv"
    "strings"
    "github.com/zorkian/go-datadog-api"
)

func getAllDashboards(client datadog.Client) []int {
    var ids []int
    dash_ids, _ := client.GetDashboards()
    for _, elem := range dash_ids {
        ids = append(ids, *elem.Id)
    }
    return ids
}

func main() {
    datadog_api_key, ok := os.LookupEnv("DATADOG_API_KEY")
    if !ok {
        log.Fatalf("Datadog API key not found, please make sure that DATADOG_API_KEY env variable is set")
    }

    datadog_app_key, ok := os.LookupEnv("DATADOG_APP_KEY")
    if !ok {
        log.Fatalf("Datadog APP key not found, please make sure that DATADOG_APP_KEY env variable is set")
    }

    var dashboards = flag.String("dashboards", "",
        "IDs of dashboards, separated by comma. If nothing is given, all dashboards are parsed")
    flag.Parse()

    client := datadog.NewClient(datadog_api_key, datadog_app_key)

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
            log.Fatalf("fatal: %s\n", err)
        }

        t := template.New("timeboard.tmpl").Funcs(template.FuncMap{"StringsJoin": strings.Join})
        t, _ = t.ParseFiles("timeboard.tmpl")
        t.Execute(os.Stdout, *dash)
    }
}
