package clickhause

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/ch-go"

	"git_p/test/insert/db/postgres"
)

// отпровляет данные в CliacHause
func SetMigrationDates(conCH *ch.Client, payloadres []postgres.Item) (err error) {
	var query string = "INSERT INTO migrations (Id, Campaignid, Name,Description,Priority,Removed,EventTime) VALUES\n"

	for _, item := range payloadres {
		piece := fmt.Sprintf("( %d , %d , '%s' , '%s' , %d , %t , toDateTime('%s') ),\n", item.Id, item.CampaignId, item.Name, item.Description, item.Priority, item.Removed, item.CreatedAt.Format("2006-01-02 15:04:05"))
		query += piece
	}
	//println(query)
	err = conCH.Do(context.Background(), ch.Query{Body: query})
	if err != nil {
		log.Println("ERROR conCH.Do:", err)
		return
	}

	return nil
}
