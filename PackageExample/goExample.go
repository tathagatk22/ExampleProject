package main

import (
	"github.com/openshift/source-to-image/pkg/api"
	"github.com/openshift/source-to-image/pkg/build/strategies/dockerfile"
	"log"
	"net/url"
	"fmt"
)

func setConfig(config *api.Config, builderImage string, displayName string, tag string, stringURL string) {
	//fmt.Printf("%t" ,config.DisplayName )
	config.DisplayName = displayName
	config.BuilderImage = builderImage
	config.Tag = tag
	parseURL, err := url.Parse(stringURL)
	if err != nil{
		log.Fatal(err)
	}
	config.Source.URL = struct {
		Scheme     string
		Opaque     string
		User       *url.Userinfo
		Host       string
		Path       string
		RawPath    string
		ForceQuery bool
		RawQuery   string
		Fragment   string
	}{Scheme: parseURL.Scheme, Opaque: parseURL.Opaque, User: parseURL.User, Host: parseURL.Host, Path: parseURL.Path, RawPath: parseURL.RawPath, ForceQuery: parseURL.ForceQuery, RawQuery: parseURL.RawQuery, Fragment: parseURL.Fragment}
}

func main() {
	stringURL := "https://github.com/openshift/sti-python.git"
	var Result *api.Result
	var config api.Config
	var builderImage = "registry.access.redhat.com/rhscl/python-35-rhel7"
	var tag = "latest"
	var displayName = "python-35-rhel7-app"
	var builder dockerfile.Dockerfile
	setConfig(&config, builderImage, displayName, tag, stringURL)
	Result, error := builder.Build(&config)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("\n", Result, "\n", error)
}
