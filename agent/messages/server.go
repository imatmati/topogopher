package messages

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
	"time"
)

type TopogopherServer struct {
}

func (s TopogopherServer) Call(location, localAddress string) float64 {

	// gRpc
	conn, err := grpc.Dial(location+":8080", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Println(err.Error())
		return math.MaxFloat64
	}
	start := time.Now()
	client := NewTopoServiceClient(conn)
	_, err = client.Ping(context.Background(), &Void{})
	if err != nil {
		log.Println(err.Error())
		return math.MaxFloat64
	}
	duration := time.Now().Sub(start)
	log.Printf("gRpc : From %s to %s duration is %e\n", localAddress, location,float64(duration)*1E-9)
	if err != nil {
		log.Println(err.Error())
		return math.MaxFloat64
	}

	return float64(duration)
}


func (s TopogopherServer) Ping(context.Context, *Void) (*PingResponse, error) {
	return &PingResponse{},nil
}

func (s *TopogopherServer) ReportDistance(ctx context.Context, loc *Locations) (*Topography, error) {
	t := &Topography{Distances: make(map[string]float64)}

	addrs,err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, l := range loc.Location {
		var distance float64
	OUT:
		for i := 0; i < 3; i++ {
			distance += s.Call(l,addrs[0].String())
			if distance == math.MaxFloat64 {
				t.Distances[l] = distance
				break OUT
			}
		}
		t.Distances[l] = distance / 3
	}
	return t, nil
}

func StartServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err.Error())
	}
	s := grpc.NewServer()

	RegisterTopoServiceServer(s, &TopogopherServer{})
	s.Serve(lis)
}
