package imapbackend

func flagsContains(flags []string, searchFor string) bool {
	for _, flag := range flags {
		if flag == searchFor {
			return true
		}
	}
	return false
}
