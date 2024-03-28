package raffle

import (
	"database/sql"
	"errors"
	"fmt"
	"itacademy/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		db *sql.DB
	}
)

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) CreateRaffle(c echo.Context) error {
	raffleDetail, err := CreateRaffle(h.db)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &raffleDetail)
}

func (h *Handler) StartRaffle(c echo.Context) error {
	raffleDetail, err := StartRaffle(h.db)

	if err != nil {
		if errors.Is(err, errors.New("there is no raffle registered yet")) {
			return c.NoContent(http.StatusOK)
		}
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &raffleDetail)
}

func (h *Handler) GenerateRaffle(c echo.Context) error {
	raffleDetail, err := GenerateRaffle(h.db)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &raffleDetail)
}

func (h *Handler) CloseRaffle(c echo.Context) error {
	raffle, err := CloseRaffle(h.db)

	if err != nil {
		if errors.Is(err, errors.New("there is no raffle registered yet")) {
			return c.NoContent(http.StatusOK)
		}
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, raffle)
}

func (h *Handler) LastRaffle(c echo.Context) error {
	raffleDetail, err := LastRaffle(h.db)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, raffleDetail)
}

func (h *Handler) WinnersStats(c echo.Context) error {
	winnersStats, err := WinnersStats(h.db)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// why just inserting the winnersstas doesnt work?

	return c.JSON(http.StatusOK, domain.WinnersStatsDetail{
		Numbers:      winnersStats.Numbers,
		Rounds:       winnersStats.Rounds,
		WinnersCount: winnersStats.WinnersCount,
		Winners:      winnersStats.Winners,
	})
}

func (h *Handler) AllGambleNumbers(c echo.Context) error {
	result, err := AllGambleNumbers(h.db)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, result)
}
