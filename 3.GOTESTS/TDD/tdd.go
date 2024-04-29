package main

func fibonacci(arg int) int {
    if arg <= 0{
        return 0
    } else if arg == 1 || arg == 2 {
        return 1
    }

    data := make([]int, arg)
    data[0] = 1
    data[1] = 1

    for i := 2; i < len(data); i++ {
        data[i] = data[i-2] + data[i-1]
    }
    return data[len(data)-1]
}
