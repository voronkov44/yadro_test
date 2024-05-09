package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Client struct {
	Name      string
	Table     int
	Arrival   string
	Departure string
}

func main() {
	var tables int
	var openTime, closeTime string
	var clients = make(map[string]Client)
	var waitList []string

	file, err := os.Open("input_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) == 1 {
			tables, _ = strconv.Atoi(line[0])
		} else if len(line) == 2 {
			openTime = line[0]
			closeTime = line[1]
			fmt.Println(openTime)
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
					fmt.Println(eventTime, 13, "ClientUnknown")
					continue
				}
				table, _ := strconv.Atoi(line[3])
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
				if len(waitList) > 0 {
					firstClient := waitList[0]
					waitList = waitList[1:]
					table, _ := strconv.Atoi(line[3])
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
}
