package loader

import (
	//	"fmt"
	"github.com/csgura/di"
	conf "github.com/go-akka/configuration"
	"os"
)

/* DI (github.com/csgura/di) module helper */
/* this module is basic binder */

//Loader is module helper
type Loader struct {
	filePath string
	cfg      *conf.Config
	injector di.Injector
	impls    *di.Implements
}

//GetConfig is return *conf.Config(using hcon)
func (l *Loader) GetConfig() *conf.Config {
	return l.cfg
}

//GetInstance is find DI instance. (careful used this method. vPtr is always ref)
func (l *Loader) GetInstance(vPtr interface{}) interface{} {

	return l.injector.GetInstance(vPtr)
}

//IsInitModule is already load modules check. return true => init ok
func (l *Loader) IsInitModule() bool {
	return l.injector != nil
}

//LoadMoudleFromFileConfig is create new injector & bind impls & load config
func LoadMoudleFromFileConfig(filePath string, impls *di.Implements) *Loader {

	cfg := NewConfigFromFile(filePath)
	if cfg == nil {
		//Error Can't Load Configure File
		//TODO -> On error handler module need.
		return nil
	}

	if impls == nil {
		//impls is none nil
		return nil
	}

	//Open Anonymous func
	impls.AddBind(func(binder *di.Binder) {
		binder.BindSingleton((*conf.Config)(nil), cfg)
	})

	mods := cfg.GetStringList("modules")
	if len(mods) <= 0 {
		//Can't load modules
		return nil
	}

	retInjector := impls.NewInjector(mods)

	newLoader := &Loader{
		filePath: filePath,
		cfg:      cfg,
		injector: retInjector,
		impls:    impls,
	}

	return newLoader
}

//NewConfigFromFile is load from file config
func NewConfigFromFile(filePath string) *conf.Config {

	if ok := FileExist(filePath); ok == false {
		return nil
	}

	newConf := conf.LoadConfig(filePath)

	return newConf
}

//FileExist is check exist file
func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
