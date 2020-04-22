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

See [Before you begin](https://cloud.google.com/functions/docs/quickstart-go#before-you-begin) for instructions on setting up the Google Cloud project. This must be done before deploying the cloud function.

See [Deploying Cloud Functions](https://cloud.google.com/functions/docs/deploying) for the different options. Our chosen method is to [deploy from source control](https://cloud.google.com/functions/docs/deploying/repo).



