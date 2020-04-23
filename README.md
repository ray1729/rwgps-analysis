# rwgps-analysis

This repository defines a Google Cloud Function to serve RideWithGPS route analysis via HTTP. The hard work is done by [gpx-utils](https://github.com/ray1729/gpx-utils/).

## Updating dependencies

As this is a thin wrapper around [gpx-utils](https://github.com/ray1729/gpx-utils/), we'll need to update the dependency to pull in new changes. To update to the latest version:

    go get -u github.com/ray1729/gpx-utils
    
Then commit and push the change to `go.mod` and `go.sum`.

## Testing Locally

We use the Google Cloud Platform [Functions Framework for Go](https://github.com/GoogleCloudPlatform/functions-framework-go) to test locally:

    cd cmd
    go build
    PORT=3000 ./cmd

Then query a route:

    curl 'http://localhost:3000/rwgps?routeId=32413693&stops=ctccambridge'

## Deploying

See [Before you begin](https://cloud.google.com/functions/docs/quickstart-go#before-you-begin) for instructions on setting up a Google Cloud project. This must be done before deploying the cloud function.

See [Deploying Cloud Functions](https://cloud.google.com/functions/docs/deploying) for the different options. 

### Deploying from Github

Our chosen method is to [deploy from source control](https://cloud.google.com/functions/docs/deploying/repo).

Creating the source repository and configuring mirroring must be done through the GCloud console as described in the [mirrored repositories guide](https://cloud.google.com/tools/cloud-repositories/docs/cloud-repositories-hosted-repository).

My project id is `master-chariot-275014` and I want to deploy from the `master` branch:

    gcloud functions deploy rwgps_analysis \
       --region europe-west2 \
       --source https://source.developers.google.com/projects/master-chariot-275014/repos/github_ray1729_rwgps-analysis/moveable-aliases/master/paths/ \
       --entry-point HandleRequest \
       --runtime go113 \
       --memory 512MB \
       --trigger-http \
       --allow-unauthenticated

This returns the trigger URL, which we can use to invoke the function:

    curl 'https://europe-west2-master-chariot-275014.cloudfunctions.net/rwgps_analysis?routeId=31822531&stops=ctccambridge'
    
### Deploying from Gitlab

We can use a Gitlab pipeline to deploy automatically from Gitlab. Note that the method described here does not need a mirrored repository. We use a service account to authenticate and grant permissions to Gitlab. See [Creating and managing service account keys](https://cloud.google.com/iam/docs/creating-managing-service-account-keys).

Create a service account for Gitlab:

    gcloud iam service-accounts create gitlab --display-name "Gitlab CI/CD"
    
Grant the service account editor permissions for the project:

    gcloud projects add-iam-policy-binding master-chariot-275014 \
        --member serviceAccount:master-chariot-275014@appspot.gserviceaccount.com \
        --role roles/editor
        
Create a service account key:

    gcloud iam service-accounts keys create key.josn --iam-account master-chariot-275014@appspot.gserviceaccount.com

Navigate to the Gitlab project and go to Settings > CI/CD. Expand the `Variables` section, and click `Add Variable`. Create a variable called `GCLOUD_SERVICE_KEY`, pasting in the value from `key.json` you created above. Select the checkbox `Protect variable` and save.

Create another variable called `GCLOUD_PROJECT` whose value is the name of your Google Cloud project. This one can also be protected.

Now navigate to Settings > Repository and expand the `Protected Branches` section. Verify that the `master` branch is protected, and has appropriate permissions. This is the branch we will deploy from.

 

      