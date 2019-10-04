package submissionbackend

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
)

func (self *SubmissionBackend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	logger.Tracef("SubmissionBackend:AnonymousLogin()")
	return nil, smtp.ErrAuthRequired
}
