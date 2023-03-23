package alert

type DiscordFactory struct {
	username string
}

func NewDiscordAlertFactory(username string) AbstractFactory {
	return &DiscordFactory{username}
}

func (f *DiscordFactory) CreateCriticalAlert(data Data) Alert {
	return makeDiscordCriticalAlert(f.username, data)
}
func (f *DiscordFactory) CreateHighAlert(data Data) Alert {
	return makeDiscordHighAlert(f.username, data)
}
func (f *DiscordFactory) CreateMediumAlert(data Data) Alert {
	return makeDiscordMediumAlert(f.username, data)
}
func (f *DiscordFactory) CreateLowAlert(data Data) Alert {
	return makeDiscordLowAlert(f.username, data)
}
