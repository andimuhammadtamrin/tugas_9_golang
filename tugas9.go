package main

import (
    "errors"
    "fmt"
    "log"
    "strconv"
)

func checkAge(umur int) (bool, error) {
    if umur <= 17 {
        return false, errors.New("Belum bisa memilih")
    }
    return true, nil

}

func main() {
    var umurPeserta string
    fmt.Print("Masukkan umur : ", umurPeserta)
    fmt.Scanln(&umurPeserta)

    defer fmt.Println("Terimakasih atas Partisipasinya")

    var number int
    var err error
    number,err = strconv.Atoi(umurPeserta)

    peserta, err := checkAge(number)

    if err != nil {
      log.Fatal(err)

    }
    fmt.Println(peserta,"Bisa memilih")
    fmt.Println()

}
