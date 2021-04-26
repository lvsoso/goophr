package main

import (
        "fmt"
)


func makeHotelReservation() {
        fmt.Println("Done making hotel reservation.")
}

func bookFlightTickets() {
        fmt.Println("Done booking flight tickets.")
}

func orderADress() {
        fmt.Println("Done ordering a dress.")
}

func payCreditCardBills() {
        fmt.Println("Done paying Credit Card bills.")
}

func writeAMail() {
        fmt.Println("Wrote 1/3rd of the mail.")
        go continueWritingMail1() // Notice the addition of `go` keyword.
}

func continueWritingMail1() {
        fmt.Println("Wrote 2/3rds of the mail.")
        go continueWritingMail2() // Notice the addition of `go` keyword.
}

func continueWritingMail2() {
        fmt.Println("Done writing the mail.")
}

func listenToAudioBook() {
        fmt.Println("Listened to 10 minutes of audio book.")
        go continueListeningToAudioBook() // Notice the addition of `go` keyword.
}

func continueListeningToAudioBook() {
        fmt.Println("Done listening to audio book.")
}

var listOfTasks = []func(){
        makeHotelReservation, bookFlightTickets, orderADress,
        payCreditCardBills, writeAMail, listenToAudioBook,
}

func main() {
        for _, task := range listOfTasks {
                task()
        }
}

