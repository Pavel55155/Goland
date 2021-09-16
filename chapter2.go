package main

//this is a comment

import "fmt"

func main (){
	fmt.Println("Введите температуру в Фаренгейтах: ")
	var F float64
	fmt.Scanf("%f", &F)

	C := (F -32) * 5/9

	fmt.Println("Температура в Цельсиях:", C)
	}