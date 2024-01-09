resource "google_monitoring_notification_channel" "default" {
  display_name = "notification-channel"
  type         = "email"
  labels = {
    email_address = var.email_address
  }
}

resource "google_monitoring_alert_policy" "default" {
  display_name = "42snippets-alert-policy"
  combiner     = "OR"

  conditions {
    display_name = "uptime-check"
    condition_threshold {
      filter     = "metric.type=\"monitoring.googleapis.com/uptime_check/check_passed\" resource.type=\"uptime_url\""
      comparison = "COMPARISON_LT"
      threshold_value = 1
      duration   = "60s"

      aggregations {
        alignment_period   = "60s"
        per_series_aligner = "ALIGN_NEXT_OLDER"
      }
    }
  }

  notification_channels = [
    google_monitoring_notification_channel.default.name
  ]
}