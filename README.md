# hyvee-vaccine-tracker
This is a small project that scans Hy-Vee's API for new vaccination appointments.

## Run Your Own Tracker
First, you'll need to clone this project.

After that, you'll need to configure Twitter OAuth credentials. The Twitter account you use will have to be enabled for Twitter Developer access.

In `/deliverers/tweet/twitter.go`, insert your credentials by replacing the corresponding keys and tokens here:
```go
oauthConfig := oauth1.NewConfig("consumerKey", "consumerSecret")
token := oauth1.NewToken("accessToken", "accessSecret")
```

Next, configure your searchable area.

In the `updatePharmacies()` function in `main.go`, update the searchParam variable to your desired radius, latitude, and longitude. Here's an example:

```go
searchParams := api.Variables{
		Radius:    75,
		Latitude:  41.2354329,
		Longitude: -95.99383390000001,
	}

```

Finally, build and run the project, and you should see tweets as new vaccine appointments become available!
