package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	//todos.add("Buy milk")
	//todos.add("Buy bread")
	//todos.toggle(0)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	//todos.print()
	//todos.delete(0)
	//todos.print()
	//fmt.Printf("%+v\n\n", todos)
	storage.Save(todos)

}
