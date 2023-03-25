package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/anibal-alpizar/go-cli-crud/tasks"
)

func main() {
	// read or create the tasks.json file
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var taskss []tasks.Task

	// get file info
	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	// check the size
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &taskss)
		if err != nil {
			panic(err)
		}
	} else {
		taskss = []tasks.Task{}
	}
	if len(os.Args) < 2 {
		printUsage()
	}
	switch os.Args[1] {
	case "list":
		tasks.ListTasks(taskss)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Nombre de la tarea:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		taskss = tasks.AddTask(taskss, name)
		tasks.SaveTasks(file, taskss)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Debes especificar el ID de la tarea a eliminar")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un número")
			return
		}
		taskss = tasks.DeleteTask(taskss, id)
		tasks.SaveTasks(file, taskss)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Debes especificar el ID de la tarea a completar")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un número")
			return
		}
		taskss = tasks.CompleteTask(taskss, id)
		tasks.SaveTasks(file, taskss)
	}
}

func printUsage() {
	fmt.Println("Uso: go-clid-crud [list|add|complete|delete]")
}
