package t1fundamentos

func  TestUserController() {
	var userController IUserController = &UserController{}
	println(userController.Get())
	println(userController.Post())
	println(userController.Put())
	println(userController.Delete())
}

type IUserController interface {
	Get() (string, error)
	Post() (string, error)
	Put() (string, error)
	Delete() (string, error)
}

type UserController struct {}

func (c *UserController) Get() (string, error) { return "Método GET executado", nil }

func (c *UserController) Post() (string, error) { return "Método POST executado", nil }

func (c *UserController) Put() (string, error) { return "Método PUT executado", nil }

func (c *UserController) Delete() (string, error) { return "Método DELETE executado", nil }
