package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	repository *Repository
}

func NewServer(repository *Repository) *Server {
	return &Server{
		repository: repository,
	}
}

func (s Server) GetUser(ctx echo.Context) error {
	email := ctx.Param("email")
	if email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid argument email")
	}
	user, err := s.repository.GetUser(email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
	}
	return ctx.JSON(http.StatusOK, user)
}

func (s Server) CreateUser(ctx echo.Context) error {
	var in User
	if err := ctx.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := s.repository.CreateUser(in)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExist) {
			return echo.NewHTTPError(http.StatusBadRequest, "user already exist")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
	}
	return ctx.JSON(http.StatusCreated, user)
}

func (s Server) UpdateUser(ctx echo.Context) error {
	email := ctx.Param("email")
	if email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid argument email")
	}
	var in User
	if err := ctx.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	in.Email = email
	user, err := s.repository.UpdateUser(in)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
	}
	return ctx.JSON(http.StatusCreated, user)
}

func (s Server) DeleteUser(ctx echo.Context) error {
	email := ctx.Param("email")
	if email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid argument email")
	}
	user, err := s.repository.DeleteUser(email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
	}
	return ctx.JSON(http.StatusOK, user)
}
