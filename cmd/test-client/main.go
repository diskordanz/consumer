package main

import (
	"fmt"
	"os"
	"os/exec"

	fHandler "github.com/diskordanz/consumer/pkg/franchise/handler"
	uHandler "github.com/diskordanz/consumer/pkg/user/handler"
	"github.com/diskordanz/consumer/service"

	fStorage "github.com/diskordanz/consumer/pkg/franchise/storage"
	uStorage "github.com/diskordanz/consumer/pkg/user/storage"
)

func main() {
	franchiseStorage := fStorage.NewInMemoryFranchiseStorage()
	userStorage := uStorage.NewInMemoryUserStorage()
	franchiseHandler := *fHandler.NewFranchiseHandler(franchiseStorage)
	userHandler := *uHandler.NewUserHandler(userStorage)
	srv := service.NewConsumerService(franchiseHandler, userHandler)
	Run(srv)
}

func Run(srv service.Service) {
	var choise string
	for {
		fmt.Println("Введите действие из предложенных: ")
		fmt.Println("1) вывести френчайзы")
		fmt.Println("2) вывести френчайз")
		fmt.Println("3) вывести всех юзеров")
		fmt.Println("4) вывести юзера")
		fmt.Println("5) вывести юзеров френчайза")
		fmt.Fscan(os.Stdin, &choise)
		if choise != "1" && choise != "2" && choise != "3" && choise != "4" && choise != "5" {
			fmt.Println("Неверное действие! попробуйте еще раз")
			continue
		}
		switch choise {
		case "1":
			fmt.Println(srv.GetFranchisesList(3, 0))
			fmt.Scanln()
		case "2":
			fmt.Println(srv.GetFranchise(2))
			fmt.Scanln()
		case "3":
			fmt.Println(srv.GetUsers(3, 0))
			fmt.Scanln()
		case "4":
			fmt.Println(srv.GetUser(2))
			fmt.Scanln()
		case "5":
			fmt.Println(srv.GetUsersOfFranchise(1, 3, 0))
			fmt.Scanln()
		}

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
}
