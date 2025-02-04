package controllers_test

import (
	"testing"

	korifiv1alpha1 "code.cloudfoundry.org/korifi/controllers/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAppWorkloadsController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

func createAppWorkload(namespace, name string) *korifiv1alpha1.AppWorkload {
	return &korifiv1alpha1.AppWorkload{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: korifiv1alpha1.AppWorkloadSpec{
			AppGUID:          "premium_app_guid_1234",
			GUID:             "guid_1234",
			Version:          "version_1234",
			Image:            "gcr.io/foo/bar",
			ImagePullSecrets: []corev1.LocalObjectReference{{Name: "some-secret-name"}},
			Command: []string{
				"/bin/sh",
				"-c",
				"while true; do echo hello; sleep 10;done",
			},
			ProcessType: "worker",
			Env:         []corev1.EnvVar{},
			Health: korifiv1alpha1.Healthcheck{
				Type:      "http",
				Port:      int32(8080),
				Endpoint:  "/healthz",
				TimeoutMs: uint(60000),
			},
			Ports:         []int32{8888, 9999},
			Instances:     1,
			MemoryMiB:     1024,
			DiskMiB:       2048,
			CPUMillicores: 5,
		},
	}
}
