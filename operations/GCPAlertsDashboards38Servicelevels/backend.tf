terraform {
  backend "gcs" {
    bucket  = "42snippets-state-bucket"
    prefix  = "terraform/state"
  }
}