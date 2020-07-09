package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func addVocab(c *gin.Context) {
	var card Card
	handleErr(c.BindJSON(&card))
	if card.Front == "" || card.Back == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{"data": "参数错误"})
		return
	}
	exec, err := db.Exec("insert into cards(type, front, back, known) values(?, ?, ?, ?)", 1,
		card.Front, card.Back, 0)
	handleErr(err)
	_, err = exec.LastInsertId()
	handleErr(err)
	c.JSON(http.StatusOK, gin.H{"data": "成功"})
}

func addCode(c *gin.Context) {
	var card Card
	handleErr(c.BindJSON(&card))
	if card.Front == "" || card.Back == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{"data": "参数错误"})
		return
	}
	exec, err := db.Exec("insert into cards(type, front, back, known) values(?, ?, ?, ?)", 2,
		card.Front, card.Back, 0)
	handleErr(err)
	_, err = exec.LastInsertId()
	handleErr(err)
	c.JSON(http.StatusOK, gin.H{"data": "成功"})
}

func cards(c *gin.Context) {
	rows, err := db.Query("select id, type, front, back, known from cards")
	defer func() {
		err := rows.Close()
		handleErr(err)
	}()
	handleErr(err)
	var cards []Card
	for rows.Next() {
		var card Card
		err := rows.Scan(&card.ID, &card.Type, &card.Front, &card.Back, &card.Known)
		handleErr(err)
		cards = append(cards, card)
	}
	c.JSON(http.StatusOK, gin.H{"data": cards})
}

func rememberNext(c *gin.Context) {
	card := getNextCard()
	c.JSON(http.StatusOK, gin.H{
		"card_id":    card.ID,
		"card_type":  card.Type,
		"card_front": card.Front,
		"card_back":  card.Back,
		"card_state":  card.State,
	})
}

func rememberPrev(c *gin.Context) {
	card := getPrevCard()
	c.JSON(http.StatusOK, gin.H{
		"card_id":    card.ID,
		"card_type":  card.Type,
		"card_front": card.Front,
		"card_back":  card.Back,
		"card_state":  card.State,
	})
}

func remember(c *gin.Context) {
	generateKnownIDs()
	card := getNextCard()
	c.HTML(http.StatusOK, "remember.html", gin.H{
		"card_id":    card.ID,
		"card_type":  card.Type,
		"card_front": card.Front,
		"card_back":  card.Back,
	})
}

func addCodeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "add-code.html", nil)
}

func addVocabIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "add-vocab.html", nil)
}
