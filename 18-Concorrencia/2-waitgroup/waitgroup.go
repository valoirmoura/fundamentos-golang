package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	//Informa que há duas GoRoutines em execução
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() //Aqui ele faz -1 goroutine

	}()

	go func() {
		escrever("Programando em GO!")
		waitGroup.Done() //Aqui mais -1 goroutine
	}()

	//Aqui finaliza a GoRoutine pois temos 0 GO Routines.... aqui não executa antes que
	//todas a GORoutines estejam executadas....
	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
