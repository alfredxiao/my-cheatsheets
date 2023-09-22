# install
`curl https://baltocdn.com/helm/signing.asc | sudo apt-key add -`
`echo "deb https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list`
`sudo apt-get update`
`sudo apt-get install -y helm`

# Helm Repository
Add a Helm chart repository.
- `helm repo add bitnami https://charts.bitnami.com/bitnami`
Update the repository
- `helm repo update`
Search
- `helm search repo bitnami`

# Intall a release
- `helm install -n myspace myapp myrepo/mychart --set extraValue=2`

# Uninstall
`helm -n mercury uninstall internal-issue-report-apache`

`helm show readme bitnami/apache`
`helm show values bitnami/apache`

# List releases
`helm list -a -A`
