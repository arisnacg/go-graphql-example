package todo

var TodoItems = []Todo{}

var count = 1

// get all todos
func GetTodos() []Todo {
	return TodoItems
}

// get one todo
func GetTodo(id int) Todo {
	for _, todo := range TodoItems {
		if todo.ID == id {
			return todo
		}
	}
	return Todo{}
}

// add new todo
func AddTodo(title string) *Todo {
	temp := &Todo{
		ID:        count,
		Title:     title,
		Completed: false,
	}
	TodoItems = append(TodoItems, *temp)
	count++
	return temp
}

// update existing todo
func UpdateTodo(id int, title string, completed bool) bool {
	for i, todo := range TodoItems {
		if todo.ID == id {
			TodoItems[i].Title = title
			TodoItems[i].Completed = completed
			return true
		}
	}
	return false
}

// delete existing todo
func DeleteTodo(id int) bool {
	for i, todo := range TodoItems {
		if todo.ID == id {
			TodoItems = append(TodoItems[:i], TodoItems[i+1:]...)
			return true
		}
	}
	return false
}
