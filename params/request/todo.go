package request

type CreateTodo struct {
	Name  string `example:"Name TODO"`
	Email string `example:"Email TODO"`
}

type UpdateTodo struct {
	Name  string `example:"Name TODO"`
	Email string `example:"Email TODO"`
}
