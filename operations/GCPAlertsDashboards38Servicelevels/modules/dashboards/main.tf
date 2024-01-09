resource "google_monitoring_dashboard" "dashboard_simple_demo" {
  dashboard_json = <<EOF
{
  "displayName": "Demo Dashboard",
  "gridLayout": {
    "widgets": [
      {
        "blank": {}
      }
    ]
  }
}

EOF
}

resource "google_monitoring_dashboard" "dashboard_grid_layout_demo" {
  dashboard_json = <<EOF
{
  "dashboardFilters": [],
  "displayName": "Apache2 Demo Dashboard with Ops Agent Metric",
  "mosaicLayout": {
    "columns": 4,
    "tiles": [
      {
        "height": 2,
        "widget": {
          "title": "workload/apache.workers [SUM]",
          "xyChart": {
            "chartOptions": {
              "mode": "COLOR"
            },
            "dataSets": [
              {
                "minAlignmentPeriod": "60s",
                "plotType": "LINE",
                "targetAxis": "Y1",
                "timeSeriesQuery": {
                  "timeSeriesFilter": {
                    "aggregation": {
                      "alignmentPeriod": "60s",
                      "perSeriesAligner": "ALIGN_SUM"
                    },
                    "filter": "metric.type=\"workload.googleapis.com/apache.workers\" resource.type=\"gce_instance\"",
                    "secondaryAggregation": {
                      "alignmentPeriod": "60s",
                      "perSeriesAligner": "ALIGN_NONE"
                    }
                  }
                }
              }
            ],
            "thresholds": [],
            "timeshiftDuration": "0s",
            "yAxis": {
              "label": "",
              "scale": "LINEAR"
            }
          }
        },
        "width": 2
      }
    ]
  }
}

EOF
}