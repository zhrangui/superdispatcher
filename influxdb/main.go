package main

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func main() {
	// You can generate a Token from the "Tokens Tab" in the UI
	const token = "0-qE1miPLBrvHtJ4BZZF631mlTJyIL8vTfooyfE-az04sSRSk6YEcTKL818jywkDjUufmmdJQmoRG0aoxXznlQ=="

	const bucket = "bucketID"
	const org = "e9a9c308a2513165"

	client := influxdb2.NewClient("https://us-west-2-1.aws.cloud2.influxdata.com", token)
	// always close client at the end
	defer client.Close()

	// get non-blocking write client
	writeApi := client.WriteApi(org, bucket)

	// write line protocol
	writeApi.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	writeApi.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 22.5, 45.0))
	// Flush writes
	writeApi.Flush()

	query := fmt.Sprintf("from(bucket:\"%v\")|> range(start: -1h) |> filter(fn: (r) => r._measurement == \"stat\")", bucket)
	// Get query client
	queryApi := client.QueryApi(org)
	// get QueryTableResult
	result, err := queryApi.Query(context.Background(), query)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}
