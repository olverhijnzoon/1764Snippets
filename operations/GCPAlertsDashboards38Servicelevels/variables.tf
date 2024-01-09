variable "project" {
  description = "The ID of the Google Cloud project."
  type        = string
}

variable "region" {
  description = "Default Region"
  type        = string
  default     = "us-central1"
}

variable "zone" {
  description = "Default Zone"
  type        = string
  default     = "us-central1-a"
}

variable "machine_type" {
  description = "Machine Type"
  type        = string
  default     = "e2-micro"
}

variable "email_address" {
  description = "The email address of the person being notified."
  type        = string
}