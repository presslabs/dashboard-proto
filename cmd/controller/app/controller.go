/*
Copyright 2014 The Kubernetes Authors.

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

// Package app implements a controller manager that runs a set of active
// controllers, like projects controller, site controller or edge controller
package app

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/appscode/kutil/tools/clientcmd"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	apiextensions_clientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"

	dashclientset "github.com/presslabs/dashboard/pkg/client/clientset/versioned"
	dashintscheme "github.com/presslabs/dashboard/pkg/client/clientset/versioned/scheme"
	dashinformers "github.com/presslabs/dashboard/pkg/client/informers/externalversions"
	projectscontroller "github.com/presslabs/dashboard/pkg/controller/projects"

	"github.com/presslabs/dashboard/cmd/controller/app/options"
	"github.com/presslabs/dashboard/pkg/controller"
	"github.com/presslabs/dashboard/pkg/version"
)

const (
	controllerAgentName = "presslabs-dashboard-controller"
)

// NewControllerManagerCommand creates a *cobra.Command object with default parameters
func NewControllerManagerCommand(stopCh <-chan struct{}) *cobra.Command {
	o := options.NewControllerManagerOptions()
	cmd := &cobra.Command{
		Use:   "presslabs-controller",
		Short: fmt.Sprintf("Presslabs Dashboard Controller (%s)", version.Get()),
		Run: func(cmd *cobra.Command, args []string) {

			if err := o.Validate(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}

			if err := Run(o, stopCh); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}

	flags := cmd.Flags()
	o.AddFlags(flags)

	return cmd
}

// Run Presslabs Controller Manager.  This should never exit.
func Run(c *options.ControllerManagerOptions, stopCh <-chan struct{}) error {
	glog.Infof("Starting Dashboard Operator Controller (%s)...", version.Get())

	run := func(_ <-chan struct{}) {
		var wg sync.WaitGroup
		// Start the Project Controller
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx, err := buildControllerContext(c)

			if err != nil {
				glog.Fatalf(err.Error())
				return
			}

			ctrl, err := projectscontroller.NewController(ctx)
			if err != nil {
				glog.Fatalf(err.Error())
				return
			}
			ctrl.Run(stopCh)
			<-stopCh
		}()

		wg.Wait()
		glog.Fatalf("Control loops exited")
	}

	if !c.LeaderElect {
		run(stopCh)
		return nil
	}

	ctx, err := buildControllerContext(c)

	if err != nil {
		return err
	}

	leaderElectionClient, err := kubernetes.NewForConfig(rest.AddUserAgent(ctx.RESTConfig, "leader-election"))

	if err != nil {
		glog.Fatalf("error creating leader election client: %s", err.Error())
	}

	startLeaderElection(c, leaderElectionClient, ctx.Recorder, run)
	panic("unreachable")
}

func buildControllerContext(c *options.ControllerManagerOptions) (*controller.Context, error) {
	// Create a Kubernetes api client
	kubeCfg, err := clientcmd.BuildConfigFromContext(c.Kubeconfig, c.KubeconfigContext)
	if err != nil {
		return nil, fmt.Errorf("error creating kubernetes rest api client: %s", err.Error())
	}
	glog.Infof("Kubernetes API server: %s", kubeCfg.Host)

	// Create a Kubernetes api client
	cl, err := kubernetes.NewForConfig(kubeCfg)

	if err != nil {
		return nil, fmt.Errorf("error creating kubernetes client: %s", err.Error())
	}

	// Create the CRD client
	crdcl, err := apiextensions_clientset.NewForConfig(kubeCfg)
	if err != nil {
		return nil, fmt.Errorf("error creating kubernetes CRD client: %s", err.Error())
	}

	// Create a Navigator api client
	intcl, err := dashclientset.NewForConfig(kubeCfg)

	if err != nil {
		return nil, fmt.Errorf("error creating internal group client: %s", err.Error())
	}

	// Create event broadcaster
	// Add oxygen types to the default Kubernetes Scheme so Events can be logged properly
	dashintscheme.AddToScheme(scheme.Scheme)
	glog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.V(4).Infof)
	eventBroadcaster.StartRecordingToSink(&corev1client.EventSinkImpl{Interface: cl.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	kubeSharedInformerFactory := informers.NewFilteredSharedInformerFactory(cl, time.Second*30, "", nil)
	dashboardInformerFactory := dashinformers.NewFilteredSharedInformerFactory(intcl, time.Second*30, "", nil)
	return &controller.Context{
		RESTConfig:                     kubeCfg,
		KubeClient:                     cl,
		KubeSharedInformerFactory:      kubeSharedInformerFactory,
		Recorder:                       recorder,
		DashboardClient:                intcl,
		DashboardSharedInformerFactory: dashboardInformerFactory,
		CRDClient:                      crdcl,
		InstallCRDs:                    c.InstallCRDs,
	}, nil
}

func startLeaderElection(c *options.ControllerManagerOptions, leaderElectionClient kubernetes.Interface, recorder record.EventRecorder, run func(<-chan struct{})) {
	// Identity used to distinguish between multiple controller manager instances
	id, err := os.Hostname()
	if err != nil {
		glog.Fatalf("error getting hostname: %s", err.Error())
	}

	// Lock required for leader election
	rl := resourcelock.EndpointsLock{
		EndpointsMeta: metav1.ObjectMeta{
			Namespace: c.LeaderElectionNamespace,
			Name:      controllerAgentName,
		},
		Client: leaderElectionClient.CoreV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity:      id + "-external-" + controllerAgentName,
			EventRecorder: recorder,
		},
	}

	// Try and become the leader and start controller manager loops
	leaderelection.RunOrDie(leaderelection.LeaderElectionConfig{
		Lock:          &rl,
		LeaseDuration: c.LeaderElectionLeaseDuration,
		RenewDeadline: c.LeaderElectionRenewDeadline,
		RetryPeriod:   c.LeaderElectionRetryPeriod,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: run,
			OnStoppedLeading: func() {
				glog.Fatalf("leaderelection lost")
			},
		},
	})
}
