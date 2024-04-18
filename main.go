package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Player struct {
	bag       Bag
	activRoom *Room
}

type Bag struct {
	slots  int
	object []string
}

type Room struct {
	doorIsOpen        bool
	startMessage      string
	objectsIn         []string
	lookAroundMessage []string
	orderPlaceIn      []string
	neibours          []string
	whereObjectsIn    map[string][]string
}

func (room Room) StartMessage() string {

	canGo := "можно пройти - " + strings.Join(room.neibours, ", ")

	return room.startMessage + canGo
}

func (room Room) LookAround() string {
	start := room.lookAroundMessage[0]
	end := room.lookAroundMessage[1]
	placeWithObject := ""
	localPlaceWithObject := ""

	for _, i := range room.orderPlaceIn {

		localPlaceWithObject = ""

		localPlaceWithObject += fmt.Sprintf("%s: ", i)
		for j := 0; j < len(room.whereObjectsIn[i]); j++ {

			if slices.Contains(room.objectsIn, room.whereObjectsIn[i][j]) {
				localPlaceWithObject += room.whereObjectsIn[i][j]
				localPlaceWithObject += ", "
			}
		}

		if localPlaceWithObject == i+": " || localPlaceWithObject == ", "+i+": " {
			continue
		}

		placeWithObject += localPlaceWithObject
	}

	if placeWithObject != "" {
		if Player1.activRoom != &Kitchen {
			placeWithObject = placeWithObject[:len(placeWithObject)-2] + ". "
		}

	} else {
		switch Player1.activRoom {
		case &Bedroom:
			placeWithObject = "пустая комната. "
		}
	}

	canGo := "можно пройти - " + strings.Join(room.neibours, ", ")

	return start + placeWithObject + end + canGo
}

func (Player) takeObject(object string) string {

	index := slices.Index(Player1.activRoom.objectsIn, object)
	if index != -1 && slices.Contains(objectsType["расходники"], object) {
		if len(Player1.bag.object) < Player1.bag.slots {
			Player1.bag.object = append(Player1.bag.object, object)
			Player1.activRoom.objectsIn[index] = Player1.activRoom.objectsIn[len(Player1.activRoom.objectsIn)-1]
			Player1.activRoom.objectsIn = Player1.activRoom.objectsIn[:len(Player1.activRoom.objectsIn)-1]

			switch object {
			case "конспекты":
				Kitchen.lookAroundMessage[1] = "надо идти в универ. "
			}
			return "предмет добавлен в инвентарь: " + object
		} else {
			return "некуда класть"
		}
	} else if slices.Contains(objectsType["сумки"], object) {
		return "используйте команду 'надеть' для того, чтобы поднимать сумки"
	} else {
		return "нет такого"
	}
}

func (Player) changeRoom(room string) string {

	if rooms[room] == nil {
		return "такой комнаты не существует"
	} else if !(slices.Contains(Player1.activRoom.neibours, room)) {
		return "нет пути в " + room
	} else {
		if rooms[room].doorIsOpen {
			Player1.activRoom = rooms[room]
			return Player1.activRoom.StartMessage()
		} else {
			return "дверь закрыта"
		}
	}
}

func (Player) putOn(object string) string {

	index := slices.Index(Player1.activRoom.objectsIn, object)
	if index != -1 && slices.Contains(objectsType["сумки"], object) {
		switch object {
		case "рюкзак":
			Player1.bag = Bag{
				slots:  3,
				object: []string{},
			}
			Player1.activRoom.objectsIn[index] = Player1.activRoom.objectsIn[len(Player1.activRoom.objectsIn)-1]
			Player1.activRoom.objectsIn = Player1.activRoom.objectsIn[:len(Player1.activRoom.objectsIn)-1]
		}

	} else if slices.Contains(objectsType["расходники"], object) {
		return "используйте команду 'взять' для того, чтобы добавить расходник в инвентарь"
	} else {
		return "нет такого"
	}

	return "вы надели: " + object
}

func (Player) useObject(firstObject, secondObject string) string {

	if slices.Contains(Player1.bag.object, firstObject) {
		if slices.Contains(objectsAction[firstObject], secondObject) {
			if firstObject == "ключи" && Player1.activRoom == &Hall {
				Outside.doorIsOpen = true
				return "дверь открыта"
			} else {
				return "нет закрытой двери"
			}

		} else {
			return "не к чему применить"
		}

	} else {
		return "нет предмета в инвентаре - " + firstObject
	}
}

var Player1 Player
var Bedroom Room
var Hall Room
var Kitchen Room
var Outside Room

var rooms map[string]*Room = map[string]*Room{
	"комната": &Bedroom,
	"кухня":   &Kitchen,
	"коридор": &Hall,
	"улица":   &Outside,
	"домой":   &Hall,
}

var objectsType map[string][]string = map[string][]string{
	"расходники": {"чай", "ключи", "конспекты"},
	"сумки":      {"рюкзак"},
}

var objectsAction map[string][]string = map[string][]string{
	"чай":       {},
	"ключи":     {"дверь"},
	"конспекты": {},
}

func main() {

	initGame()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(handleCommand(scanner.Text()))
	}
}

func initGame() {
	Player1 = Player{
		bag: Bag{slots: 0,
			object: []string{}},
		activRoom: &Kitchen}
	Bedroom = CreateBedroom()
	Hall = CreateHall()
	Kitchen = CreateKitchen()
	Outside = CreateOutside()
}

func handleCommand(command string) string {

	com := strings.Split(command, " ")

	action := com[0]
	parametrs := com[1:]

	switch action {

	case "осмотреться":
		return Player1.activRoom.LookAround()

	case "идти":
		return Player1.changeRoom(parametrs[0])

	case "взять":
		return Player1.takeObject(parametrs[0])

	case "надеть":
		return Player1.putOn(parametrs[0])

	case "применить":
		return Player1.useObject(parametrs[0], parametrs[1])

	case "инвентарь":
		return "инвентарь: " + strings.Join(Player1.bag.object, ", ")

	case "перезапуск":
		initGame()
		return "игра перезапущена"
	}

	return "неизвестная команда"
}
