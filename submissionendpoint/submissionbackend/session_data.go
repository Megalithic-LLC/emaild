package submissionbackend

import (
	"io"
	"io/ioutil"

	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Data(r io.Reader) error {
	logger.Tracef("Submission:Session:Data()")

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return self.backend.localDelivery.Deliver(data, self.recipients)
}
