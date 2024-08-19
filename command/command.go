package command

type PrimaryCommand string

const (
	GET PrimaryCommand = "get"
	SET PrimaryCommand = "set"
)

type Command struct {
	PrimaryCommand PrimaryCommand
	SubCommand     string
	Args           []string
}
