package messages

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func StartClient(port string,locations []string) {
	for i,location := range locations {
		if i == len(locations) -1 {
			break
		}
		conn, err := grpc.DialContext(context.Background(), location+":"+port, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err.Error())
		}

		client := NewTopoServiceClient(conn)
		topography, err := client.ReportDistance(context.Background(), &Locations{
			Location: locations[i+1:],
		})
		if err != nil {
			log.Println(err.Error())
			continue
		}
		for key, value := range topography.Distances {
			log.Printf("Location %s is at %e distance\n", key, value*1E-9)
		}
	}
}