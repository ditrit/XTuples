package cron_module

type CreateCronBody struct {
	Taskname string `json:"taskname" validate:"required"`
	At       string `json:"at" validate:"required"`
}

type UpdateCronBody struct {
	Taskname string `json:"taskname" validate:"required"`
	At       string `json:"at" validate:"required"`
}
