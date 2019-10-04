package submissionbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
)

type Session struct {
	account    *model.Account
	backend    *SubmissionBackend
	recipients []*model.Account
}
