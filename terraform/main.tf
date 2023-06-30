resource "google_cloud_run_service" "api_service" {
  name     = "go-orders"
  location = ""

  template {
    spec {
      containers {
        image = "gcr.io/ics-jamii-1632986730022/go-order-api"

        env {
          name  = "FIRESTORE_PROJ_ID"
          value = var.firestore_proj_id
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
