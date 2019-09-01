package imapbackend

type User struct {
	backend  *ImapBackend
	username string
}

func (self *User) Username() string {
	return self.username
}
