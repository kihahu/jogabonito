/*

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

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	fantasyv1 "k8s.io/api/v1"
)

// PlayerStatsReconciler reconciles a PlayerStats object
type PlayerStatsReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=fantasy.jogabonito.io,resources=playerstats,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=fantasy.jogabonito.io,resources=playerstats/status,verbs=get;update;patch

func (r *PlayerStatsReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("playerstats", req.NamespacedName)

	team := []fantasyv1.PlayerStatsSpec{}

	// your logic here
	fmt.Println("Inside the controller")

	// print players with > 200 stat
	player := &fantasyv1.PlayerStats{}
	err := r.Get(context.TODO(), req.NamespacedName, player)
	// fmt.Println(player.Status.IctIndex)

	if err != nil {
		fmt.Print(err)
	}

	if player.Spec.IctIndex == "200" {
		team = append(team, player.Spec)
	}

	// print team
	for _, t := range team {
		fmt.Println(t)
	}

	return ctrl.Result{}, nil
}

func (r *PlayerStatsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&fantasyv1.PlayerStats{}).
		Complete(r)
}
