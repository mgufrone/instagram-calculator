package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mgufrone/insta-calculator/server/helpers"
	"github.com/mgufrone/insta-calculator/server/service/instagram"
	"net/http"
	"regexp"
)

type Instagram struct {
	ig *instagram.Service
}

func NewHandler(ig *instagram.Service) *Instagram {
	return &Instagram{ig: ig}
}

type Response struct {
	Engagement float64 `json:"engagement"`
	Followers  int64   `json:"followers"`
	Username   string  `json:"username"`
}

type Request struct {
	Username string `json:"username" form:"username"`
}

func (i *Instagram) GetProfile(ctx *gin.Context) {
	var (
		err      error
		errCode  int
		eng      *instagram.Engagement
		data     *Response
		req      *Request
		username string
	)
	errCode = http.StatusInternalServerError
	if err = helpers.TryOrError(func() error {
		return ctx.BindJSON(&req)
	}, func() (err error) {
		username = req.Username
		fmt.Println("requesting profile of ", username)
		validUsername := regexp.MustCompile("^([a-zA-Z0-9_.]{3,30})$")
		if !validUsername.MatchString(username) {
			err = errors.New("invalid username")
			errCode = http.StatusBadRequest
		}
		return
	}, func() (err error) {
		eng, err = i.ig.RetrieveProfile(username)
		return
	}, func() (err error) {
		data = &Response{
			Engagement: eng.Rate(),
			Followers:  eng.Followers,
			Username:   username,
		}
		return
	}); err != nil {
		err1 := ctx.Error(err)
		ctx.JSON(errCode, err1)
		return
	}
	ctx.JSON(http.StatusOK, data)
}
