package handler

import (
	"fmt"
	"net/http"

	"github.com/JuDyas/GolangPractice/try_jenkins/internal/service"
	"github.com/labstack/echo/v4"
)

type WordCountRequest struct {
	Text string `json:"text"`
}

func CountWords(c echo.Context) error {
	var req WordCountRequest
	fmt.Println(req)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	svc := service.NewWordCountService()
	wordCount := svc.CountWords(req.Text)
	//Делаем коммиты двааа

	//Коммит в фюче?
	return c.JSON(http.StatusOK, map[string]int{"word_count": wordCount})
}
