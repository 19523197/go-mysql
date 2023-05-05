package database

import (
	"context"
	"log"
	"testing"
)

func TestExecSql(t *testing.T) {
	var result []Formulir
	var formulir Formulir
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	param1 := "2209090001"
	param2 := "1"
	query1 := "SELECT no_upcm, niu,kd_status_bayar FROM formulir WHERE niu = ? AND kd_status = ?"
	rows, err := db.QueryContext(ctx, query1, param1, param2)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&formulir.no_upcm, &formulir.niu, &formulir.kd_status_bayar)
		if err != nil {
			panic(err)
		}

		result = append(result, formulir)

	}

}

func TestPrepareStatement(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query1 := "SELECT no_upcm, niu,kd_status_bayar FROM formulir WHERE niu = ? AND kd_status = ?"
	statement, err := db.PrepareContext(ctx, query1)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		niu := i
		kd_status_bayar := "Y"

		result, err := statement.QueryContext(ctx, niu, kd_status_bayar)
		if err != nil {
			panic(err)
		}
		log.Println(result)
	}

}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	param1 := "2209090001"
	param2 := "1"
	query1 := "SELECT no_upcm, niu,kd_status_bayar FROM formulir WHERE niu = ? AND kd_status = ?"
	rows, err := db.ExecContext(ctx, query1, param1, param2)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insertId, err := rows.LastInsertId()

	if err != nil {
		panic(err)
	}

	log.Println(insertId)
}
