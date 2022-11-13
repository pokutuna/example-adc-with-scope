example-adc-with-scope
===

Example of accessing Google Spreadsheet using Application Default Credentials.

See [How Application Default Credentials works](https://cloud.google.com/docs/authentication/application-default-credentials?hl=en).


## on local

The client library refers to `~/.config/gcloud/application_default_credentials.json` to authenticate.

`$ make application-default-credentials` puts the file.


## on Google Cloud (Cloud Functions)

The client library get a credentials of service account attached on runtime from metadata server.

[google-auth-library](https://www.npmjs.com/package/google-auth-library)(JavaScript) requires to pass `scopes` to get `GoogleAuth` object accessible to spreadsheet.

[google.golang.org/api](https://pkg.go.dev/google.golang.org/api)(Go) doesn't need that. `Service` resolves scopes internally.
