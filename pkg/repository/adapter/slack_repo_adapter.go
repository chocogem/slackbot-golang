package adapter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/chocogem/slackbot-golang/pkg/config"
	"github.com/chocogem/slackbot-golang/pkg/domain"
	"github.com/chocogem/slackbot-golang/pkg/repository"
    "github.com/chocogem/slackbot-golang/pkg/driver/client"
	"github.com/slack-go/slack"
)

type slackAdapter struct {
	slackConfig domain.SlackConfig
    slackClient client.ApiClient
}

func NewSlackRepository(cfg config.Config,slackClient client.ApiClient) (repository.SlackRepository, error) {
	return &slackAdapter{
		slackConfig: domain.SlackConfig{
                     BotAuthToken: cfg.SlackAuthToken,
                     ChannelId: cfg.SlackChannelId,
                     SlackMessageUrl:  cfg.SlackMessageUrl,
                    },
        slackClient: slackClient,}, nil

}

func (c *slackAdapter) SendHello() error {
	var param slack.PostMessageParameters
	client := slack.New(c.slackConfig.BotAuthToken, slack.OptionDebug(true))

	attachment := slack.Attachment{
		Pretext: "Slack Bot Message",
		Text:    "Hello This is dummy message",
		Color:   "#AF1F45",
		Fields: []slack.AttachmentField{
			{
				Title: "This Time is",
				Value: time.Now().String(),
			},
		},
	}
	jparam := []byte(`{"username":"Custom name bot"}`)
	if err := json.Unmarshal(jparam, &param); err != nil {
		panic(err)
	}
	slackOptions := slack.MsgOptionPostMessageParameters(
        slack.PostMessageParameters{
            User:"Custom Username",
        },
    )
	channel, timestamp, err := client.PostMessage(
		c.slackConfig.ChannelId,
		slack.MsgOptionAttachments(attachment),
        slackOptions,
	)

	if err != nil {
		return err
	}
	fmt.Printf("Message sent to %s at %s", channel, timestamp)

	return nil
}

func (c *slackAdapter) SendCustomMessage(payload []byte)([]byte,error) {
	respBody,err := c.slackClient.PostCustom(c.slackConfig.SlackMessageUrl,payload,c.slackConfig.BotAuthToken)
    if(err!=nil){
        return nil,err
    }
	return respBody,nil
}