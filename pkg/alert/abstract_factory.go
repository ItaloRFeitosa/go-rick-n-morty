package alert

type AbstractFactory interface {
	CreateCriticalAlert(data Data) Alert
	CreateHighAlert(data Data) Alert
	CreateMediumAlert(data Data) Alert
	CreateLowAlert(data Data) Alert
}
