package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

type Todo struct {
  Title string
  Done bool
}

type Todos []Todo

func (t *Todos)AddTodo(title string){
  fmt.Println("add todo")
  todo := Todo{
    Title: title,
    Done: false,
  }
  *t = append(*t, todo)
  t.PrintTodo()
}

func (t *Todos) validateIndex(index int) error {
  if index < 0 || index >= len(*t){
    return errors.New("Invalid Index")
  }
  return nil
}

func (todos *Todos) DeleteTodo(id int){
  
  t := *todos

  if err := t.validateIndex(id); err != nil{
    panic(err)
  }

  *todos = append(t[:id], t[id+1:]...)
}


func (t *Todos) ToggleDone(id int){
  if err := t.validateIndex(id); err != nil{
    panic(err)
  }

  
  todos := *t
  todos[id].Done = !todos[id].Done
  t.PrintTodo()
}

func (t *Todos) EditTodo(id int, title string){
  if err := t.validateIndex(id); err != nil{
    panic(err) 
  }

  todos := *t
  todos[id].Title = title
  t.PrintTodo()
}

func (todos *Todos) PrintTodo(){
  table := table.New(os.Stdout)

  table.SetRowLines(false)
  table.SetHeaders("ID", "Title", "Done")

  for index, todo := range *todos{
    var doneMark string 
    if todo.Done {
      doneMark = "✅"
    } else{
      doneMark = "❌"
    }
    table.AddRow(strconv.Itoa(index), todo.Title, doneMark)
  }

  table.Render()
}
