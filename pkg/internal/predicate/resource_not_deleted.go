/*
Copyright 2019 Pressinfra SRL.

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

package predicate

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// ResourceNotDeleted filters out resources with non-zero deletion timestamp and
// deletion events
var ResourceNotDeleted predicate.Predicate = predicate.Funcs{
	// We need to check the CreateEvent since this might refer to items in cache,
	// not to items in store, and for an existing object we'll get CreateEvent
	// not UpdateEvent
	CreateFunc: func(e event.CreateEvent) bool {
		return e.Meta.GetDeletionTimestamp().IsZero()
	},
	UpdateFunc: func(e event.UpdateEvent) bool {
		return e.MetaNew.GetDeletionTimestamp().IsZero()
	},
	DeleteFunc: func(e event.DeleteEvent) bool {
		return false
	},
	GenericFunc: func(e event.GenericEvent) bool {
		return e.Meta.GetDeletionTimestamp().IsZero()
	},
}