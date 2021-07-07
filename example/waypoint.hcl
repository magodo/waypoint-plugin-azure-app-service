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
      resource_group_name = "acctestRG-mgdappservice"
      app_service_name = "acctestAS-123"
    }
  }

  /* release { */
  /*     use "azure-app-service" { */
  /*     } */
  /* } */
}
