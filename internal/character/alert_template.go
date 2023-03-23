package character

import (
	"context"
	"errors"

	pkgalert "github.com/italorfeitosa/go-rick-n-morty/pkg/alert"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/ricknmorty"
)

type RickNMortyAlertManager interface {
	SendAlert(context.Context, error) error
}

type rickNMortyAlertManagerTemplate struct {
	alertFactory  pkgalert.AbstractFactory
	alertStrategy pkgalert.Strategy
}

func NewRickNMortyClientAlertManager(
	alertFactory pkgalert.AbstractFactory,
	alertStrategy pkgalert.Strategy,
) RickNMortyAlertManager {
	return &rickNMortyAlertManagerTemplate{alertFactory, alertStrategy}
}

func (rnm *rickNMortyAlertManagerTemplate) SendAlert(ctx context.Context, err error) error {
	var (
		data  pkgalert.Data
		alert pkgalert.Alert
	)

	data.Title = "Error on Rick N Morty Integration"
	data.Message = err.Error()

	alert = rnm.alertFactory.CreateLowAlert(data)

	if errors.Is(err, ricknmorty.ErrUnknown) {
		alert = rnm.alertFactory.CreateCriticalAlert(data)
	}

	if errors.Is(err, ricknmorty.ErrResponse) {
		alert = rnm.alertFactory.CreateHighAlert(data)
	}

	if errors.Is(err, ricknmorty.ErrNotFound) {
		alert = rnm.alertFactory.CreateLowAlert(data)
	}

	return rnm.alertStrategy.EmitAlert(ctx, alert)
}
