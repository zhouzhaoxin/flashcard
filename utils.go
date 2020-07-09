package main

import (
	"golang.org/x/exp/rand"
	"log"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

var knownIDs []int
var current int32

func generateKnownIDs() {
	var concatID string
	current = -1
	knownIDs = knownIDs[0:0]
	row := db.QueryRow("select group_concat(id) from cards where known=0")
	handleErr(row.Scan(&concatID))
	for _, id := range strings.Split(concatID, ",") {
		id, err := strconv.Atoi(id)
		handleErr(err)
		knownIDs = append(knownIDs, id)
	}
	rand.Shuffle(len(knownIDs), func(i, j int) { knownIDs[i], knownIDs[j] = knownIDs[j], knownIDs[i] })
}

func getCardByID(id int) *Card {
	var card Card
	row := db.QueryRow("select id, type, front, back, known from cards where known=0 and id=?", id)
	handleErr(row.Scan(&card.ID, &card.Type, &card.Front, &card.Back, &card.Known))
	return &card
}

func getNextCard() *Card {
	var card Card
	current++
	if int(current) >= len(knownIDs) {
		current--
		card.Back = "没有啦"
		card.State = 1
		card.Front = "没有啦"
	} else {
		return getCardByID(knownIDs[current])
	}
	return &card
}

func getPrevCard() *Card {
	var card Card
	current--
	if current < 0 {
		current++
		card.Back = "没有啦"
		card.State = 1
		card.Front = "没有啦"
	} else {
		return getCardByID(knownIDs[current])
	}

	return &card
}
