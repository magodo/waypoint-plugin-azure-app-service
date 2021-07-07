project = "example"

app "example" {
  build {
    use "docker-pull" {
        image = "nginxdemos/hello"
        #image = "yeasy/simple-web"
        tag = "latest"
    }
  }
  deploy {
    use "azure-app-service" {
        app_service_id = "/subscriptions/67a9759d-d099-4aa8-8675-e6cfd669c3f4/resourceGroups/acctestRG-mgdappservice/providers/Microsoft.Web/sites/acctestAS-123"
    }
  }

  /* release { */
  /*     use "azure-app-service" { */
  /*     } */
  /* } */
}
