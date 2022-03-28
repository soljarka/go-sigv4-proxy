package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/kelseyhightower/envconfig"
	"github.com/soljarka/go-sigv4-proxy/config"
)

var cfg config.Config

// Serve a reverse proxy for a given url
func serveReverseProxy(res http.ResponseWriter, req *http.Request) {

	fmt.Println("Request: ", req.URL.String())
	// parse the url
	url, _ := url.Parse(cfg.Endpoint)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	creds := credentials.NewCredentials(
		&credentials.StaticProvider{
			Value: credentials.Value{
				AccessKeyID:     cfg.AwsAccessKeyID,
				SecretAccessKey: cfg.AwsSecretAccessKey,
				SessionToken:    cfg.AwsSessionToken,
			},
		})
	// Sign the request
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(res, "Error reading request body: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	signer := v4.NewSigner(creds)
	signer.Sign(req, bytes.NewReader(body), cfg.Service, cfg.Region, time.Now())

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

func main() {
	err := envconfig.Process("gosigv4proxy", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Handle("/", http.HandlerFunc(serveReverseProxy))
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil)
}
