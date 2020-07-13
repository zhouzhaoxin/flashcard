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
	row := db.QueryRow("select id, front, back, known from cards where id=?", id)
	handleErr(row.Scan(&card.ID, &card.Front, &card.Back, &card.Known))
	return &card
}

func editCardKnown(card *Card) int64 {
	stmt, err := db.Prepare("UPDATE cards SET known=? WHERE id=?")

	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		handleErr(stmt.Close())
	}()
	rs, err := stmt.Exec(card.Known, card.ID)
	handleErr(err)
	ra, err := rs.RowsAffected()
	handleErr(err)
	if err != nil {
		log.Fatalln(err)
	}
	return ra
}

func editCard(card *Card) int64 {
	stmt, err := db.Prepare("UPDATE cards SET front=?, back=? WHERE id=?")

	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		handleErr(stmt.Close())
	}()
	rs, err := stmt.Exec(card.Front, card.Back, card.ID)
	handleErr(err)
	ra, err := rs.RowsAffected()
	handleErr(err)
	if err != nil {
		log.Fatalln(err)
	}
	return ra
}

func knownCard(id int) int64 {
	stmt, err := db.Prepare("UPDATE cards SET known=1 WHERE id=?")

	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		handleErr(stmt.Close())
	}()
	rs, err := stmt.Exec(id)
	handleErr(err)
	ra, err := rs.RowsAffected()
	handleErr(err)
	if err != nil {
		log.Fatalln(err)
	}
	knownIDs = append(knownIDs[:current], knownIDs[current+1:]...)
	current--
	return ra
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

func getCards() []*Card {
	rows, err := db.Query("select id, front, back, known from cards")
	defer func() {
		err := rows.Close()
		handleErr(err)
	}()
	handleErr(err)
	var cards []*Card
	for rows.Next() {
		var card Card
		err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.Known)
		handleErr(err)
		cards = append(cards, &card)
	}
	return cards
}

func getCardsByKnown(known int) []*Card {
	rows, err := db.Query("select id, front, back, known from cards where known=?", known)
	defer func() {
		err := rows.Close()
		handleErr(err)
	}()
	handleErr(err)
	var cards []*Card
	for rows.Next() {
		var card Card
		err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.Known)
		handleErr(err)
		cards = append(cards, &card)
	}
	return cards
}

func addCard(card *Card) {
	exec, err := db.Exec("insert into cards(front, back, known) values(?, ?, ?)",
		card.Front, card.Back, 0)
	handleErr(err)
	_, err = exec.LastInsertId()
	handleErr(err)
	return
}

func deleteCard(id int) int64 {
	rs, err := db.Exec("DELETE FROM cards WHERE id=?", id)
	handleErr(err)
	ra, err := rs.RowsAffected()
	handleErr(err)
	return ra
}
