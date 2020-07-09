package server

import (
	"github.com/labstack/echo/v4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ctx struct {
	echo.Context
	KubeClient *kubernetes.Clientset
}

func createContext(clientset *kubernetes.Clientset) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ctx{c, clientset}
			return next(cc)
		}
	}
}

func getNamespaces(c echo.Context) error {
	cc := c.(*ctx)

	namespacesClient := cc.KubeClient.CoreV1().Namespaces()
	namespaces, err := namespacesClient.List(metav1.ListOptions{})

	if err != nil {
		return echo.NewHTTPError(400, err)
	}

	// TODO(mark): collect .namespaces.items[*].name
	// if .namespaces.items[*].status.phase is "Active"

	return cc.JSON(200, namespaces)
}

func getNodes(c echo.Context) error {
	cc := c.(*ctx)

	// TODO(mark): query param labels to filter

	nodesClient := cc.KubeClient.CoreV1().Nodes()
	nodes, err := nodesClient.List(metav1.ListOptions{})

	if err != nil {
		return echo.NewHTTPError(400, err)
	}

	return cc.JSON(200, nodes)
}
