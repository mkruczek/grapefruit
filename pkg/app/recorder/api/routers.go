package api

func (hc *HttpController) SetRouters() {

	apiV1Group := hc.server.Group("/api/v1")
	apiV1Group.GET("/healthz", healthz())

}
