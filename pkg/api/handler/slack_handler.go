package handler

import (
	"encoding/json"
	"net/http"

	"github.com/chocogem/slackbot-golang/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type SlackHandler struct {
	slackUseCase usecase.SlackUseCase
}

func NewSlackHandler(su usecase.SlackUseCase) *SlackHandler {
	return &SlackHandler{
		slackUseCase: su,
	}
}

// @Description  sendHello
// @Tags         sendHello
// @Summary      Send dummy message to slack channel
// @Produce      json
// @Success      200  {object}  SlackResponse
// @Router       /slack/hello (post)
func (cr *SlackHandler) SendHello(c *gin.Context) {
	err := cr.slackUseCase.SendHello()
	if err != nil {
		c.Error(err)
	} else {
		var response SlackResponse
		response.Success = true
		c.JSON(http.StatusOK, response)
	}

}

// @Description  sendCustomMessage
// @Tags         sendCustomMessage
// @Summary      Send custom message to slack channel
// @Param        data  body json of slack format https://api.slack.com/messaging/composing/layouts
// @Produce      json
// @Success      200  {object} json
// @Router       /slack/sendCustomMessage (post)
func (cr *SlackHandler) SendCustomMessage(c *gin.Context) {
	payload, err := c.GetRawData()
	if err != nil {
		c.Error(err)
	}
	respBody, err := cr.slackUseCase.SendCustomMessage(payload)
	if err != nil {
		c.Error(err)
	} else {
		var js map[string]interface{}
		err := json.Unmarshal(respBody, &js)
		if err != nil {
			c.Error(err)
		}
		c.JSON(http.StatusOK, js)
	}

}
