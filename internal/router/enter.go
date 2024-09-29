package router

type GroupRouter struct {
	Products ProductRouter
}

var AppRouter = new(GroupRouter)
