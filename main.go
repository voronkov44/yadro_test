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
	Name         string
	Table        int
	Arrival      string
	Departure    string
	PricePerHour int
}

func (c *Client) calculateRevenue(openTime, closeTime string) int {
	if c.Departure == "" {
		return 0
	}

	arrivalHour, _ := strconv.Atoi(strings.Split(c.Arrival, ":")[0])
	departureHour, _ := strconv.Atoi(strings.Split(c.Departure, ":")[0])
	totalHours := departureHour - arrivalHour

	if totalHours < 1 {
		totalHours = 1
	}

	if c.Arrival < openTime {
		arrivalHour, _ = strconv.Atoi(strings.Split(openTime, ":")[0])
	}
	if c.Departure > closeTime {
		departureHour, _ = strconv.Atoi(strings.Split(closeTime, ":")[0])
	}

	return totalHours * c.PricePerHour
}

func main() {
	var tables int
	var openTime, closeTime string
	var clients = make(map[string]Client)
	var waitList []string
	totalRevenue := make(map[int]int)
	workingTime := make(map[int]int)

	file, err := os.Open("C:\\Users\\drop-\\GolandProjects\\awesomeProject\\yadro_test\\test\\input_file.txt")
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

				revenue := client.calculateRevenue(openTime, closeTime)
				totalRevenue[clients[eventName].Table] += revenue

				arrivalHour, _ := strconv.Atoi(strings.Split(clients[eventName].Arrival, ":")[0])
				departureHour, _ := strconv.Atoi(strings.Split(clients[eventName].Departure, ":")[0])
				workingHours := departureHour - arrivalHour
				workingTime[clients[eventName].Table] += workingHours

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

	// Print total revenue and working time for each table
	for tableNum, revenue := range totalRevenue {
		workingHours := workingTime[tableNum]
		workingMinutes := workingHours % 60
		workingHours /= 60
		fmt.Printf("%d %d %02d:%02d\n", tableNum, revenue, workingHours, workingMinutes)
	}
}
