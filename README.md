# rwgps-analysis

This repository defines a Google Cloud Function to serve RideWithGPS route analysis via HTTP. The hard work is done by [gpx-utils](https://github.com/ray1729/gpx-utils/).

## Testing Locally

We use the Google Cloud Platform [Functions Framework for Go](https://github.com/GoogleCloudPlatform/functions-framework-go) to test locally:

    cd cmd
    go build
    PORT=3000 ./cmd

Then query a route:

    curl 'http://localhost:3000/rwgps?routeId=32413693&stops=ctccambridge'

## Deploying

See [Before you begin](https://cloud.google.com/functions/docs/quickstart-go#before-you-begin) for instructions on setting up a Google Cloud project. This must be done before deploying the cloud function.

See [Deploying Cloud Functions](https://cloud.google.com/functions/docs/deploying) for the different options. Our chosen method is to [deploy from source control](https://cloud.google.com/functions/docs/deploying/repo).

Creating the source repository and configuration mirroring must be done through the GCloud console as described in the [mirrored repositories guide](https://cloud.google.com/tools/cloud-repositories/docs/cloud-repositories-hosted-repository).

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