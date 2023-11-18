# The way to access dashboard in gg cloud

## In virtual machine
<code>kubectl proxy</code>

## In local machine
<code>
gcloud compute ssh --zone "asia-southeast1-a" "qthuy2609@instance-2" --project "thesis-course-registration" --ssh-flag="-L 8001:localhost:8001"
</code>

### Page
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/pod?namespace=thesis-management-backend

## Get token
kubectl -n kubernetes-dashboard create token admin-user