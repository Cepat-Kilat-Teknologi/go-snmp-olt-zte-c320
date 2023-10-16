package handler

import (
	"fmt"
	"net/http"
)

type PonHanlder struct{}

func (o *PonHanlder) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an customer")
}

func (o *PonHanlder) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all customer")
}

func (o *PonHanlder) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an customer by ID")
}

func (o *PonHanlder) GetByPonID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an customer by pon ID")
}

func (o *PonHanlder) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an customer by ID")
}

func (o *PonHanlder) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an customer by ID")
}
