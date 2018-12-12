package main

import (
	"github.com/openshift/source-to-image/pkg/api"
	"github.com/openshift/source-to-image/pkg/build/strategies"
	"github.com/openshift/source-to-image/pkg/docker"
	s2ierr "github.com/openshift/source-to-image/pkg/errors"
	"github.com/openshift/source-to-image/pkg/scm/git"
	utilglog "github.com/openshift/source-to-image/pkg/util/glog"
	"log"
	"os"
	"fmt"
)

func getApiConfig(builderImage string, displayName string, stringURL string, dockerConfig *api.DockerConfig, tag string) api.Config {

	var config api.Config

	config.DisplayName = displayName
	config.BuilderImage = builderImage
	config.Tag = tag
	config.Source, _ = git.Parse(stringURL)
	config.BuilderPullPolicy = api.DefaultBuilderPullPolicy
	config.DockerConfig = dockerConfig

	return config
}

func getDockerConfig() api.DockerConfig {

	var dockerConfig api.DockerConfig
	dockerHostEndPoint := os.Getenv("DOCKER_HOST")
	dockerCertLocation := os.Getenv("DOCKER_CERT_PATH")

	dockerConfig.UseTLS = true
	dockerConfig.TLSVerify = true
	dockerConfig.KeyFile = dockerCertLocation + "/key.pem"
	dockerConfig.CertFile = dockerCertLocation + "/cert.pem"
	dockerConfig.CAFile = dockerCertLocation + "/ca.pem"
	dockerConfig.Endpoint = dockerHostEndPoint

	return dockerConfig
}

func main() {

	stringURL := "https://github.com/vorburger/s2i-java-example"
	//stringURL := "https://github.com/openshift/ruby-hello-world"
	//var builderImage = "centos/ruby-22-centos7"
	var builderImage = "fabric8/s2i-java"
	var tag = "vorburger:s2i-java-example"
	//var tag = "hello-world-app"
	var displayName = "latest"
	dockerConfig := getDockerConfig()
	client, err := docker.NewEngineAPIClient(&dockerConfig)
	if err != nil {
		log.Println(err)
	}
	apiConfig := getApiConfig(builderImage, displayName, stringURL, &dockerConfig, tag)
	builder, _, error := strategies.GetStrategy(client, &apiConfig)
	if error != nil {
		log.Println(error)
	}
	var glog = utilglog.StderrLog
	result, err := builder.Build(&apiConfig)
	if err != nil {
		glog.V(0).Infof("Build failed")
		s2ierr.CheckError(err)
	} else {
		if len(apiConfig.AsDockerfile) > 0 {
			glog.V(0).Infof("Application dockerfile generated in %s", apiConfig.AsDockerfile)
		} else {
			glog.V(0).Infof("Build completed successfully")

		}
	}
	for _, message := range result.Messages {
		fmt.Println(message)
		glog.V(1).Infof(message)
	}
}
