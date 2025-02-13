package apiv1

// GetWorkerJoinInfoRPC is the path for the GetWorkerJoinInfo RPC.
const GetWorkerJoinInfoRPC = "k8sd/worker/info"

// GetWorkerJoinInfoRequest is the request message for the GetWorkerJoinInfo RPC.
type GetWorkerJoinInfoRequest struct {
	// Address is the address of the worker node.
	Address string `json:"address"`
}

// GetWorkerJoinInfoResponse is the response message for the GetWorkerJoinInfo RPC.
type GetWorkerJoinInfoResponse struct {
	// CACert is the PEM encoded certificate authority of the cluster.
	CACert string `json:"ca,omitempty"`
	// ClientCACert is the PEM encoded certificate authority of the cluster clients.
	ClientCACert string `json:"client-ca,omitempty"`
	// APIServers is a list of kube-apiserver endpoints of the cluster.
	APIServers []string `json:"apiServers"`
	// KubeletClientCert is the certificate to use in kubelet to authenticate with kube-apiserver.
	KubeletClientCert string `json:"kubeletClientCert"`
	// KubeletClientKey is the private key to use in kubelet to authenticate with kube-apiserver.
	KubeletClientKey string `json:"kubeletClientKey"`
	// KubeProxyClientCert is the certificate to use in kube-proxy to authenticate with kube-apiserver.
	KubeProxyClientCert string `json:"kubeProxyClientCert"`
	// KubeProxyClientKey is the private key to use in kube-proxy to authenticate with kube-apiserver.
	KubeProxyClientKey string `json:"kubeProxyClientKey"`
	// PodCIDR is the configured CIDR for pods in the cluster.
	PodCIDR string `json:"podCIDR"`
	// ServiceCIDR is the configured CIDR for services in the cluster.
	ServiceCIDR string `json:"serviceCIDR"`
	// ClusterDNS is the DNS server address of the cluster.
	ClusterDNS string `json:"clusterDNS,omitempty"`
	// ClusterDomain is the DNS domain of the cluster.
	ClusterDomain string `json:"clusterDomain,omitempty"`
	// CloudProvider is the cloud provider used in the cluster.
	CloudProvider string `json:"cloudProvider,omitempty"`
	// KubeletCert is the certificate to use for kubelet TLS. It will be empty if the cluster is not using self-signed certificates.
	KubeletCert string `json:"kubeletCrt,omitempty"`
	// KubeletKey is the private key to use for kubelet TLS. It will be empty if the cluster is not using self-signed certificates.
	KubeletKey string `json:"kubeletKey,omitempty"`
	// K8sdPublicKey is the public key that can be used to validate authenticity of cluster messages.
	K8sdPublicKey string `json:"k8sdPublicKey,omitempty"`
	// Annotations is a map of strings that can be used to store arbitrary metadata configuration.
	Annotations map[string]string `json:"annotations,omitempty"`
}
