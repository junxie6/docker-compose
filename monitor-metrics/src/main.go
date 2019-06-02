// Reference:
// https://github.com/bentol/prometheus-golang-example
// https://github.com/brancz/prometheus-example-app
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	DB *sqlx.DB

	goroutinesGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "mysql_row_count",
			Help: "number of rows",
		},
		[]string{"hostname"},
	)
)

func observer() {
	hostname, _ := os.Hostname()

	var err error
	var value int

	for {
		if value, err = GetCount(); err != nil {
			value = -1
		}

		goroutinesGauge.With(prometheus.Labels{
			"hostname": hostname,
		}).Set(float64(value))

		time.Sleep(1 * time.Second)
	}
}

func InitDB() error {
	var err error
	DB, err = sqlx.Connect("mssql", "sqlserver://sa:Change_Admin_Pass6@127.0.0.1:1433")

	if err != nil {
		return err
	}

	return nil
}

func GetCount() (int, error) {
	var err error
	var rows *sqlx.Rows
	data := struct {
		Number int `db:"Number"`
	}{}

	sq := `
SELECT COUNT(*) AS Number FROM [exp].[dbo].[Bug] WHERE CreatedDT >= DATEADD(MINUTE, -5, SYSDATETIME())
`
	if rows, err = DB.Queryx(sq); err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&data)

		if err != nil {
			return 0, err
		}
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return data.Number, nil
}

func main() {
	var err error

	if err = InitDB(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	defer DB.Close()

	//
	prometheus.MustRegister(goroutinesGauge)
	go observer()

	http.Handle("/metrics", prometheus.Handler())

	if err = http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
}
