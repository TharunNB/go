package main

import (
	// "context"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func getReplicas(dep *appsv1.Deployment) int32 {
	if dep.Spec.Replicas == nil {
		return 1
	}
	return *dep.Spec.Replicas
}

func getClient() (*kubernetes.Clientset, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset, nil
}

func Patcher(dep *appsv1.Deployment, clientset *kubernetes.Clientset) {
	patch := map[string]interface{}{
		"spec": map[string]interface{}{
			"replicas": 3,
		},
	}

	patchBytes, _ := json.Marshal(patch)

	_, err := clientset.AppsV1().
		Deployments(dep.Namespace).
		Patch(
			context.TODO(),
			dep.Name,
			types.StrategicMergePatchType,
			patchBytes,
			metav1.PatchOptions{},
		)

	if err != nil {
		fmt.Println("Patch failed:", err)
		return
	}

	fmt.Printf("[SUCCESS] %s corrected to 3 replicas\n", dep.Name)
}

func ReplicaWatcher() {

	clientset, err := getClient()
	if err != nil {
		panic(err)
	}

	// Create Informer Factory
	informerFactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)

	// GEtting Informer for Pods
	podInformer := informerFactory.Core().V1().Pods()

	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*corev1.Pod)
			fmt.Printf("Pod Added: %s/%s\n", pod.Namespace, pod.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldPod := oldObj.(*corev1.Pod)
			newPod := newObj.(*corev1.Pod)
			fmt.Printf("Pod updated: %s/%s\n", newPod.Namespace, newPod.Name)
			fmt.Printf("  Old phase: %s, New phase: %s\n", oldPod.Status.Phase, newPod.Status.Phase)
		},
		DeleteFunc: func(obj interface{}) {
			var pod *corev1.Pod

			switch t := obj.(type) {
			case *corev1.Pod:
				pod = t
			case cache.DeletedFinalStateUnknown:
				pod = t.Obj.(*corev1.Pod)
			default:
				return
			}
			fmt.Printf("Pod DELETED: %s/%s\n", pod.Namespace, pod.Name)
		},
	})

	// Deployment Informer
	DeploymentInformer := informerFactory.Apps().V1().Deployments()

	DeploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			dep := obj.(*appsv1.Deployment)

			replicas := getReplicas(dep)

			if replicas > 3 {
				fmt.Printf("[WARN] %s has too many replicas: %d\n", dep.Name, replicas)
				Patcher(dep, clientset)
			} else {
				fmt.Printf("[OK] %s deployed with %d replicas\n", dep.Name, replicas)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldDep := oldObj.(*appsv1.Deployment)
			newDep := newObj.(*appsv1.Deployment)
			if oldDep.ResourceVersion == newDep.ResourceVersion {
				return
			}

			oldReplicas := getReplicas(oldDep)
			newReplicas := getReplicas(newDep)

			// Only react if replicas actually changed
			if oldReplicas != newReplicas {
				fmt.Printf(
					"Deployment %s scaled: %d → %d\n",
					newDep.Name,
					oldReplicas,
					newReplicas,
				)
			}

			// Enforce rule
			if newReplicas > 3 {
				fmt.Printf(
					"[WARN] %s exceeds replica limit: %d\n",
					newDep.Name,
					newReplicas,
				)
				Patcher(newDep, clientset)
			}
		},
	})

	stopCh := make(chan struct{})

	informerFactory.Start(stopCh)

	if !cache.WaitForCacheSync(stopCh, podInformer.Informer().HasSynced, DeploymentInformer.Informer().HasSynced) {
		panic("failed to sync caches")
	}

	fmt.Println("All Informers synced, Running...")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("Shutting down informer")
	close(stopCh)

}
