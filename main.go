package main

import (
	"fmt"
)

type Room struct {
	ID       int
	Type     string
	Status   bool // True:Available -- False:Reserved
	BedCount int
	Price    int
}

var Rooms []Room = GenerateRooms()

func main() {
	input := ""
	for input != "exit" {
		fmt.Println("Enter a Number :")
		fmt.Println("1. Available Rooms List :")
		fmt.Println("2. Add Room :")
		fmt.Println("3. Reserve Room :")
		fmt.Println("4. Remove Room :")
		fmt.Scanln(&input)
		switch input {
		case "1":
			GetRoomList()
		case "2":
			AddRoom()
		case "3":
			ReserveRoom()
		case "4":
			RemoveRoom()
		case "exit":
			fmt.Println("Exiting ...")
		default:
			fmt.Println("Invalid Command")
		}
	}

}

func GetRoomList() {
	foundRooms := false

	for _, room := range Rooms {
		if room.Status {
			foundRooms = true
			fmt.Printf("<-----------------> \n")
			fmt.Printf("Room ID: %d\nType: %s\nBed Count: %d\nPrice: %d\n", room.ID, room.Type, room.BedCount, room.Price)
			fmt.Printf("<----------------->\n\n")
		}
	}

	if !foundRooms {
		fmt.Printf("<-----------------> \n")
		fmt.Println("No rooms available.")
		fmt.Printf("<-----------------> \n")

	}
}
func AddRoom() {
	var roomType string
	var bedCount int
	var price int

	fmt.Print("Enter room type: ")
	fmt.Scanln(&roomType)

	fmt.Print("Enter bed count: ")
	fmt.Scanln(&bedCount)

	fmt.Print("Enter price: ")
	fmt.Scanln(&price)

	lastRoom := Rooms[len(Rooms)-1]
	newID := lastRoom.ID + 1
	status := true
	newRoom := Room{
		ID:       newID,
		Type:     roomType,
		Status:   status,
		BedCount: bedCount,
		Price:    price,
	}
	updatedRooms := append(Rooms, newRoom)
	if len(updatedRooms) == len(Rooms) {
		fmt.Println("Operation Failed!")
	} else {
		Rooms = updatedRooms
		lastRoom := Rooms[len(Rooms)-1]
		fmt.Println("<-------------------------------->")
		fmt.Println("Room successfully added.")
		fmt.Printf("Room ID:%d\nRoom Type:%s\nStatus:%t\nBed Count:%d\nPrice:%d\n", lastRoom.ID, lastRoom.Type, lastRoom.Status, lastRoom.BedCount, lastRoom.Price)
		fmt.Println("<-------------------------------->")
	}

}
func RemoveRoom() {
	var RoomID int
	fmt.Println("Enter Room ID:")
	fmt.Scanln(&RoomID)
	fmt.Println(Rooms)
	Rooms = append(Rooms[:RoomID-1], Rooms[RoomID:]...)
	fmt.Printf("Room with ID: %d Deleted successfully\n", RoomID)
}

func ReserveRoom() {
	var id int
	fmt.Println("Enter Room ID to reserve :")
	fmt.Scanln(&id)
	for i, room := range Rooms {
		if room.ID == id {
			Rooms[i].Status = false
			fmt.Println(Rooms[i])
			fmt.Printf("<-----------------> \n")
			fmt.Printf("Room %d reserved successfully\n", id)
			fmt.Printf("<-----------------> \n")
			return
		}
	}

}

func GenerateRooms() []Room {
	rooms := []Room{}
	rooms = append(rooms, Room{ID: 1, Type: "Single", Status: false, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID: 2, Type: "Single", Status: false, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID: 3, Type: "Double", Status: false, BedCount: 2, Price: 200})
	rooms = append(rooms, Room{ID: 4, Type: "Double", Status: false, BedCount: 2, Price: 250})
	rooms = append(rooms, Room{ID: 5, Type: "Double", Status: false, BedCount: 2, Price: 270})
	rooms = append(rooms, Room{ID: 6, Type: "Double", Status: false, BedCount: 2, Price: 150})
	rooms = append(rooms, Room{ID: 7, Type: "Standard", Status: false, BedCount: 3, Price: 300})
	rooms = append(rooms, Room{ID: 8, Type: "Standard", Status: false, BedCount: 3, Price: 300})
	rooms = append(rooms, Room{ID: 9, Type: "Standard", Status: false, BedCount: 3, Price: 400})
	rooms = append(rooms, Room{ID: 10, Type: "Standard", Status: true, BedCount: 3, Price: 310})
	return rooms
}
