package database

import (
	"context"
)

type Formulir struct {
	no_upcm         string
	niu             string
	kd_status_bayar string
}

func GetAll() (result []Formulir) {
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
	return result
}
