package tech

import (
	"fmt"
	"surlit/internal/app/tech/errors"
)

type Option struct {
	Name string
	Handle
}

func (a *App) printMenu(options []Option) {
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option.Name)
	}
}

func (a *App) chooseOption(options []Option) (*Option, error) {
	fmt.Printf("choose option: ")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		// flush stdout
		var discard string
		_, _ = fmt.Scanln(&discard)
		return nil, errors.ErrWrongInput
	}
	if choice < 1 || choice > len(options) {
		return nil, errors.ErrWrongChoice
	}
	return &options[choice-1], nil
}

func (a *App) Serve() {
	options := []Option{
		{
			Name:   "Login",
			Handle: a.service.loginHandler,
		},
		{
			Name:   "Register",
			Handle: a.service.registrationHandler,
		},
		{
			Name:   "User projects",
			Handle: a.service.getUserProjectsHandler,
		},
		{
			Name:   "User links",
			Handle: a.service.getUserLinksHandler,
		},
		{
			Name:   "Link routes",
			Handle: a.service.getLinkRoutesHandler,
		},
		{
			Name:   "Route stats",
			Handle: a.service.getLinkRoutesHandler,
		},
	} // TODO add more options
	for {
		a.printMenu(options)
		option, err := a.chooseOption(options)
		if err != nil {
			fmt.Println(err)
			continue
		}
		option.Handle()
	}
}
