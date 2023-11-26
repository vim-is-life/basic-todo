// TODO refactor to put writing and saving in own file? maybe myio.go or sum
// TODO add input validation where appropriate

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	todoList := []Todo{}
	readTodos(&todoList)

	var choice UserChoice
	for {
		printMenu()
		fmt.Scan(&choice)
		// check if user wants to quit, then for valid input, then exec desired
		// action
		// if choice == 9 {
		// 	return
		// }
		switch choice {
		case ChoiceViewTodos:
			viewTodos(todoList)

		case ChoiceAddTodo:
			addTodo(&todoList)
			viewTodos(todoList)

		case ChoiceUpdateTodo:
			updateTodo(&todoList, getDesiredTodoIdFromUser())
			viewTodos(todoList)

		case ChoiceDeleteTodo:
			deleteOneTodo(&todoList, getDesiredTodoIdFromUser())
			viewTodos(todoList)

		case ChoiceDeleteAllTodos:
			deleteAllTodos(&todoList)
			viewTodos(todoList)

		default:
			// fmt.Println("Invalid option, please try again.")
			// continue
			os.Exit(0)
		}

		switch choice {
		case ChoiceAddTodo, ChoiceUpdateTodo, ChoiceDeleteTodo, ChoiceDeleteAllTodos:
			saveTodos(todoList)
		}
	}
}

// printMenu outputs the menu text to the user.
func printMenu() {
	menuText := `What would you like to do?
	1. View list of todos
	2. Add to the list
	3. Update an item in the list
	4. Delete an item from the list
	5. Delete all items from the list

Any other input quits the program. Your choice?: `
	fmt.Print(menuText)
}

// getDesiredTodoIdFromUser will prompt for the ID of todo user wants to change
// TODO validate input to make sure id user gives is within range
// returns the id that the user wanted
func getDesiredTodoIdFromUser() uint {
	fmt.Print("Id to modify? ")
	var userGivenId uint
	fmt.Scan(&userGivenId)
	return userGivenId
}

// viewTodos will display all todos in todoList.
func viewTodos(todoList []Todo) {
	fmt.Println()
	fmt.Printf("%5s%40s\t%10s\t%10s\n", "id", "name", "kind", "state")
	fmt.Printf("%5s%40s\t%10s\t%10s\n", "==", "====", "====", "=====")
	for _, todo := range todoList {
		fmt.Printf("%4d)%40s\t%10s\t%10s\n",
			todo.id, todo.name, todo.kind, todo.state)
	}
	fmt.Println()
}

// addTodo will add one todo item to todoList
func addTodo(todoList *[]Todo) {
	var newTodoItem Todo
	var sc = bufio.NewScanner(os.Stdin)
	newTodoItem.id = uint(len(*todoList)) + 1

	// get the name of the todo
	fmt.Print("Todo name?\n> ")
	sc.Scan()
	newTodoItem.name = sc.Text()

	// get the type of todo
	fmt.Printf(
		"Todo type? (%d for project, %d for homework, %d for reading)\n> ",
		CategoryProject, CategoryHomework, CategoryReading)
	//TODO input validation
	fmt.Scan(&newTodoItem.kind)

	// set the state of the todo
	newTodoItem.state = StateTodo
	*todoList = append(*todoList, newTodoItem)
}

// updateTodo will modify the entry in todoList with identifier id.
func updateTodo(todoList *[]Todo, id uint) {
	var choice int
	fmt.Print("Update how? (1 for type, 2 for state) ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf(
			"New type? (%d for project, %d for homework, %d for reading) ",
			CategoryProject, CategoryHomework, CategoryReading)
		fmt.Scan(&(*todoList)[id-1].kind)
	case 2:
		fmt.Printf(
			"New state? (%d for todo, %d for in progress, %d for done) ",
			StateTodo, StateInProgress, StateDone)
		fmt.Scan(&(*todoList)[id-1].state)
	}

}

// updateTodo will delete the entry in todoList with identifier id.
func deleteOneTodo(todoList *[]Todo, idToDel uint) {
	// this is one way of how to delete something from a slice, but not sure if
	// it's the most optimized necessarily
	var idxToDel int
	for idx, todo := range *todoList {
		if todo.id == idToDel {
			idxToDel = idx
			break
		}
	}
	*todoList = append((*todoList)[:idxToDel], (*todoList)[idxToDel+1:]...)
}

// deleteAllTodos removes all entries in todoList (in-place).
// The function will, however, ask for user confirmation and wait 2 seconds
// before allowing a response, so as to prevent accidental deletion.
func deleteAllTodos(todoList *[]Todo) {
	fmt.Println("Waiting 2 seconds...")
	time.Sleep(2 * time.Second)

	var userResp string
	fmt.Print("Are you sure? [y/N] ")
	fmt.Scan(&userResp)

	if strings.ToUpper(userResp) == "Y" {
		*todoList = []Todo{}
	} else {
		fmt.Println("Aborting...")
	}
}
