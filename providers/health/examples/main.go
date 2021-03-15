// Author: recallsong
// Email: songruiguo@qq.com

package main

import (
	"fmt"
	"os"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/health"
	_ "github.com/erda-project/erda-infra/providers/httpserver"
)

type define struct{}

func (d *define) Service() []string      { return []string{"hello"} }
func (d *define) Dependencies() []string { return []string{"health"} }
func (d *define) Description() string    { return "hello for example" }
func (d *define) Creator() servicehub.Creator {
	return func() servicehub.Provider {
		return &provider{}
	}
}

type provider struct {
	L logs.Logger
}

func (p *provider) Init(ctx servicehub.Context) error {
	h := ctx.Service("health").(health.Interface)
	h.Register(p.HealthCheck)
	return nil
}

func (p *provider) HealthCheck() error {
	return fmt.Errorf("error message")
}

func init() {
	servicehub.RegisterProvider("hello", &define{})
}

func main() {
	hub := servicehub.New()
	hub.Run("examples", "", os.Args...)
}