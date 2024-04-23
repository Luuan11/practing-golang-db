package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insercao(arr []int) {
    for i := 1; i < len(arr); i++ {
        chave := arr[i]
        j := i - 1
        for j >= 0 && arr[j] > chave {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = chave
    }
}

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func selecao(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        menor := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[menor] {
                menor = j
            }
        }
        arr[i], arr[menor] = arr[menor], arr[i]
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    variavel := rand.Perm(100)
    variavel2 := rand.Perm(1000)
    variavel3 := rand.Perm(10000)

    // Medir o tempo de execução e ordenar 100 números

	println("===================== INICIO DOS 100 NÚMEROS =====================")

    inicio := time.Now()
    insercao(variavel)
    fmt.Println("Tempo de execução do Insertion Sort para 100 números:", time.Since(inicio))

    inicio = time.Now()
    bubbleSort(variavel)
    fmt.Println("Tempo de execução do Bubble Sort para 100 números:", time.Since(inicio))

    inicio = time.Now()
    selecao(variavel)
    fmt.Println("Tempo de execução do Selection Sort para 100 números:", time.Since(inicio))

	println("===================== FIM DOS 100 NÚMEROS =====================")

	println("===================== INICIO DOS 1000 NÚMEROS =====================")
    inicio = time.Now()
    insercao(variavel2)
    fmt.Println("Tempo de execução do Insertion Sort para 1000 números:", time.Since(inicio))

    inicio = time.Now()
    bubbleSort(variavel2)
    fmt.Println("Tempo de execução do Bubble Sort para 1000 números:", time.Since(inicio))

    inicio = time.Now()
    selecao(variavel2)
    fmt.Println("Tempo de execução do Selection Sort para 1000 números:", time.Since(inicio))

	println("===================== FIM DOS 1000 NÚMEROS =====================")

	println("===================== INICIO DOS 10000 NÚMEROS =====================")

    inicio = time.Now()
    insercao(variavel3)
    fmt.Println("Tempo de execução do Insertion Sort para 10000 números:", time.Since(inicio))

    inicio = time.Now()
    bubbleSort(variavel3)
    fmt.Println("Tempo de execução do Bubble Sort para 10000 números:", time.Since(inicio))

    inicio = time.Now()
    selecao(variavel3)
    fmt.Println("Tempo de execução do Selection Sort para 10000 números:", time.Since(inicio))

	println("===================== FIM DOS 10000 NÚMEROS =====================")
}