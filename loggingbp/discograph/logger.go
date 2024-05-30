package discograph

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

const defaultTimeFormat = "15:04:05 02.01.2006"

type DiscordLogger interface {
	Request() *logRequestBuilder
	Close() error
}

type discordLogger struct {
	sess       *discordgo.Session
	timeFormat string
}

type discordLoggerOpt func(*discordLogger)

func WithCustomTime(timeFormat string) discordLoggerOpt {
	return func(dl *discordLogger) {
		dl.timeFormat = timeFormat
	}
}

func NewDiscordLogger(apiKey string, opts ...discordLoggerOpt) (DiscordLogger, error) {
	session, err := discordgo.New("Bot " + apiKey)
	if err != nil {
		return nil, err
	}

	if err := session.Open(); err != nil {
		return nil, err
	}

	dl := &discordLogger{sess: session}
	dl.timeFormat = defaultTimeFormat

	for _, opt := range opts {
		opt(dl)
	}

	return dl, nil
}

func (dl *discordLogger) Close() error {
	if dl.sess != nil {
		return dl.sess.Close()
	}
	return nil
}

func (dl *discordLogger) sendLogRequest(msg *logRequest) error {
	var content string
	if msg.isCritical {
		content += "👮 Urgent intervention required 👮\n\n"
	}

	if len(msg.assignees) > 0 {
		content += "Assigned to: "
		for _, assignee := range msg.assignees {
			content += fmt.Sprintf("%s ", assignee)
		}
		content += "\n\n"
	}

	if msg.message != nil {
		if msg.message.Err != nil {
			content += fmt.Sprintf("Message: %s\nError: %v\n", msg.message.Message, msg.message.Err)
		} else {
			content += fmt.Sprintf("Message: %s\n", msg.message.Message)
		}
	}

	content += fmt.Sprintf("\nTime: %s", time.Now().Format(dl.timeFormat))
	for _, channel := range msg.channels {
		_, err := dl.sess.ChannelMessageSend(channel, content)
		if err != nil {
			return err
		}
	}

	return nil
}
