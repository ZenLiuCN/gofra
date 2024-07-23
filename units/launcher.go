package units

import (
	"context"
	"flag"
	cfg "github.com/ZenLiuCN/gofra/conf"
	"github.com/golang/glog"
	"os"
	"path/filepath"
	"sync/atomic"
)

type (
	GracefulShutdown interface {
		Shutdown(ctx context.Context) (err error) //shutdown graceful
	}
	// ReloadableServer is an application entry server. which can reload without recreate.
	ReloadableServer interface {
		Reload(config cfg.Config) (reloading bool, err error) // invoke when server is reloading
		Launch(config cfg.Config) (reloading bool, err error) // invoke when server is launch first time
	}
)
type (
	Server interface {
		GracefulShutdown
		Configure(config cfg.Config) (srv Server, err error) //Configure and returns new Server
		Run() (err error)                                    //Launch until Server is shutdown
	}
	BaseReloadable struct {
		Srv      Server
		reloaded atomic.Bool
	}
)

// IsReloading current is reloading status
func (b *BaseReloadable) IsReloading() bool {
	return b.reloaded.Load()
}

// DoneReload config ending status not a reloading
func (b *BaseReloadable) DoneReload() {
	b.reloaded.Store(false)
}

// WillReload config ending status is a reloading
func (b *BaseReloadable) WillReload() {
	b.reloaded.Store(true)
}
func (b *BaseReloadable) Reload(config cfg.Config) (reloading bool, err error) {
	b.Srv, err = b.Srv.Configure(config)
	if err != nil {
		return
	}
	err = b.Srv.Run()
	reloading = b.reloaded.Load()
	return
}

func (b *BaseReloadable) Launch(config cfg.Config) (reloading bool, err error) {
	b.Srv, err = b.Srv.Configure(config)
	if err != nil {
		return
	}
	err = b.Srv.Run()
	reloading = b.reloaded.Load()
	return
}

var (
	confFile string
)

// GLogLaunch entry of application with [github.com/golang/glog] as logger
func GLogLaunch(srv ReloadableServer) {
	flag.StringVar(&confFile, "conf", "app.conf", "configuration file")
	os.Args = append(os.Args, "-alsologtostderr", "-log_dir=logs")
	flag.Parse()
	cfg.Initialize(confFile)
	defer glog.Flush()
	var reload bool
	var err error
	var name = filepath.Base(os.Args[0])
reload:
	if reload {
		glog.Infof("reload %s server", name)
		reload, err = srv.Reload(cfg.GetConfig())
		if err != nil {
			glog.Error(err)
		}
	} else {
		glog.Infof("launch %s server", name)
		reload, err = srv.Launch(cfg.GetConfig())
		if err != nil {
			glog.Error(err)
		}
	}
	if reload {
		goto reload
	}
	glog.Infof("shutdown %s server", name)
}
