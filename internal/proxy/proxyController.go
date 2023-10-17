package proxy

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/common"
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
}

func (c *ReverseProxyController) CreateReverseProxy(context *gin.Context) {

	serviceName := context.Param("name")

	serviceEntity, err := c.ServiceRepository.FindByName(serviceName)
	if err != nil {
		context.JSON(http.StatusNotFound, common.ErrorResponse(err.Error()))
	}

	// TODO load balancer
	hostname := serviceEntity.Hosts[0].Url
	port := serviceEntity.Hosts[0].Port

	urlString := hostname + ":" + strconv.Itoa(port)
	targetURL, _ := url.Parse(urlString)
	log.Printf("Proxy at : %s", targetURL.String())
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	log.Printf("Forwarding request to %s\n", context.Request.URL.String())
	// Let the reverse proxy do its job
	proxy.ServeHTTP(context.Writer, context.Request)

}
