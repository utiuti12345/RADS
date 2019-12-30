package slack

import "github.com/ashwanthkumar/slack-go-webhook"

type Slack struct {
	WebhookUrl string
	Channel string
	UserName string
}

func NewSlack(webhookUrl string,channel string,userName string) Slack {
	return Slack{
		WebhookUrl: webhookUrl,
		Channel:    channel,
		UserName:   userName,
	}
}

//const (
//	WEBHOOKURL = "https://hooks.slack.com/services/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"
//	CHANNEL    = "develop"
//	USERNAME   = "sample"
//)

func (s Slack) PostMessage(title string,message string)(errs []error) {
	field1 := slack.Field{Title: title, Value: message}
	//field2 := slack.Field{Title: "AnythingKey", Value: "AnythingValue"}

	attachment := slack.Attachment{}
	//attachment.AddField(field1).AddField(field2)
	attachment.AddField(field1)
	color := "good"
	attachment.Color = &color
	payload := slack.Payload {
		Username:    s.UserName,
		Channel:     s.Channel,
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(s.WebhookUrl, "", payload)
	if err != nil {
		return err
	}

	return nil
}
