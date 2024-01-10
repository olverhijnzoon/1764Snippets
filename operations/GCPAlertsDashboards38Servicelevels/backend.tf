terraform {
  backend "gcs" {
    bucket  = "1764snippets-state-bucket"
    prefix  = "terraform/state"
  }
}