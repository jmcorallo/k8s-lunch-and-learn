# k8s-lunch-and-learn
Resources for Kubernetes Lunch and Learn presentation

Runbook:

# Create cluster

bucket_name=lunch-and-learn-cluster

aws s3api create-bucket --bucket ${bucket_name} --region us-east-1

export KOPS_CLUSTER_NAME=demo.k8s.local

export KOPS_STATE_STORE=s3://${bucket_name}

kops create cluster --node-count=3 --node-size=t2.micro --zones=us-east-1a --master-size=t2.micro --master-count=1
kops update cluster --yes