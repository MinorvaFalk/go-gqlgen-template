Go-Gqlgen-Template
---
Template for developing Go Graphql API using Gqlgen

## TODO List 📃

- [ ] TODO #1
- [ ] TODO #2
- [ ] TODO #3


## Getting Started 🚀

Rover Command:
```bash
# Upload schema
rover subgraph publish go-gqlgen-template@current \
--name mock --schema ./graph/schema/schema.graphql \
--routing-url http://localhost:4001/query

# Introspection
rover subgraph publish go-gqlgen-template@current \
--name mock --schema - \
--routing-url http://localhost:4001/query
```

### Starting Application
```bash
make run
```