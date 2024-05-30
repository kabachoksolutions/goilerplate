package discograph

type logMessage struct {
	Message string
	Err     error
}

type logRequest struct {
	isCritical bool
	message    *logMessage
	assignees  []string
	channels   []string
}

type logRequestBuilder struct {
	req *logRequest
	dl  *discordLogger
}

func (dl *discordLogger) Request() *logRequestBuilder {
	return &logRequestBuilder{
		req: &logRequest{},
		dl:  dl,
	}
}

func (b *logRequestBuilder) Critical() *logRequestBuilder {
	b.req.isCritical = true
	return b
}

func (b *logRequestBuilder) Message(msg string) *logRequestBuilder {
	if b.req.message == nil {
		b.req.message = &logMessage{}
	}
	b.req.message.Message = msg
	return b
}

func (b *logRequestBuilder) Error(err error) *logRequestBuilder {
	if b.req.message == nil {
		b.req.message = &logMessage{}
	}
	b.req.message.Err = err
	return b
}

func (b *logRequestBuilder) Assignees(assignees ...string) *logRequestBuilder {
	b.req.assignees = assignees
	return b
}

func (b *logRequestBuilder) Channels(channels ...string) *logRequestBuilder {
	b.req.channels = channels
	return b
}

func (b *logRequestBuilder) Send() error {
	return b.dl.sendLogRequest(b.req)
}
