A simple utility to convert DataDog Dashboard to Terraform format. 

Requires `DATADOG_API_KEY` and `DATADOG_APP_KEY` environment variables.

Useful, if you had all dashboards configured adhoc and now want to follow DevOps style :)

# How to build
Just run (GOPATH and sometimes GOBIN have to be set):
```bash
go get && go build
```

# Examples
Export all dashboards:
```bash
DATADOG_API_KEY=xxx DATADOG_APP_KEY=xxx ./dd2tf
```

Export one particular dashboard (where `1111` is the ID of the dashboard):
```bash
DATADOG_API_KEY=xxx DATADOG_APP_KEY=xxx ./dd2tf -dashboards 1111
```

Write dashboards to corresponding files:
```bash
DATADOG_API_KEY=xxx DATADOG_APP_KEY=xxx ./dd2tf -files
```

You can find api/app keys in settings, under `Integrations -> API` section.
