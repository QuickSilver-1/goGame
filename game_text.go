package main

func CreateBedroom() Room {

	startMessageBedroom := "ты в своей комнате. "
	objectsInBedroom := []string{"ключи", "конспекты", "рюкзак"}
	objectsInBedroom1 := []string{"ключи", "конспекты"}
	objectsInBedroom2 := []string{"рюкзак"}
	orderPlaceInBedroom := []string{"на столе", "на стуле"}
	neiboursBedroom := []string{"коридор"}
	whereObjectsInBedroom := map[string][]string{
		orderPlaceInBedroom[0]: objectsInBedroom1,
		orderPlaceInBedroom[1]: objectsInBedroom2}
	lookAroundMessageBedroom := []string{"", ""}
	return Room{
		doorIsOpen:        true,
		startMessage:      startMessageBedroom,
		objectsIn:         objectsInBedroom,
		lookAroundMessage: lookAroundMessageBedroom,
		orderPlaceIn:      orderPlaceInBedroom,
		neibours:          neiboursBedroom,
		whereObjectsIn:    whereObjectsInBedroom}
}

func CreateHall() Room {

	startMessageHall := "ничего интересного. "
	objectsInHall := []string{}
	orderPlaceInHall := []string{}
	neiboursHall := []string{"кухня", "комната", "улица"}
	whereObjectsInHall := map[string][]string{}
	lookAroundMessageHall := []string{"ничего интересного. ", ""}

	return Room{
		doorIsOpen:        true,
		startMessage:      startMessageHall,
		objectsIn:         objectsInHall,
		lookAroundMessage: lookAroundMessageHall,
		orderPlaceIn:      orderPlaceInHall,
		neibours:          neiboursHall,
		whereObjectsIn:    whereObjectsInHall}
}

func CreateKitchen() Room {

	startMessageKitchen := "кухня, ничего интересного. "
	objectsInKitchen := []string{"чай"}
	orderPlaceInKitchen := []string{"на столе"}
	neiboursKitchen := []string{"коридор"}
	whereObjectsInKitchen := map[string][]string{
		orderPlaceInKitchen[0]: objectsInKitchen}
	lookAroundMessageKitchen := []string{"ты находишься на кухне, ", "надо собрать рюкзак и идти в универ. "}

	return Room{
		doorIsOpen:        true,
		startMessage:      startMessageKitchen,
		objectsIn:         objectsInKitchen,
		lookAroundMessage: lookAroundMessageKitchen,
		orderPlaceIn:      orderPlaceInKitchen,
		neibours:          neiboursKitchen,
		whereObjectsIn:    whereObjectsInKitchen}
}

func CreateOutside() Room {

	startMessageOutside := "на улице весна. "
	objectsInOutside := []string{}
	orderPlaceInOutside := []string{}
	neiboursOutside := []string{"домой"}
	whereObjectsInOutside := map[string][]string{}
	lookAroundMessageOutside := []string{"на улице весна. ", "нужно идти в универ. "}

	return Room{
		doorIsOpen:        false,
		startMessage:      startMessageOutside,
		objectsIn:         objectsInOutside,
		lookAroundMessage: lookAroundMessageOutside,
		orderPlaceIn:      orderPlaceInOutside,
		neibours:          neiboursOutside,
		whereObjectsIn:    whereObjectsInOutside}
}
