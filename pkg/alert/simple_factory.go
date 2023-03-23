package alert

type SimpleFactory struct{}

func NewSimpleAlertFactory() AbstractFactory {
	return SimpleFactory{}
}

func (SimpleFactory) CreateCriticalAlert(data Data) Alert {
	return makeSimpleCriticalAlert(data)
}
func (SimpleFactory) CreateHighAlert(data Data) Alert {
	return makeSimpleHighAlert(data)
}
func (SimpleFactory) CreateMediumAlert(data Data) Alert {
	return makeSimpleMediumAlert(data)
}
func (SimpleFactory) CreateLowAlert(data Data) Alert {
	return makeSimpleLowAlert(data)
}
