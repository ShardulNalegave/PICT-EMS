package tsdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ShardulNalegave/PICT-EMS/sessions"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type TSDB struct {
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI
}

func ConnectToTSDB() *TSDB {
	addr := os.Getenv("INFLUX_ADDR")
	token := os.Getenv("INFLUX_TOKEN")
	org := os.Getenv("INFLUX_ORG")
	bucket := os.Getenv("INFLUX_BUCKET")

	client := influxdb2.NewClient(addr, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	queryAPI := client.QueryAPI(org)

	return &TSDB{client: client, writeAPI: writeAPI, queryAPI: queryAPI}
}

func (t *TSDB) WriteSessionData(s *sessions.Session) error {
	exit := time.Now()
	p := influxdb2.NewPointWithMeasurement(os.Getenv("INFLUX_MEASUREMENT")).
		AddTag("location", s.Location).
		AddTag("registration_id", s.RegistrationID).
		AddField("entry_time", s.EntryTime).
		AddField("exit_time", exit).
		SetTime(s.EntryTime)

	return t.writeAPI.WritePoint(context.Background(), p)
}

func (t *TSDB) GetSessions(start time.Time, stop time.Time, location string) []Session {
	q := fmt.Sprintf(`
	from(bucket: "%s")
		|> range(start: %s, stop: %s)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["location"] == "%s")
		|> sort(columns: ["_time"])
		|> pivot(rowKey: ["_time", "location", "registration_id"], columnKey: ["_field"], valueColumn: "_value")
	`, os.Getenv("INFLUX_BUCKET"), start.Format(time.RFC3339), stop.Format(time.RFC3339), os.Getenv("INFLUX_MEASUREMENT"), location)
	res, err := t.queryAPI.Query(context.Background(), q)
	if err != nil {
		panic(err)
	}
	defer res.Close()

	s := make([]Session, 0)

	for res.Next() {
		record := res.Record()
		entry, _ := time.Parse(time.RFC3339, record.ValueByKey("entry_time").(string))
		exit, _ := time.Parse(time.RFC3339, record.ValueByKey("exit_time").(string))
		s = append(s, Session{
			Location:       record.ValueByKey("location").(string),
			RegistrationID: record.ValueByKey("registration_id").(string),
			EntryTime:      entry,
			ExitTime:       exit,
		})
	}

	if res.Err() != nil {
		log.Fatal(res.Err())
	}

	return s
}
