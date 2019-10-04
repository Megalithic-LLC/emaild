package submissionbackend

import (
	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Mail(from string) error {
	logger.Tracef("Submission:Session:Mail(%s)", from)
	return nil
}
