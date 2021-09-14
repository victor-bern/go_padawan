package controllers

import (
	"fmt"
	"go_padawan/src/database"
	"go_padawan/src/models"
	"go_padawan/src/services"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ExchangeValue(c *fiber.Ctx) error {
	db := database.GetDatabase()
	var exc *models.Exchange
	var excFinded *models.Exchange

	from := c.Params("from")
	to := c.Params("to")

	if !services.ValidateCurrency(from, to) {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"Error": "Send a valid currency",
			"ValidCurrencies": []string{
				"BRL",
				"USD",
				"EUR",
				"BTC",
			},
		})
	}

	if strings.EqualFold(from, to) {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"Error": "Currencies can't be same",
		})
	}

	amount, err := strconv.ParseFloat(c.Params("amount"), 64)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"error": "Format invalid",
		})
	}

	rate, err := strconv.ParseFloat(c.Params("rate"), 64)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"error": "Format invalid",
		})
	}
	exc = &models.Exchange{
		From:   from,
		To:     to,
		Amount: amount,
		Rate:   rate,
	}

	db.Where("`from` = ? AND `to` = ? AND `amount` = ? AND `rate` = ? AND RESULT > 0", exc.From, exc.To, exc.Amount, exc.Rate).First(&excFinded)

	if len(excFinded.From) > 0 {
		result := excFinded.Result
		symbol, _ := services.GetCurrency(excFinded.To)

		c.SendStatus(200)
		return c.JSON(fiber.Map{
			"valorConvertido": result,
			"simboloMoeda":    symbol,
		})
	}

	result := amount * rate
	exc.Result = fmt.Sprintf("%.2f", result)

	db.Create(&exc)

	symbol, _ := services.GetCurrency(exc.To)

	c.SendStatus(200)
	return c.JSON(fiber.Map{
		"valorConvertido": fmt.Sprintf("%.2f", result),
		"simboloMoeda":    symbol,
	})
}
