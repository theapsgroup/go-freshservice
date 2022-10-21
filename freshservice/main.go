package freshservice

import (
	"context"
	"log"
)

func main() {
	company := "xxx"
	apiKey := "xxx"
	reqId := 123
	ctx := context.Background()
	Fs, err := NewClient(ctx, company, apiKey)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Obtain info for a user (Requester)
	//requester, _, err := Fs.Requesters.GetRequester(reqId)
	//log.Printf("%s %s - %s\n", requester.FirstName, requester.LastName, requester.Email)
	//
	//services, _, err := Fs.Services.ListServiceItems()
	//for _, service := range services.Collection {
	//	log.Printf("Services available: %d (%s)\n", service.ID, service.Name)
	//}

	customObjects, _, err := Fs.CustomObject.GetCustomObjectRecords(reqId)
	log.Println(customObjects.Title, customObjects.Description, customObjects.ID)
	if err != nil {
		log.Fatalf("Error while retriving custom objects")
	}
}
