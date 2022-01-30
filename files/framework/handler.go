package framework

type (
	Command func(Context)

	CommandStruct struct {
		command Command
		help    string
	}

	CommandMap map[string]CommandStruct

	CommandHandler struct {
		commands CommandMap
	}
)

// Function to create a new command handler
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CommandMap)}
}

// Function to find the given command. Returns false if it's not there
func (handler CommandHandler) Get(name string) (*Command, bool) {
	if cmd, present := handler.commands[name]; present {
		return &cmd.command, present
	} else {
		return nil, present
	}
}

// Function to return all registered commands
func (handler CommandHandler) GetCommands() CommandMap {
	return handler.commands
}

func (handler CommandHandler) Register(name string, command Command, helpmsg string) {
	// Massage the arguments into a "Full command"
	cmdStruct := CommandStruct{command: command, help: helpmsg}
	handler.commands[name] = cmdStruct
	if len(name) > 1 {
		handler.commands[name[:1]] = cmdStruct
	}
}

func (command CommandStruct) GetHelp() string {
	return command.help
}

/* This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	logrus.Info("Processing msg : ", m.Content, " from ", m.Author.Username)

	err := helper.CheckPrefix(m.Content)
	if err != nil {
		logrus.Info("Message is not a command")
		return // Do nothing if the message is not a command
	}

	msg := strings.TrimSpace(m.Content)
	logrus.Info(msg)
	command := strings.Split(msg[1:], " ")
	logrus.Info(command)
	if len(command) == 1 {
		switch command[0] {
		case "btc", "bitcoin":
			s.ChannelMessageSend(m.ChannelID, crypto.GetTokenByID(90))
			return
		case "eth", "ethereum":
			s.ChannelMessageSend(m.ChannelID, crypto.GetTokenByID(80))
			return
		case "help":
			s.ChannelMessageSend(m.ChannelID, entity.Help)
			return
		}
	}

	if len(command) == 2 && command[0] == "covid" {
		s.ChannelMessageSend(m.ChannelID, covid.GetCountryInfo(command[1]))
	}
}*/
