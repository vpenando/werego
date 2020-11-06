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
	CommandRoles      = CommandPrefix + "roles"
	CommandCleanVotes = CommandPrefix + "cleanvotes"
	CommandHelp       = CommandPrefix + "help"
	CommandAlive      = CommandPrefix + "alive"
	CommandConnect    = CommandPrefix + "connect"
	CommandDisconnect = CommandPrefix + "disconnect"
	CommandSound      = CommandPrefix + "sound"
)

type detailledCommand struct {
	Text         string
	HelpText     string
	ExpectedArgs int
}

var commands = []detailledCommand{
	{
		Text:     CommandJoin,
		HelpText: "Joins the next game if it hasn't started.",
	},
	{
		Text:     CommandStart,
		HelpText: "Starts the game if there are enough players.",
	},
	{
		Text:     CommandStop,
		HelpText: "Stops the game.",
	},
	{
		Text:         CommandVote,
		HelpText:     "Adds one vote to a player. Usage: '!vote @player'",
		ExpectedArgs: 1,
	},
	{
		Text:         CommandKill,
		HelpText:     "Removes a player from the game. Usage: '!kill @player'",
		ExpectedArgs: 1,
	},
	{
		Text:     CommandJoined,
		HelpText: "Lists the players that have joined the game.",
	},
	{
		Text:         CommandRole,
		HelpText:     "Displays the role of a player.",
		ExpectedArgs: 1,
	},
	{
		Text:     CommandRoles,
		HelpText: "Sends the role of each player in DM.",
	},
	{
		Text:     CommandCleanVotes,
		HelpText: "Clears the votes.",
	},
	{
		Text:     CommandAlive,
		HelpText: "Returns the name of players that has not been killed.",
	},
	{
		Text:         CommandConnect,
		HelpText:     "Connects to a voice channel. Usage: '!connect channel'",
		ExpectedArgs: 1,
	},
	{
		Text:     CommandDisconnect,
		HelpText: "Disconnects from a voice channel, if connected.",
	},
}

func help() string {
	result := "Available commands:\n"
	for _, cmd := range commands {
		result += "**" + cmd.Text + " **\n *" + cmd.HelpText + "*\n"
	}
	return result
}

func checkCommand(command string, argsCount int) error {
	for _, cmd := range commands {
		if cmd.Text == command && cmd.ExpectedArgs != argsCount {
			argsText := "arg"
			if cmd.ExpectedArgs > 1 {
				argsText += "s"
			}
			return fmt.Errorf("Error: expected %d %s, got %d", cmd.ExpectedArgs, argsText, argsCount)
		}
	}
	return nil
}
