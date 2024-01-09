variable "goal" {
  description = "The target value of the SLO."
  type        = number
}

variable "service" {
  description = "The name of the custom service."
  type        = string
}

variable "slo_id" {
  description = "The ID of the Service Level Objective."
  type        = string
}