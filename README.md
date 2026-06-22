# go-clean-architecture-starter

This repository is a Kikplate template for generating a Go HTTP API starter with Chi PostgreSQL layered clean architecture health checks graceful shutdown and Docker deployment defaults.

The template is defined by plate.yaml values.yaml and the files inside templates with tmpl extensions. You can customize the project name module path Go version API port database settings runtime timeouts and optional modules through the schema in plate.yaml.

## Build locally from this template

Run this command from the repository root.

```bash
kik generate --template . -f values.yaml --output-dir ./generated-project
```

This generates a project in generated-project using the local template files.

## Generate from Kikplate server

Run this command when using the published template name.

```bash
kik generate go-clean-architecture-starter -f values.yaml
```

This generates the same project shape using the remote template registry source.

## How to use the generated project

Change into the generated project directory then start the stack with Docker Compose or run the Go service directly after exporting the environment variables.

```bash
cd generated-project
docker compose up --build
```
