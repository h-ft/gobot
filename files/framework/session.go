package framework

type (
	Session struct {
		guildId, ChannelID string
	}

	SessionManager struct {
		sessions map[string]*Session
	}

	JoinProperties struct {
		Muted    bool
		Deafened bool
	}
)

func newSession(guildId, channelId string) *Session {
	session := new(Session)
	session.guildId = guildId
	session.ChannelID = channelId
	return session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{make(map[string]*Session)}
}

func (manager SessionManager) GetByGuild(guildId string) *Session {
	for _, sess := range manager.sessions {
		if sess.guildId == guildId {
			return sess
		}
	}
	return nil
}

func (manager SessionManager) GetByChannel(channelId string) (*Session, bool) {
	sess, found := manager.sessions[channelId]
	return sess, found
}
