package main

import (
        "html/template"
        "log"
        "os"
        "strconv"
        "github.com/zorkian/go-datadog-api"
)

func main() {
    datadog_api_key, ok := os.LookupEnv("DATADOG_API_KEY")
    if !ok {
        log.Printf("Datadog API key not found, please make sure that DATADOG_API_KEY env variable is set")
        os.Exit(1)
    }

    datadog_app_key, ok := os.LookupEnv("DATADOG_APP_KEY")
    if !ok {
        log.Printf("Datadog APP key not found, please make sure that DATADOG_APP_KEY env variable is set")
        os.Exit(1)
    }

    datadog_dashboard, ok := os.LookupEnv("DATADOG_DASH_ID")
    if !ok {
        log.Printf("Datadog dashboard id not found, please make sure that DATADOG_DASH_ID env variable is set")
        os.Exit(1)
    }

    client := datadog.NewClient(datadog_api_key, datadog_app_key)

    dash_id, err := strconv.Atoi(datadog_dashboard)

    dash, err := client.GetDashboard(dash_id)
    if err != nil {
        log.Fatalf("fatal: %s\n", err)
    }

    t := template.New("timeboard.tmpl")
    t, _ = t.ParseFiles("timeboard.tmpl")
    t.Execute(os.Stdout, *dash)
}
