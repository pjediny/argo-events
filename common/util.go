/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"fmt"
	"github.com/rs/zerolog"
	"k8s.io/client-go/kubernetes"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
)

// DefaultConfigMapName returns a formulated name for a configmap name based on the sensor-controller deployment name
func DefaultConfigMapName(controllerName string) string {
	return fmt.Sprintf("%s-configmap", controllerName)
}

// DefaultSensorDeploymentName returns a formulated name for sensor deployment
func DefaultSensorDeploymentName(deploymentName string) string {
	return fmt.Sprintf("%s-sensor-deployement", deploymentName)
}

// DefaultSensorJobName returns a formulated name for a sensor job
func DefaultSensorJobName(jobName string) string {
	return fmt.Sprintf("%s-sensor-job", jobName)
}

// DefaultGatewayPodName returns a formulated name for a gateway deployment
func DefaultGatewayPodName(deploymentName string) string {
	return fmt.Sprintf("%s-gateway-deployment", deploymentName)
}

// DefaultGatewayServiceName returns a formulated name for a gateway service
func DefaultGatewayServiceName(serviceName string) string {
	return fmt.Sprintf("%s-gateway-svc", serviceName)
}

// DefaultSensorServiceName returns a formulated name for a sensor service
func DefaultSensorServiceName(serviceName string) string {
	return fmt.Sprintf("%s-sensor-svc", serviceName)
}

// DefaultGatewayConfigurationName returns a formulated name for a gateway configuration
func DefaultGatewayConfigurationName(gatewayName string, configurationName string) string {
	return fmt.Sprintf("%s/%s", gatewayName, configurationName)
}

// GetClientConfig return rest config, if path not specified, assume in cluster config
func GetClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

// ServerResourceForGroupVersionKind finds the API resources that fit the GroupVersionKind schema
func ServerResourceForGroupVersionKind(disco discovery.DiscoveryInterface, gvk schema.GroupVersionKind) (*metav1.APIResource, error) {
	resources, err := disco.ServerResourcesForGroupVersion(gvk.GroupVersion().String())
	if err != nil {
		return nil, err
	}
	for _, r := range resources.APIResources {
		if r.Kind == gvk.Kind {
			return &r, nil
		}
	}
	return nil, fmt.Errorf("server is unable to handle %s", gvk)
}

// SendSuccessResponse sends http success response
func SendSuccessResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(SuccessResponse))
}

// SendErrorResponse sends http error response
func SendErrorResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write([]byte(ErrorResponse))
}

// LoggerConf returns standard logging configuration
func LoggerConf() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return fmt.Sprintf("| %-6s|", i)
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("msg: %s | ", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	return  output
}

// GetLoggerContext returns a logger with input options
func GetLoggerContext(opt zerolog.ConsoleWriter) zerolog.Context {
	return zerolog.New(opt).With().Timestamp()
}

// GetSecret retrieves the secret value from the secret in namespace with name and key
func GetSecret(client kubernetes.Interface, namespace string, name, key string) (string, error) {
	secret, err := client.CoreV1().Secrets(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	val, ok := secret.Data[key]
	if !ok {
		return "", fmt.Errorf("secret '%s' does not have the key '%s'", name, key)
	}
	return string(val), nil
}
