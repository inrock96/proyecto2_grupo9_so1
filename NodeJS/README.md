docker login
docker build -t webpage .
docker tag pagina bernaldrpp/pagina
docker push bernaldrpp/pagina

kubectl create deployment pagina -n proyecto --image=bernaldrpp/pagina
kubectl -n proyecto expose deployment pagina --port 3000 --target-port=3000 --type NodePort --name=pagina-svc