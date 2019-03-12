/*
Copyright 2018 Pressinfra SRL

This file is subject to the terms and conditions defined in file LICENSE,
which is part of this source code package.
*/

package organization

import (
	"fmt"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/testing_frameworks/integration/addr"

	logf "github.com/presslabs/controller-util/log"
	"github.com/presslabs/dashboard/pkg/apis"
	"github.com/presslabs/dashboard/pkg/apiserver"
	"github.com/presslabs/dashboard/pkg/apiserver/internal/metadata"
)

var cfg *rest.Config
var t *envtest.Environment

func TestAPIServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Organization API Server Suite", []Reporter{envtest.NewlineReporter{}})
}

var _ = BeforeSuite(func() {
	var err error

	logf.SetLogger(logf.ZapLoggerTo(GinkgoWriter, true))

	t = &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "..", "..", "config", "crds"),
			filepath.Join("..", "..", "..", "..", "vendor/github.com/coreos/prometheus-operator/example/prometheus-operator-crd"),
			filepath.Join("..", "..", "..", "..", "vendor/github.com/presslabs/mysql-operator/config/crds"),
			filepath.Join("..", "..", "..", "..", "vendor/github.com/presslabs/wordpress-operator/config/crds"),
		},
	}
	apis.AddToScheme(scheme.Scheme)

	cfg, err = t.Start()
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	t.Stop()
})

func SetupAPIServer(mgr manager.Manager) *apiserver.APIServer {
	grpcPort, _, err := addr.Suggest()
	Expect(err).To(Succeed())

	httpPort, _, err := addr.Suggest()
	Expect(err).To(Succeed())

	opts := &apiserver.APIServerOptions{
		Manager:  mgr,
		HTTPAddr: fmt.Sprintf(":%d", httpPort),
		GRPCAddr: fmt.Sprintf(":%d", grpcPort),
		AuthFunc: metadata.FakeAuth,
	}

	server, err := apiserver.NewAPIServer(opts)
	Expect(err).To(Succeed())

	mgr.Add(server)

	return server
}

func StartTestManager(mgr manager.Manager) chan struct{} {
	stop := make(chan struct{})
	go func() {
		defer GinkgoRecover()
		Expect(mgr.Start(stop)).To(Succeed())
	}()
	return stop
}
