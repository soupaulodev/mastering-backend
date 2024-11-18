package t1fundamentos

func  TestIController() {
	var IController iController = &userController{}
	println(IController.Get())
	println(IController.Post())
	println(IController.Put())
	println(IController.Delete())
}

type iController interface {
	Get() (string, error)
	Post() (string, error)
	Put() (string, error)
	Delete() (string, error)
}

type userController struct {}

func (c *userController) Get() (string, error) { return "Método GET executado", nil }

func (c *userController) Post() (string, error) { return "Método POST executado", nil }

func (c *userController) Put() (string, error) { return "Método PUT executado", nil }

func (c *userController) Delete() (string, error) { return "Método DELETE executado", nil }
