# Getting started with Go on Bluemix
To get started, we'll take you through a sample Go hello world app, help you set up a development environment and deploy to Bluemix.

## Prerequisites

You'll need the following:
* [Bluemix account](https://console.ng.bluemix.net/registration/)
* [Cloud Foundry CLI](https://github.com/cloudfoundry/cli#downloads)
* [Git](https://git-scm.com/downloads)
* [Go](https://golang.org/dl/)

## 1. Clone the sample app

Now you're ready to start working with the simple Go *hello world* app. Clone the repository and change to the directory where the sample app is located.
  ```
git clone https://github.com/vshasha/gowebapp
cd gowebapp
  ```

Peruse the files in the *gowebapp* directory to familiarize yourself with the contents.

## 2. Run the app locally

Build and run the app.
  ```
go run goserver.go
  ```

View your app at: http://localhost:13100

## 3. Prepare the app for deployment


To deploy to Bluemix, it can be helpful to set up a manifest.yml file. One is provided for you with the sample. Take a moment to look at it.

The manifest.yml includes basic information about your app, such as the name, how much memory to allocate for each instance and the route. In this manifest.yml **random-route: true** generates a random route for your app to prevent your route from colliding with others.  You can replace **random-route: true** with **host: myChosenHostName**, supplying a host name of your choice. [Learn more...](https://console.bluemix.net/docs/manageapps/depapps.html#appmanifest)
 ```
 applications:
 - name: myFirstGoWebApp
   random-route: true
   memory: 128M
 ```

## 4. Deploy the app

You can use the Cloud Foundry CLI to deploy apps.

Choose your API endpoint
   ```
cf api <API-endpoint>
   ```

Replace the *API-endpoint* in the command with an API endpoint from the following list.

|URL                             |Region          |
|:-------------------------------|:---------------|
| https://api.ng.bluemix.net     | US South       |
| https://api.eu-de.bluemix.net  | Germany        |
| https://api.eu-gb.bluemix.net  | United Kingdom |
| https://api.au-syd.bluemix.net | Sydney         |

Login to your Bluemix account

  ```
cf login
  ```

From within the *get-started-go* directory push your app to Bluemix
  ```
cf push
  ```

This can take a minute. If there is an error in the deployment process you can use the command `cf logs <Your-App-Name> --recent` to troubleshoot.

When deployment completes you should see a message indicating that your app is running.  View your app at the URL listed in the output of the push command.  You can also issue the

 ```
cf apps
 ```
command to view your apps status and see the URL.
