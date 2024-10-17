package exec_module

type ExecBody struct {
	Taskname string `json:"taskname" validate:"required"`
}
