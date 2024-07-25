package units

import cfg "github.com/ZenLiuCN/gofra/conf"

// Reloadable an component that support inner reloading.
type Reloadable interface {
	Reload(conf cfg.Config) error
}
