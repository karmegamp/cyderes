Please select "Code" view before reading this document

Prerequest
==========
git clone git@github.com:karmegamp/cyderes-deploy.git

A. Continueous Deployment steps

1. Install ArgoCD 
=================
$ brew install argocd

2. Apply ArgoCD related resources to K8s
========================================
$ kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

3. Create argocd namespace
===========================
$ kubectl create namespace argocd

4. Set the context
==================
$ kubectl config set-context --current --namespace=argocd

5. Login argoCD CLI 
====================
$ argocd login --core

6. Create and link argoCD application with GitHub repository for K8s resource sync
==================================================================================
$ argocd app create cyderes-deploy --repo https://github.com/karmegamp/cyderes-deploy.git --path cyderes-deploy --dest-server https://kubernetes.default.svc --dest-namespace default

7. Once K8s resource updated in GitHub, update the cluster 
===============================================================
$ argocd app sync cyderes-deploy

Example:


Before Git update for datatx resources
--------------------------------------
$kubectl get svc,deploy,pod

NAME                                       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)             AGE
service/argocd-applicationset-controller   ClusterIP   10.111.189.149   <none>        7000/TCP,8080/TCP   5h5m
service/argocd-metrics                     ClusterIP   10.104.80.13     <none>        8082/TCP            5h5m
service/argocd-redis                       ClusterIP   10.111.93.199    <none>        6379/TCP            5h5m
service/argocd-repo-server                 ClusterIP   10.103.235.30    <none>        8081/TCP,8084/TCP   5h5m

NAME                                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/argocd-applicationset-controller   1/1     1            1           5h5m
deployment.apps/argocd-redis                       1/1     1            1           5h5m
deployment.apps/argocd-repo-server                 1/1     1            1           5h5m

NAME                                                    READY   STATUS    RESTARTS   AGE
pod/argocd-application-controller-0                     1/1     Running   0          5h5m
pod/argocd-applicationset-controller-65dd5b65c6-t2qm7   1/1     Running   0          5h5m
pod/argocd-redis-6d87b6f77d-ptltp                       1/1     Running   0          5h5m
pod/argocd-repo-server-546c4dd678-pmc64                 1/1     Running   0          5h5m

GitHub update and sync the cluster 
----------------------------------

 $ argocd app sync cyderes-deploy
{"level":"info","msg":"unknown field \"status.sourceHydrator\"","time":"2025-06-29T07:34:33+05:30"}
TIMESTAMP                  GROUP        KIND   NAMESPACE                  NAME    STATUS    HEALTH        HOOK  MESSAGE
2025-06-29T07:34:33+05:30            Service      argocd                datatx  OutOfSync  Missing
2025-06-29T07:34:33+05:30   apps  Deployment      argocd                datatx  OutOfSync  Missing
2025-06-29T07:34:33+05:30            Service      argocd                datatx  OutOfSync  Healthy
2025-06-29T07:34:33+05:30            Service      argocd                datatx  OutOfSync  Healthy              service/datatx created
2025-06-29T07:34:33+05:30   apps  Deployment      argocd                datatx  OutOfSync  Missing              deployment.apps/datatx created
2025-06-29T07:34:33+05:30   apps  Deployment      argocd                datatx  OutOfSync  Progressing              deployment.apps/datatx created

Name:               argocd/cyderes-deploy
Project:            default
Server:             https://kubernetes.default.svc
Namespace:          default
URL:                http://localhost:58260/applications/cyderes-deploy
Source:
- Repo:             https://github.com/karmegamp/cyderes-deploy.git
  Target:
  Path:             cyderes-deploy
SyncWindow:         Sync Allowed
Sync Policy:        Manual
Sync Status:        OutOfSync from  (ddd95fa)
Health Status:      Progressing

Operation:          Sync
Sync Revision:      ddd95faa94643f6a8aa4c32df073446b7c1e48fd
Phase:              Succeeded
Start:              2025-06-29 07:34:33 +0530 IST
Finished:           2025-06-29 07:34:33 +0530 IST
Duration:           0s
Message:            successfully synced (all tasks run)

GROUP  KIND        NAMESPACE  NAME    STATUS     HEALTH       HOOK  MESSAGE
       Service     argocd     datatx  OutOfSync  Healthy            service/datatx created
apps   Deployment  argocd     datatx  OutOfSync  Progressing        deployment.apps/datatx created

After Git update of datatx resources
------------------------------------
$ kubectl get svc,deploy,pod
NAME                                       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)             AGE
service/argocd-applicationset-controller   ClusterIP   10.111.189.149   <none>        7000/TCP,8080/TCP   5h3m
service/argocd-metrics                     ClusterIP   10.104.80.13     <none>        8082/TCP            5h3m
service/argocd-redis                       ClusterIP   10.111.93.199    <none>        6379/TCP            5h3m
service/argocd-repo-server                 ClusterIP   10.103.235.30    <none>        8081/TCP,8084/TCP   5h3m
service/datatx                             NodePort    10.111.164.73    <none>        8888:30279/TCP      39m

NAME                                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/argocd-applicationset-controller   1/1     1            1           5h3m
deployment.apps/argocd-redis                       1/1     1            1           5h3m
deployment.apps/argocd-repo-server                 1/1     1            1           5h3m
deployment.apps/datatx                             2/2     2            2           39m

NAME                                                    READY   STATUS    RESTARTS   AGE
pod/argocd-application-controller-0                     1/1     Running   0          5h3m
pod/argocd-applicationset-controller-65dd5b65c6-t2qm7   1/1     Running   0          5h3m
pod/argocd-redis-6d87b6f77d-ptltp                       1/1     Running   0          5h3m
pod/argocd-repo-server-546c4dd678-pmc64                 1/1     Running   0          5h3m
pod/datatx-84cdfc479f-8hpkq                             1/1     Running   0          39m
pod/datatx-84cdfc479f-t22sm                             1/1     Running   0          39m


