package alert

import (
	"context"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type httpStrategy struct {
	resty *resty.Client
	url   string
}

func NewHttpStrategy(url string) Strategy {
	return &httpStrategy{resty.New(), url}
}

func (h *httpStrategy) EmitAlert(ctx context.Context, alert Alert) error {
	res, err := h.resty.R().SetContext(ctx).SetBody(alert).Post(h.url)
	if err != nil {
		return err
	}

	if res.IsError() {
		err = fmt.Errorf("error on notify to http strategy; status code: %d", res.StatusCode())
		log.Println(err)
		return err
	}

	return nil
}
