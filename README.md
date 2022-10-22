# go-freshservice

An unofficial Go client for the [FreshService](https://api.freshservice.com/) API.

## Usage

```go
import "github.com/santiago8781/go-freshservice/freshservice"
```

Simply create a new FreshService client, then use the various services on the client to access the different resource 
types on the FreshService API.

```go
ctx := context.Background()
fs, err := freshservice.NewClient(ctx, "company", "MY-API-TOKEN")
if err != nil {
log.Fatalf("Failed to create client: %v", err)
}
```

### Example

The below example aims to give a short introduction in how to use the client and services.

```go
package main

import (
    "context"
    "github.com/theapsgroup/go-freshservice/freshservice"
    "log"
)

func main() {
    ctx := context.Background()
    fs, err := freshservice.NewClient(ctx, "company", "MY-API-TOKEN")
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    // Obtain info for a user (Requester)
    requester, _, err := fs.Requesters.GetRequester(123)
    log.Printf("%s %s - %s\n", requester.FirstName, requester.LastName, requester.Email)
    
    // Obtain second page of Tickets for the Requester
    opt := freshservice.ListTicketsOptions{
        Email: &requester.Email,
        ListOptions: freshservice.ListOptions{
            Page: 2,
        },
    }
    
    tickets, _, err := fs.Tickets.ListTickets(&opt)
    for _, ticket := range tickets.Collection {
        log.Printf("Ticket: %d (%s)\n", ticket.ID, ticket.Subject)
    }
}
```
