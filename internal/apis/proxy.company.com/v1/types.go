package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type SidecarSpec struct {
	AppName string `json:"appName"`
	Cb      struct {
		Timeout int `json:"timeout"`
		Tries   int `json:"tries"`
	} `json:"cb"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Sidecar struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SidecarSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SidecarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Sidecar `json:"items"`
}
