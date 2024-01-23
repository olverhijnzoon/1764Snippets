apply:
	terraform apply

destroy:
	terraform destroy

prep:
# 	admin first
# 	export GOOGLE_APPLICATION_CREDENTIALS

# 	service account
	if ! gcloud iam service-accounts describe snippets1764@snippets1764.iam.gserviceaccount.com > /dev/null 2>&1; then gcloud iam service-accounts create snippets1764 --description="snippets1764" --display-name="snippets1764"; fi

# 	create key
	if [ ! -f ./key.json ]; then gcloud iam service-accounts keys create ./key.json --iam-account snippets1764@snippets1764.iam.gserviceaccount.com; fi

# 	set permissions and use key
	gcloud projects add-iam-policy-binding snippets1764 --member=serviceAccount:snippets1764@snippets1764.iam.gserviceaccount.com --role=roles/storage.admin
	gcloud projects add-iam-policy-binding snippets1764 --member=serviceAccount:snippets1764@snippets1764.iam.gserviceaccount.com --role=roles/compute.admin

# 	use key
#	export GOOGLE_APPLICATION_CREDENTIALS=./key.json

prep_terraform:
# 	terraform state bucket
	if ! gsutil ls -b gs://1764snippets-state-bucket; then gsutil mb -l "EU" gs://1764snippets-state-bucket; fi

	terraform init