package main

/*
Problem 08* - Find the spammer
Write a program which reads a set of lines which shows all your email messages, with date as first column, sender as second column & subject as final column.

Find what’s the average emails per day you receive by sender.
The total number of days is based on the range of days which you have given. For example, if you’re given emails from 2021-01-01 to 2021-01-03 only,
the total number of days is three.

Finally, print the list of senders which are likely to be spammers - those which send you more than 3 emails per day on average.

Print the results sorted based on the average emails per day in descending order.

*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	m := make(map[string]int)
	CreateFileAndPopulateItIfEmpty()

	file, err := os.Open("emails")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}
		xs := strings.Split(line, " ")
		for _, v := range xs {
			if strings.Contains(v, "@") {
				m[v] += 1
			}
		}
	}

	for k, v := range m {
		if v > 3 {
			fmt.Printf("Email '%s' has sent you %d emails, so it is a spam.\n", k, v)
		} else {
			fmt.Printf("Email '%s' has sent you %d emails, so it is NOT a spam.\n", k, v)
		}
	}

}

func CreateFileAndPopulateItIfEmpty() {
	_, err := os.Stat("emails")
	if os.IsNotExist(err) {
		f, err := os.Create("emails")
		if err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println("File successfully created.")
		}
		defer f.Close()

		f.WriteString(`
		2021-01-01 steve@gmail.com Long time no see...
		2021-01-02 info@alibaba.com Your order is on the way...
		2021-01-03 steve@gmail.com The secrets of the rich...
		2021-01-03 info@amazon.com Your book is on its way!
		2021-01-03 info@alibaba.com Your order has been delivered.
		2021-01-01 steve@gmail.com Happy new year!
		2021-01-02 michael@yahoo.com Hope you enjoyed your holidays!
		2021-01-02 steve@gmail.com Be successful this year with my new course!
		2021-01-03 info@alibaba.com Rate our services!
		2021-01-03 steve@gmail.com Buy bitcoin in 2021!
		2021-01-03 kate@gmail.com I found this great course!
		2021-01-03 steve@gmail.com See what others have to say about it...
		2021-01-03 michael@yahoo.com Let's meet up!
		2021-01-02 steve@gmail.com Don't miss out...
		2021-01-01 kate@gmail.com Happy Holidays!
		2021-01-01 steve@gmail.com Did you make your goals for 2021? I did!
		2021-01-03 kate@gmail.com How are you doing?
		2021-01-03 steve@gmail.com Here's what the rich do and the poor don't!
		2021-01-03 steve@gmail.com Last chance!
		`)
	} else {
		fmt.Println("File already exist.")
		fmt.Println()
	}
}
