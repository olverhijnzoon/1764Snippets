resource "google_monitoring_slo" "apache_idle_workers_slo" {
  service = var.service
  slo_id  = "apache-idle-workers-slo"
  goal    = 0.95  # 95% of the time, the number of idle workers should be below 50.

  rolling_period_days = 7

windows_based_sli {
    window_period = "60s"
    metric_sum_in_range {
      time_series = join(" AND ", [
        # There seems to be a new API now and this did not work on first attempt but should be close
        "metric.type=\"agent.googleapis.com/apache/idle_workers\"",
        "resource.label.instance_id=\"4284582424701292929\"",
      ])
      range {
        min = 50
        max = 60
      }
    }
  }
}