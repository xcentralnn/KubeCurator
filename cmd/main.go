package main

import (
	"os"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	curatorv1 "github.com/xcentralnn/curator/api/v1alpha1"
	"github.com/xcentralnn/curator/internal/controller"
)

func main() {
	scheme := runtime.NewScheme()
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(curatorv1.AddToScheme(scheme))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		Metrics: server.Options{
			BindAddress: ":8080",
		},
	})
	if err != nil {
		os.Exit(1)
	}

	if err = (&controller.SmartScalerReconciler{
		Client: mgr.GetClient(),
	}).SetupWithManager(mgr); err != nil {
		os.Exit(1)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		os.Exit(1)
	}
}
