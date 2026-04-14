# Test technique Merkl - DevOps Intern

Bonjour ! Voici mon rendu pour le challenge technique DevOps. J'ai bien structuré mon travail en deux parties, comme défini dans les consignes.

## Partie 1 : Les manifests Kubernetes

Pour cette première partie, j'ai préparé les définitions classiques (Namespace, Deployment, ConfigMap, Quotas, Probes). 

Pour tout déployer facilement sur votre cluster de test, il suffit de lancer :

```bash
kubectl apply -f task1-workload.yaml
kubectl apply -f task2-configmap.yaml
kubectl apply -f task3-resources.yaml
kubectl apply -f task4-probes.yaml
```

## Partie 2 : Le Custom Kubernetes Operator (HelloApp)

J'ai utilisé `operator-sdk` pour générer le projet dans `hello-operator`.

> ** Fix technique** : Le scaffold de base utilisait une vieille version de `controller-tools` (v0.12.0) incompatible avec Go 1.25. J'ai mis à jour le `Makefile` vers la `v0.16.5` pour corriger les erreurs de compilation.

### Implémentation :
- **L'API** (`helloapp_types.go`) : Ajout des champs `Message` et `Replicas`, puis ajout de `AvailableReplicas` dans le status.
- **La boucle Reconcile** (`helloapp_controller.go`) : J'ai fait très simple. On récupère la ressource `HelloApp` et on gère sa suppression silencieusement via un `IgnoreNotFound`.
