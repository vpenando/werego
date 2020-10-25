package werego

import "fmt"

// Discord commands
const (
	CommandPrefix     = "!"
	CommandJoin       = CommandPrefix + "join"
	CommandStart      = CommandPrefix + "start"
	CommandStop       = CommandPrefix + "stop"
	CommandVote       = CommandPrefix + "vote"
	CommandVotes      = CommandPrefix + "votes"
	CommandKill       = CommandPrefix + "kill"
	CommandJoined     = CommandPrefix + "joined"
	CommandRole       = CommandPrefix + "role"
	CommandCleanVotes = CommandPrefix + "cleanvotes"
	CommandHelp       = CommandPrefix + "help"
	CommandAlive      = CommandPrefix + "alive"
)

type detailledCommand struct {
	Text         string
	HelpText     string
	ExpectedArgs int
}

var commands = []detailledCommand{
	{
		Text:         CommandJoin,
		HelpText:     "Joins the next game if it hasn't started.",
		ExpectedArgs: 0,
	},
	{
		Text:         CommandStart,
		HelpText:     "Starts the game if there are enough players.",
		ExpectedArgs: 0,
	},
	{
		Text:         CommandStop,
		HelpText:     "Stops the game.",
		ExpectedArgs: 0,
	},
	{
		Text:         CommandVote,
		HelpText:     "Add one vote to a player. Usage: '!vote @player'",
		ExpectedArgs: 1,
	},
	{
		Text:         CommandKill,
		HelpText:     "Removes a player from the game. Usage: '!kill @player'",
		ExpectedArgs: 1,
	},
	{
		Text:         CommandJoined,
		HelpText:     "Lists the players that have joined the game.",
		ExpectedArgs: 1,
	},
	{
		Text:         CommandRole,
		HelpText:     "Displays the role of a player.",
		ExpectedArgs: 1,
	},
	{
		Text:     CommandCleanVotes,
		HelpText: "Clears the votes.",
	},
}

func help() string {
	result := "Available commands:\n"
	for _, cmd := range commands {
		result += "**  " + cmd.Text + " **\n  *" + cmd.HelpText + "*\n"
	}
	return result
}

func checkCommand(command string, argsCount int) error {
	for _, cmd := range commands {
		if cmd.Text == command && cmd.ExpectedArgs != argsCount {
			return fmt.Errorf("Error: expected %d args, got %d", cmd.ExpectedArgs, argsCount)
		}
	}
	return nil
}
