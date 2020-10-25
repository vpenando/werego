package werego

import "fmt"

// Discord commands
const (
	CommandPrefix = "!"
	CommandJoin   = CommandPrefix + "join"
	CommandStart  = CommandPrefix + "start"
	CommandStop   = CommandPrefix + "stop"
	CommandVote   = CommandPrefix + "vote"
	CommandVotes  = CommandPrefix + "votes"
	CommandKill   = CommandPrefix + "kill"
	CommandJoined = CommandPrefix + "joined"
	CommandRole   = CommandPrefix + "role"
)

var commandWithExpectedArgs = map[string]int{
	CommandJoin:   0,
	CommandStart:  0,
	CommandStop:   0,
	CommandVote:   1,
	CommandKill:   1,
	CommandJoined: 0,
}

func checkCommand(command string, argsCount int) error {
	for cmd, args := range commandWithExpectedArgs {
		if cmd == command && args != argsCount {
			return fmt.Errorf("Error: expected %d args, got %d", args, argsCount)
		}
	}
	return nil
}
