package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct{
  Add string
  Del int
  Edit string
  Toggle int
  List bool
}

func NewCmdFlags() *CmdFlags{
  cf := CmdFlags{}
  flag.StringVar(&cf.Add, "add", "", "Add a new Todo")
  flag.StringVar(&cf.Edit, "edit", "", "Edit a Todo")
  flag.IntVar(&cf.Del, "del", -1, "Delete a Todo by index")
  flag.IntVar(&cf.Toggle, "toggle", -1, "toggle a Todo by index")
  flag.BoolVar(&cf.List, "list", false, "List all Todos")

  flag.Parse()

  return &cf
}

func (cf *CmdFlags) Execute(todos *Todos){
  switch{
  case cf.Add != "":
    todos.AddTodo(cf.Add)
  case cf.Del != -1:
    todos.DeleteTodo(cf.Del)
  case cf.Edit != "":
    parts := strings.Split(cf.Edit, ":")
    if len(parts) != 2{
      fmt.Println("Invalid Edit Command, please use id:newTitle")
      os.Exit(1)
    }
    index, err := strconv.Atoi(parts[0])
    if err != nil{
      fmt.Println("Invalid Edit Command, please use id:newTitle")
      os.Exit(1)
    }
    todos.EditTodo(index, parts[1])
  case cf.Toggle != -1:
    todos.ToggleDone(cf.Toggle)
  case cf.List:
    todos.PrintTodo()
  default:
    fmt.Println("Invalid Command")
  }
}
