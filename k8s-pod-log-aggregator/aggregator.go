package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/sync/errgroup"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func int64Ptr(i int64) *int64 {
	return &i
}

func getClient() *kubernetes.Clientset {
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

	return clientset

}

func ListPods(namespace string) {

	clientset := getClient()
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}

func FetchLogsConcurrently(namespace string, maxWorkers int) (map[string]string, []string, error) {
	clientset := getClient()

	g, ctx := errgroup.WithContext(context.Background())

	sem := make(chan struct{}, maxWorkers)

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	logsMap := make(map[string]string)
	var errors []string
	var mu sync.Mutex

	for _, pod := range pods.Items {

		pod := pod

		sem <- struct{}{}

		g.Go(func() error {
			defer func() { <-sem }()

			req := clientset.CoreV1().Pods(namespace).GetLogs(pod.Name, &v1.PodLogOptions{
				TailLines: int64Ptr(100),
			})

			stream, err := req.Stream(ctx)
			if err != nil {
				mu.Lock()
				errors = append(errors, fmt.Sprintf("%s: %v", pod.Name, err))
				mu.Unlock()
				return err
			}
			defer stream.Close()

			data, err := io.ReadAll(stream)
			if err != nil {
				mu.Lock()
				errors = append(errors, fmt.Sprintf("%s: read error %v", pod.Name, err))
				mu.Unlock()
				return err
			}

			mu.Lock()
			logsMap[pod.Name] = string(data)
			mu.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Encountered Error Exiting %v ", err)
	}
	return logsMap, errors, nil
}
