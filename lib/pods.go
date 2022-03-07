package lib

import (
	"fmt"

	"k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func AdmitPods(ar v1.AdmissionReview) *v1.AdmissionResponse {
	podResource := metav1.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}

	// 判断资源类型
	if ar.Request.Resource != podResource {
		err := fmt.Errorf("expect resource to be %s", podResource)
		klog.Error(err)
		return ToV1AdmissionResponse(err)
	}

	// 反序列化数据到Pod
	raw := ar.Request.Object.Raw
	pod := corev1.Pod{}
	deserializer := Codecs.UniversalDeserializer()
	if _, _, err := deserializer.Decode(raw, nil, &pod); err != nil {
		klog.Error(err)
		return ToV1AdmissionResponse(err)
	}

	// 业务逻辑
	reviewResponse := v1.AdmissionResponse{}
	if pod.Name == "xiaolatiao" {
		reviewResponse.Allowed = false
		reviewResponse.Result = &metav1.Status{Code: 503, Message: "pod name cannot be xiaolatiao"}
	} else {
		reviewResponse.Allowed = true
	}

	return &reviewResponse
}
