/*
Copyright 2020 The klocust Authors

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

package schemas

// Values for Create & Apply Locust Cluster
type LocustValues struct {
	// Namespace of locust cluster
	Namespace string `yaml:"namespace"`

	// Name of locust cluster
	LocustName string `yaml:"locustName"`

	// Locust ConfigMap
	ConfigMap struct {

		// filename of locustfile
		LocustFilename     string `yaml:"locustFilename"`
		LocustFileContents string `yaml:"locustFileContents"`
	}

	// Locust Main Node
	Main struct {
		// Docker image of Main Node
		Image string `yaml:"image"`

		// Resource Requests of Main Node
		Requests struct {
			// Request of Main Node CPU
			CPU string `yaml:"cpu"`
			// Request of Main Node Memory
			Memory string `yaml:"memory"`
		}

		// Resource Limit of Main Node
		Limits struct {
			// Limit of Main Node CPU
			CPU string `yaml:"cpu"`
			// Limit of Main Node Memory
			Memory string `yaml:"memory"`
		}

		// NodeSelector of Main Node
		NodeSelector map[string]string `yaml:"nodeSelector"`

		// Tolerations of Main Node
		Tolerations []map[string]string `yaml:"tolerations"`

		// Affinity of Main Node
		Affinity map[string]string `yaml:"affinity"`

		// Lables of Main Node
		Labels map[string]string `yaml:"labels"`

		// Annotations of Main Node
		Annotations map[string]string `yaml:"annotations"`
	} `yaml:"main"`

	// Locust Worker Node
	Worker struct {
		// Docker image of worker node
		Image string `yaml:"image"`

		// Count of Worker Node
		Count int `yaml:"count"`

		// Resource Requests of Worker Node
		Requests struct {
			// Request of Worker Node CPU
			CPU string `yaml:"cpu"`
			// Request of Main Node Memory
			Memory string `yaml:"memory"`
		}
		// Resource Limits of Worker Node
		Limits struct {
			// Request of Worker Node CPU
			CPU string `yaml:"cpu"`
			// Request of Worker Node Memory
			Memory string `yaml:"memory"`
		}

		// NodeSelector of Worker Node
		NodeSelector map[string]string `yaml:"nodeSelector"`

		// Tolerations of Worker Node
		Tolerations []map[string]string `yaml:"tolerations"`

		// Affinty of Worker Node
		Affinity map[string]string `yaml:"affinity"`

		// Labels of Worker Node
		Labels map[string]string `yaml:"labels"`

		// Annotations of Worker Node
		Annotations map[string]string `yaml:"annotations"`
	} `yaml:"worker"`

	// Locust Main Service
	Service struct {
		// Port of Main Service
		Port int `yaml:"port"`

		// Lables of Main Service
		Labels map[string]string `yaml:"labels"`

		// Annotations of Main Service
		Annotations map[string]string `yaml:"annotations"`

		// True if using EKS on Fargate
		EnableFargate bool `yaml:"enableFargate"`
	} `yaml:"service"`

	// Locust Main Ingress
	Ingress struct {

		// Class of Main Ingress (alb/nginx)
		Class string `yaml:"class"`

		// host of Main Ingress (ex) locust-hello.example.com)
		Host string `yaml:"host"`

		// ALB Ingress controller (if Class = 'alb')
		ALB struct {
			// Scheme  of ALB
			Scheme string `yaml:"scheme"`
			// SecurityGroups of ALB (separate with ,)
			SecurityGroups string `yaml:"securityGroups"`
			// CetificateARN of ALB (if ALB with SSL)
			CertificateARN string `yaml:"certificateARN"`
		} `yaml:"alb"`

		// Labels of Main Ingress
		Labels map[string]string `yaml:"labels"`

		// Annotations of Main Ingress
		Annotations map[string]string `yaml:"annotations"`
	} `yaml:"ingress"`
}
