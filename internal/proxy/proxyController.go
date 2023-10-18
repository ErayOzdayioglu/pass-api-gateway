package proxy

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/common"
	"github.com/ErayOzdayioglu/api-gateway/internal/loadbalancer"
	"github.com/ErayOzdayioglu/api-gateway/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

type ReverseProxyController struct {
	ServiceRepository service.ServiceRepository
	LoadBalancer      loadbalancer.LoadBalancer
}

func (c *ReverseProxyController) CreateReverseProxy(context *gin.Context) {

	path := context.Param("name")

	serviceEntity, err := c.ServiceRepository.FindByPath("/" + path)
	if err != nil {
		context.JSON(http.StatusNotFound, common.ErrorResponse(err.Error()))
	}

	host, err := c.LoadBalancer.SelectTheRouteWithRoundRobin(serviceEntity)

	hostname := host.Url
	port := host.Port

	urlString := hostname + ":" + strconv.Itoa(port)
	targetURL, _ := url.Parse(urlString)
	log.Printf("Proxy at : %s", targetURL.String())
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	log.Printf("Forwarding request to %s\n", context.Request.URL.String())
	// Let the reverse proxy do its job
	proxy.ServeHTTP(context.Writer, context.Request)

}
