package submissionbackend

import (
	"github.com/on-prem-net/emaild/model"
)

type Session struct {
	account    *model.Account
	backend    *SubmissionBackend
	recipients []*model.Account
}
