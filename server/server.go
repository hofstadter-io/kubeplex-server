package server

import (
	"github.com/labstack/echo/v4"
	"k8s.io/client-go/kubernetes"
)

type RestServer struct {
	KubeClient *kubernetes.Clientset
}

func (r *RestServer) Start() {
	e := echo.New()

	e.Use(createContext(r.KubeClient))

	e.GET("/namespaces", getNamespaces)
	e.GET("/nodes", getNodes)

	// TODO(mark) additional endpoints
	// - GET /namespace/:name label, resources

	e.Start(":1323")
}
