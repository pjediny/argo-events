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

package main

import (
	"context"
	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/controllers/sensor"
	"os"
)

func main() {
	// kubernetes configuration
	kubeConfig, _ := os.LookupEnv(common.EnvVarKubeConfig)
	restConfig, err := common.GetClientConfig(kubeConfig)
	if err != nil {
		panic(err)
	}

	// sensor-controller configuration
	configMap, ok := os.LookupEnv(common.EnvVarConfigMap)
	if !ok {
		configMap = common.DefaultConfigMapName(common.DefaultSensorControllerDeploymentName)
	}

	namespace, ok := os.LookupEnv(common.SensorNamespace)
	if !ok {
		namespace = common.DefaultControllerNamespace
	}

	// create a new sensor controller
	controller := sensor.NewSensorController(restConfig, configMap, namespace)
	// watch updates to sensor controller configuration
	err = controller.ResyncConfig(namespace)
	if err != nil {
		panic(err)
	}

	go controller.Run(context.Background(), 1, 1)

	// Wait forever
	select {}
}
