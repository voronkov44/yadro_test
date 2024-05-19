package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Name         string
	Table        int
	Arrival      string
	Departure    string
	PricePerHour int
}

func main() {
	var tables int
	//var rev int
	var openTime, closeTime string
	var clients = make(map[string]Client)
	var waitList []string
	workingTime := make(map[int]int)

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Чтение количества столов
	scanner.Scan()
	numTablesStr := scanner.Text()
	numTables, _ := strconv.Atoi(numTablesStr)

	// Чтение времени работы
	scanner.Scan()
	workingHours := strings.Split(scanner.Text(), " ")
	if len(workingHours) >= 2 {
		openTime = workingHours[0]
		closeTime = workingHours[1]
		fmt.Println(openTime)
	}

	// Чтение цены стола за час
	scanner.Scan()
	tablePriceStr := scanner.Text()
	tablePrice, _ := strconv.Atoi(tablePriceStr)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка чтения файла", err)
		}

		if len(line) == 1 {
			tables, _ = strconv.Atoi(line[0])
		} else if len(line) == 2 {
			openTime = line[0]
			closeTime = line[1]
		} else if len(line) >= 3 {
			eventTime := line[0]
			eventID, _ := strconv.Atoi(line[1])
			eventName := strings.Join(line[2:], " ")

			switch eventID {
			case 1:
				if _, ok := clients[eventName]; ok {
					fmt.Println(eventTime, 13, "YouShallNotPass")
					continue
				}
				if eventTime < openTime || eventTime > closeTime {
					fmt.Println(eventTime, 13, "NotOpenYet")
					continue
				}
				clients[eventName] = Client{Name: eventName, Arrival: eventTime}
				fmt.Println(eventTime, eventID, eventName)
			case 2:
				client, ok := clients[eventName]
				if !ok {
					fmt.Println(eventTime, 2, eventName)
					continue
				}
				table, _ := strconv.Atoi(line[1])
				if client.Table != 0 && client.Table == table {
					continue
				}
				if client.Table != 0 {
					fmt.Println(eventTime, 13, "PlaceIsBusy")
					continue
				}
				client.Table = table
				clients[eventName] = client
				fmt.Println(eventTime, eventID, eventName, table)
			case 3:
				if len(clients) < tables {
					fmt.Println(eventTime, 13, "ICanWaitNoLonger!")
					continue
				}
				waitList = append(waitList, eventName)
				fmt.Println(eventTime, eventID, eventName)
			case 4:
				client, ok := clients[eventName]
				if !ok {
					fmt.Println(eventTime, 13, "ClientUnknown")
					continue
				}

				client.Departure = eventTime
				delete(clients, eventName)
				fmt.Println(eventTime, eventID, eventName)

				arrivalHour, _ := strconv.Atoi(strings.Split(clients[eventName].Arrival, ":")[0])
				departureHour, _ := strconv.Atoi(strings.Split(clients[eventName].Departure, ":")[0])
				workingHours := departureHour - arrivalHour
				workingTime[clients[eventName].Table] += workingHours

				if len(waitList) > 0 {
					firstClient := waitList[0]
					waitList = waitList[1:]
					table, _ := strconv.Atoi(line[1])
					client.Table = table
					clients[firstClient] = client
					fmt.Printf("%s %d %s %d\n", eventTime, 12, firstClient, table)
				}
			}
		}
	}

	var remainingClients []string
	for clientName := range clients {
		remainingClients = append(remainingClients, clientName)
	}
	sort.Strings(remainingClients)
	for _, clientName := range remainingClients {
		fmt.Printf("%s 11 %s\n", closeTime, clientName)
	}

	fmt.Println(closeTime)

	for i := 1; i <= numTables; i++ {
		profit := (tablePrice * (timeDiffInMinutes(workingHours[0], workingHours[1]) / 60))

		fmt.Printf("%d %d %s\n", i, profit, workingHours[1])
	}

}
func timeDiffInMinutes(startTime string, endTime string) int {
	start, _ := timeFromStr(startTime)
	end, _ := timeFromStr(endTime)

	diff := end.Sub(start)
	return int(diff.Minutes())
}

func timeFromStr(str string) (time.Time, error) {
	t, err := time.Parse("15:04", str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	return t, err
}
