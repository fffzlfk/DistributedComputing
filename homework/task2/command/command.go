package command

type commandID int

const (
	CMD_ADD_BOOK commandID = iota
	CMD_QUERY_BY_ID
	CMD_QUERY_BY_NAME
	CMD_DELETE
)

type command struct {
	id   commandID
	args []string
}
