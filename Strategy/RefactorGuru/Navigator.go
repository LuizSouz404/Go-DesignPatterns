package main

type Navigator struct {
	BuildRouteBehavior
}

func (n *Navigator) SetRouteType(buildRoute BuildRouteBehavior) {
	n.BuildRouteBehavior = buildRoute
}
