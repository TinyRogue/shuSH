package commands

const (
	Exit = "exit"
	Echo = "echo"
	Type = "type"
	PWD  = "pwd"
)

type Command struct {
	primary     string
	subcommands []string
	raw         string
}

func HandleCommand(input string) {
	command := parse(input)
	switch command.primary {
	case Exit:
		command.handleExit()
	case Echo:
		command.handleEcho()
	case Type:
		command.handleType()
	case PWD:
		command.handlePrintWorkingDirectory()
	default:
		command.handleExternal()
	}
}
