package main

func main(){
  t := Todos{}
  storage := NewStorage[Todos]("todos.json")
  storage.Load(&t)

  CmdFlags := NewCmdFlags()
  CmdFlags.Execute(&t)

  storage.Save(t)
}
