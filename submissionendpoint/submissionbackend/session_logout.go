package submissionbackend

import (
	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Logout() error {
	logger.Tracef("Submission:Session:Logout()")
	return nil
}
