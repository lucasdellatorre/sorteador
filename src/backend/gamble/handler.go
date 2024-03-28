package gamble

import (
	"database/sql"
	"errors"
	"itacademy/domain"
	"itacademy/util"
	"net/http"

	"github.com/go-playground/validator/v10"
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

func (h *Handler) Create(c echo.Context) error {
	validate := validator.New()
	gamble := new(domain.Gamble)

	if err := validate.RegisterValidation("cpf", util.ValidateCPF); err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("bad cpf"))
	}
	if err := validate.RegisterValidation("numbers", util.ValidateNumbers); err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("bad gambling numbers"))
	}
	if err := c.Bind(gamble); err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("bad gambling numbers"))
	}
	if err := validate.Struct(gamble); err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("bad request"))
	}

	id, err := Create(h.db, gamble)

	if err != nil {
		if errors.Is(err, errors.New("raffle not active")) {
			c.NoContent(http.StatusUnauthorized)
		}
		return c.NoContent(http.StatusInternalServerError)
	}

	result := &domain.GambleDetail{GambleId: id, Name: gamble.Name, Cpf: gamble.Cpf, Numbers: gamble.Numbers, RaffleId: gamble.RaffleID}

	return c.JSON(http.StatusCreated, result)
}

func (h *Handler) List(c echo.Context) error {
	result, err := List(h.db)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, result)
}
