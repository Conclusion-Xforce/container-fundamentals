# Everything is YAML

The goal of this exercise is to deploy the application that we built in [GO GO GO](../01-container/README.md) to a Kubernetes cluster. There are multiple ways to approach this: write a YAML manifest from scratch or use `kubectl` with the  `--dry-run=client` option to generate a manifest for you. Either way, you can use your own image or one of the images from our Quay repository on [quay.io/jitseklomp_xforce/golang-fib]. 

## Service
The next step is to leverage the built-in load balancing capabilities. Make sure to run multiple instances of the container and use a Kubernetes `Service` to direct traffic to all replicas of your app. Test your work with `kubectl port-forward`.
