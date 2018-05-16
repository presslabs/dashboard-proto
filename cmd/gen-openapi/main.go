/*
Copyright 2018 Pressinfra SRL

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
	"fmt"
	"os"

	kutilv1 "github.com/appscode/kutil/apiextensions/v1beta1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/presslabs/dashboard/pkg/apis/projects"
	projectsv1 "github.com/presslabs/dashboard/pkg/apis/projects/v1alpha1"
	dashboardversion "github.com/presslabs/dashboard/pkg/version"
)

var (
	group       = projects.GroupName
	version     = projectsv1.SchemeGroupVersion.Version
	resourceCRD = ""
)

func initFlags(resourceCRD *string, output *string, cfg *kutilv1.Config, fs *pflag.FlagSet) *pflag.FlagSet {
	fs.Var(&cfg.Labels, "labels", "Lables")
	fs.Var(&cfg.Annotations, "annotations", "Annotations")
	fs.StringVar(resourceCRD, "crd", "plural.example.com", "Custom resource definition")
	fs.StringVar(output, "output", "yaml", "Output format")
	return fs
}

func main() {
	cfg := &kutilv1.Config{}
	resourceCRD := ""
	output := ""

	cmd := &cobra.Command{
		Use:   "dashboard-gen-crd",
		Short: fmt.Sprintf("Generate Dashboard CRDs (%s)", dashboardversion.Get()),
		Run: func(cmd *cobra.Command, args []string) {
			if CRDcfg, exists := projectsv1.CRDs[resourceCRD]; !exists {
				fmt.Fprintf(os.Stderr, "%s CRD does not exist\n", resourceCRD)
				os.Exit(1)
			} else {
				CRDcfg.Labels = cfg.Labels
				CRDcfg.Annotations = cfg.Annotations
				crd := kutilv1.NewCustomResourceDefinition(CRDcfg)
				kutilv1.MarshallCrd(os.Stdout, crd, output)
			}
		},
	}

	flags := cmd.Flags()
	initFlags(&resourceCRD, &output, cfg, flags)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
