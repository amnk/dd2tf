[![Build Status](https://travis-ci.org/amnk/dd2tf.svg?branch=master)](https://travis-ci.org/amnk/dd2tf)

A simple utility to convert DataDog dashboards and/or monitors to Terraform format. 

Requires `DATADOG_API_KEY` and `DATADOG_APP_KEY` environment variables.

Useful, if you had all dashboards configured adhoc and now want to follow DevOps style :)

If you are using non-US datadog versions (e.g. `eu`), you can change the API URL using `DATADOG_BASE_URL`.

# How to build
Just run (GOPATH and sometimes GOBIN have to be set):
```bash
dep ensure
go generate && go build
```

# Examples
Export all dashboards:
```bash
DATADOG_API_KEY=xxx DATADOG_APP_KEY=xxx ./dd2tf dashboards --all
```

Export one particular dashboard (where `1111` is the ID of the dashboard):
```bash
DATADOG_API_KEY=xxx DATADOG_APP_KEY=xxx ./dd2tf dashboards --ids 111
```

Export one particular dashboard in the EU DataDog region (where `1111` is the ID of the dashboard):
```bash
DATADOG_BASE_URL="https://api.datadoghq.eu" DATADOG_API_KEY=xxx DATADOG_APP_KEY=xxx ./dd2tf dashboards --ids 111
```

Write dashboards to corresponding files:
```bash
DATADOG_API_KEY=xxx DATADOG_APP_KEY=xxx ./dd2tf dashboards --files --all
```

Datadog monitor can be exported with this command:
```bash
DATADOG_API_KEY=xxx DATADOG_APP_KEY=xxx ./dd2tf monitors --ids 1706011
```

And Datadog Screenboard:
```bash
DATADOG_API_KEY=xxx DATADOG_APP_KEY=xxx ./dd2tf screenboards --all
```

You can find api/app keys in settings, under `Integrations -> API` section.
