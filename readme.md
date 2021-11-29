# Description

This image could be deployed in Docker Swarm or Kubernetes.

- v0.1 and v0.2 comes with different background color.
- /ping to get the container / pod ID.

Example curl command run to showcase load balancing among the containers/pods

while true; do curl {url}/ping; printf "\n"; sleep 1; done

## Hello World Kubernetes Useful Command

docker build -t vsgdev/go-demo:0.1 .

docker run --rm -p 4000:4000 vsgdev/go-demo:0.1

docker push vsgdev/go-demo:0.1

kubectl apply -f go-demo-deployment.yaml

kubectl expose deployment go-demo-deployment --type=LoadBalancer --port=4000

minikube service go-demo-deployment
