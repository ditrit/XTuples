package cron_module

import (
	"fmt"
	"go-http/pkg/settings/database"
)

type SCron struct {
	Id        string `json:"cron_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Taskname  string `json:"task_name"`
	At        string `json:"at"`
}

func GetAllCronsSQL() (any, error) {
	query := "SELECT * FROM crons"

	rows, err := database.GetDB().Query(query)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var crons []SCron
	for rows.Next() {
		var cron SCron
		err = rows.Scan(&cron.Id, &cron.CreatedAt, &cron.UpdatedAt, &cron.Taskname, &cron.At)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		crons = append(crons, cron)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return crons, nil
}

func GetCronSQL(id string) (any, error) {
	query := fmt.Sprintf("SELECT * FROM crons WHERE cron_id='%s'", id)

	var cron SCron
	row := database.GetDB().QueryRow(query)
	err := row.Scan(&cron.Id, &cron.CreatedAt, &cron.UpdatedAt, &cron.Taskname, &cron.At)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cron, nil
}

func CreateCronSQL(data CreateCronBody) (any, error) {
	query := fmt.Sprintf("INSERT INTO crons (task_name,at) VALUES ('%s','%s')", data.Taskname, data.At)

	res, err := database.GetDB().Exec(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	AddCron(SCron{At: data.At, Taskname: data.Taskname})
	return res, nil
}

func DeleteCronSQL(id string) (any, error) {
	query := fmt.Sprintf("DELETE FROM crons WHERE cron_id='%s'", id)

	res, err := database.GetDB().Exec(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ResetCrons()
	return res, nil
}

func UpdateCronSQL(id string, data UpdateCronBody) (any, error) {
	query := fmt.Sprintf("UPDATE crons SET task_name='%s', at='%s' WHERE cron_id='%s'", data.Taskname, data.At, id)

	res, err := database.GetDB().Exec(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ResetCrons()
	return res, nil
}
