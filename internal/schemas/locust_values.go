package schemas

// Values for Create & Apply Locust Cluster
type LocustValues struct {
	// Namespace of locust cluster
	Namespace  string `yaml:"namespace"`

	// Name of locust cluster
	LocustName string `yaml:"locustName"`

	// Loucst Master Node
	Master     struct {
		// Resource Requests of Master Node
		Requests struct {
			// Request of Master Node CPU
			CPU    string `yaml:"cpu"`
			// Request of Master Node Memory
			Memory string `yaml:"memory"`
		}

		// Resource Limit of Master Node
		Limits struct {
			// Limit of Master Node CPU
			CPU    string `yaml:"cpu"`
			// Limit of Master Node Memory
			Memory string `yaml:"memory"`
		}

		// NodeSelector of Master Node
		NodeSelector map[string]string   `yaml:"nodeSelector"`

		// Tolerations of Master Node
		Tolerations  []map[string]string `yaml:"tolerations"`

		// Affinity of Master Node
		Affinity     map[string]string   `yaml:"affinity"`

		// Lables of Master Node
		Labels       map[string]string   `yaml:"labels"`

		// Annotations of Master Node
		Annotations  map[string]string   `yaml:"annotations"`
	} `yaml:"master"`

	// Locust Worker Node
	Worker struct {
		// Count of Worker Node
		Count    int `yaml:"count"`

		// Resource Requests of Worker Node
		Requests struct {
			// Request of Worker Node CPU
			CPU    string `yaml:"cpu"`
			// Request of Master Node Memory
			Memory string `yaml:"memory"`
		}
		// Resource Limits of Worker Node
		Limits struct {
			// Request of Worker Node CPU
			CPU    string `yaml:"cpu"`
			// Request of Worker Node Memory
			Memory string `yaml:"memory"`
		}

		// NodeSelector of Worker Node
		NodeSelector map[string]string   `yaml:"nodeSelector"`

		// Tolerations of Worker Node
		Tolerations  []map[string]string `yaml:"tolerations"`

		// Affinty of Worker Node
		Affinity     map[string]string   `yaml:"affinity"`

		// Labels of Worker Node
		Labels       map[string]string   `yaml:"labels"`

		// Annotations of Worker Node
		Annotations  map[string]string   `yaml:"annotations"`
	} `yaml:"worker"`

	// Locust Master Service
	Service struct {
		// Port of Master Service
		Port          int               `yaml:"port"`

		// Lables of Master Service
		Labels        map[string]string `yaml:"labels"`

		// Annotations of Master Service
		Annotations   map[string]string `yaml:"annotations"`

		// True if using EKS on Fargate
		EnableFargate bool              `yaml:"enableFargate"`
	} `yaml:"service"`

	// Locust Master Ingress
	Ingress struct {

		// Class of Master Ingress (alb/nginx)
		Class string `yaml:"class"`

		// host of Master Ingress (ex) locust-hello.example.com)
		Host  string `yaml:"host"`

		// ALB Ingress controller (if Class = 'alb')
		ALB   struct {
			// Scheme  of ALB
			Scheme         string `yaml:"scheme"`
			// SecurityGroups of ALB (separate with ,)
			SecurityGroups string `yaml:"securityGroups"`
			// CetificateARN of ALB (if ALB with SSL)
			CertificateARN string `yaml:"cetificateARN"`
		} `yaml:"alb"`

		// Labels of Master Ingress
		Labels      map[string]string `yaml:"labels"`

		// Annotations of Master Ingress
		Annotations map[string]string `yaml:"annotations"`
	} `yaml:"ingress"`
}
