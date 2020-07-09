package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 增
func cardsAddIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "cards-add.html", nil)
}

func cardsAddHandler(c *gin.Context) {
	var card Card
	handleErr(c.BindJSON(&card))
	if card.Front == "" || card.Back == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{"data": "参数错误"})
		return
	}
	addCard(&card)
	c.JSON(http.StatusOK, gin.H{"data": "成功"})
}

// 删
func cardsDeleteHandler(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.Atoi(id)
	handleErr(err)
	ra := deleteCard(i)
	c.JSON(http.StatusOK, gin.H{"data": ra})
}

// 改
func cardsEditIndex(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.Atoi(id)
	handleErr(err)
	card := getCardByID(i)
	c.HTML(http.StatusOK, "cards-edit.html", gin.H{"id": card.ID, "front": card.Front, "back": card.Back})
}

func cardsEditHandler(c *gin.Context) {
	var card Card
	handleErr(c.BindJSON(&card))
	ra := editCard(&card)
	c.JSON(http.StatusOK, gin.H{"data": ra})
}

// 查
func cardsIndex(c *gin.Context) {
	cards := getCards()
	c.HTML(http.StatusOK, "cards.html", gin.H{"cards": cards})
}

func cards(c *gin.Context) {
	cards := getCards()
	c.JSON(http.StatusOK, gin.H{"data": cards})
}

func remember(c *gin.Context) {
	generateKnownIDs()
	card := getNextCard()
	c.HTML(http.StatusOK, "remember.html", gin.H{
		"card_id":    card.ID,
		"card_front": card.Front,
		"card_back":  card.Back,
	})
}

func rememberNext(c *gin.Context) {
	card := getNextCard()
	c.JSON(http.StatusOK, gin.H{
		"card_id":    card.ID,
		"card_front": card.Front,
		"card_back":  card.Back,
		"card_state": card.State,
	})
}

func rememberPrev(c *gin.Context) {
	card := getPrevCard()
	c.JSON(http.StatusOK, gin.H{
		"card_id":    card.ID,
		"card_front": card.Front,
		"card_back":  card.Back,
		"card_state": card.State,
	})
}

func known(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.Atoi(id)
	handleErr(err)
	knownCard(i)
	card := getNextCard()
	c.JSON(http.StatusOK, gin.H{
		"card_id":    card.ID,
		"card_front": card.Front,
		"card_back":  card.Back,
		"card_state": card.State,
	})
}
