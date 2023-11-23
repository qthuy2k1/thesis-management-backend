# The way to access dashboard in gg cloud

## In virtual machine
<code>kubectl proxy</code>

## In local machine
<code>
ssh -L 9001:103.82.132.17:8001 new_user@103.82.132.1
</code>

### Page
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/pod?namespace=thesis-management-backend

## Get token
kubectl -n kubernetes-dashboard create token admin-user