package commands

const (
	Exit = "exit"
	Echo = "echo"
	Type = "type"
)

var BuiltIns = map[string]struct{}{Exit: {}, Echo: {}, Type: {}}

func HandleCommand(input string) {
	command, args := preprocess(input)
	switch command {
	case Exit:
		handleExit(args)
	case Echo:
		handleEcho(args)
	case Type:
		handleType(args)
	default:
		handleUnknown(command)
	}
}
