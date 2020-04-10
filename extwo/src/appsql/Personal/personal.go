package personal

import (
	"fmt"
	"net/http"
	"encoding/json"
	"appsql/Config"	
)

type dataPersonal struct {
	Id 			int
	Name 		string `json:"Name"`
	Age 		int
	Is_trash 	int
	Status 		int
}

func ReadPersonal(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db := config.DBOp()
	tsql := fmt.Sprintf("SELECT id, name, age, is_trash, status FROM _personal_golang;")
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
	}
	
	var res []dataPersonal

	for rows.Next() {
		var personal = dataPersonal{}
		errs := rows.Scan(&personal.Id, &personal.Name, &personal.Age, &personal.Is_trash, &personal.Status)

		if errs != nil {
			fmt.Println("Error reading rows: " + errs.Error())
		}

		res = append(res, personal)
	}

	var strPersonal, errs = json.Marshal(res)
	if errs != nil {
		fmt.Println("Error reading rows: " + errs.Error())
	}

	w.Write(strPersonal)
	defer db.Close()
}