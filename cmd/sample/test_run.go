package main

import (
	"fmt"
	"github.com/csgura/di"
	"github.com/go-akka/configuration"
	"github.com/nersion86/nlib/loader"
)

/* This Document is how to using nlib sample page */

//ConfigTest is HOCON TEST.
func ConfigTest() {

	cfg := loader.NewConfigFromFile("test.conf")

	strArr := cfg.GetStringList("modules")
	for _, v := range strArr {
		fmt.Println(v)
	}

	intV := cfg.GetInt32("my.config.hoy-int", 0)
	fmt.Println(intV)

	strV := cfg.GetString("my.config.hoy-string", "")
	fmt.Println(strV)

	//done
}

// This area is DI using loader module. loader is not safety. but esay binding di

//TestProvider is empty struct. this is basic. next step. not using empty struct.
type TestProvider struct {
}

//HelloWorld is sample provider. (member is all public. it just sample... this is not good)
type HelloWorld struct {
	Hello string
	World int
	Conf  *configuration.Config //this module is dependency injection !!
}

//NewHelloWorld is Create HelloWorld. HelloWorld is Need cfg module. cfg is already inject default!!
func NewHelloWorld(cfg *configuration.Config) *HelloWorld {
	//HOCON return int32 or int64. not int. golang is int != int32.
	//cast int32 to int.
	world := int(cfg.GetInt32("my.config.hoy-int", 0)) //type cast int32 to int.
	hello := cfg.GetString("my.config.hoy-string", "not_load")

	newHello := &HelloWorld{
		Hello: hello,
		World: world,
		Conf:  cfg,
	}

	return newHello
}

//Configure is DI abstract module method. !!
func (t *TestProvider) Configure(binder *di.Binder) {
	provider := func(injector di.Injector) interface{} {

		//not need Type Assertion. but default is using type assertion. (abstract -> class)
		cfg := injector.GetInstance((*configuration.Config)(nil)).(*configuration.Config)

		res := NewHelloWorld(cfg)
		return res
	}

	binder.BindProvider((*HelloWorld)(nil), provider)
	//or MyStyle... using this.
	//binder.Bind((*HelloWorld)(nil)).ToProvider(provider) //<- more flexilbe style.
}

//ImplimentsMake is sample of how to do configure impls (Read github.com/csgura/di)
func ImplimentsMake() *di.Implements {

	impl := di.NewImplements()

	//test.conf : testModule matched (modules array in this name)
	impl.AddImplement("testModule", &TestProvider{})
	return impl
}

//End of Configure DI sample.

//LoderTest is DI helper module test.
func LoaderTest() {

	ld := loader.LoadMoudleFromFileConfig("test.conf", ImplimentsMake())

	//hello module get.
	hello := ld.GetInstance((*HelloWorld)(nil)).(*HelloWorld)

	fmt.Println(hello.Hello)
	fmt.Println(hello.World)

	oString := hello.Conf.GetStringList("other.o-config")

	for _, v := range oString {
		fmt.Println(v)
	}

}

func main() {

	fmt.Println("NLIB Sample Test .... not using go test version")
	ConfigTest()

	fmt.Println("Loader Test ... ")
	LoaderTest()

}
