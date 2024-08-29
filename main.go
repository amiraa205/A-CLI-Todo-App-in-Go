package main

//import "golang.org/x/mod/sumdb/storage"

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	todos.print()
	storage.Save(todos)

}
