// cmd/main.go

package main

import (
    "os"

    ctrl "sigs.k8s.io/controller-runtime"

    "github.com/xcentralnn/kubecurator/internal/controller"
)

func main() {
    mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
        MetricsBindAddress: ":8080",
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
