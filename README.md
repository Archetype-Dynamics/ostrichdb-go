# ostrichdb-go

Go SDK for interacting with OstrichDB

## Prerequisites

Before using this SDK, you need:

- [Go](https://go.dev/doc/install) 1.23.1 or higher
- [The Odin Compiler](https://odin-lang.org/docs/install/)
- Python 3.x

You also need to have OstrichDB running:

1. Create a `.env` file in root of the project that you are using the SDK in
2. In that `.env` file create a new entry WITHOUT a value: `OSTRICHDB_JWT=`
3. Clone or fork the [Open-OstrichDB repository](https://github.com/archetype-dynamics/Open-OstrichDB)
*For step 3, It is recommended to fork the repo and put it on root of your working project*
  
4. Once the Open-OstrichDB repo is on your machine `cd` into it
*The following steps assume your CWD is the root of the Open-OstrichDB repo*
2. Obtain a JWT authentication token by running the JWT creation script: `python3 create_test_jwt.py`
3. Store the printed value in the `.env` file you created in step 2.
3. Start the OstrichDB server:`./scripts/local_run.sh`

## Installation

Import the SDK into your Go project:

```go
import "ostrichdb-go/src/sdk"
```

## Usage

### Basic Setup

Every interaction with OstrichDB requires initializing a config and client:

```go
package main

import (
    "fmt"
    "ostrichdb-go/src/sdk"
)

func main() {
    // Initialize configuration (loads JWT from .env)
    config := sdk.NewConfigBuilder()

    // Create a client with the configuration
    client := sdk.NewClientBuilder(config)
}
```

### Creating a Project

Projects are top-level containers in OstrichDB. If the project already exists, use its existing name:

```go
// Create or reference a project
project := sdk.NewProjectBuilder(client, "myCoolProject")
err := sdk.CreateProject(project)
if err != nil {
    fmt.Println("Error creating project:", err)
}
```

### Creating a Collection

Collections (databases) belong to projects:

```go
// Create a collection within the project
collection := sdk.NewCollectionBuilder(project, "myFirstCollection")
err := sdk.CreateCollection(collection)
if err != nil {
    fmt.Println("Error creating collection:", err)
}
```

### Working with Clusters and Records

Clusters are groups of related records within a collection:

```go
// Create a cluster
cluster := sdk.NewClusterBuilder(collection, "myCluster")
err := sdk.CreateCluster(cluster)
if err != nil {
    fmt.Println("Error creating cluster:", err)
}

// Create a record within the cluster
record := sdk.NewRecordBuilder(cluster, "recordName", "string", "myValue")
err = sdk.CreateRecord(record)
if err != nil {
    fmt.Println("Error creating record:", err)
}

// Delete a record
err = sdk.DeleteRecord(record)
if err != nil {
    fmt.Println("Error deleting record:", err)
}
```

## Configuration

The SDK requires a `.env` file in your project root containing your OstrichDB JWT token:

```
OSTRICHDB_JWT=your_jwt_token_here
```

## Example

See [main.go](./main.go) for a complete working example.