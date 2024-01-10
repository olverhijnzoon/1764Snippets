terraform {
  required_version = "~> 1.5.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.40"
    }
  }
}

provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
}

resource "google_storage_bucket" "terraform_state" {
  name     = "42snippets-state-bucket"
  location = "EU"
}

resource "google_service_account" "ops_agent" {
  account_id   = "ops-agent-service-account"
  display_name = "Service Account for Google Cloud Ops Agent"
}

resource "google_project_iam_member" "monitoring_writer" {
  project = var.project
  role    = "roles/monitoring.metricWriter"
  member  = "serviceAccount:${google_service_account.ops_agent.email}"
}

resource "google_project_iam_member" "monitoring_viewer" {
  project = var.project
  role    = "roles/monitoring.viewer"
  member  = "serviceAccount:${google_service_account.ops_agent.email}"
}

resource "google_project_iam_member" "logs_writer" {
  project = var.project
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.ops_agent.email}"
}

resource "google_compute_instance" "default" {
  name         = "snippets42"
  machine_type = var.machine_type
  tags         = ["http-server"]

  service_account {
    email  = google_service_account.ops_agent.email
    scopes = ["cloud-platform"]
  }

  boot_disk {
    initialize_params {
      image = "debian-11"
    }
  }

  network_interface {
    network = "default"
    # Enable external IP
    access_config {}
  }

  metadata_startup_script = <<SCRIPT
  apt-get update
  apt-get install -y apache2 php7.0
  systemctl start apache2
  curl -sSO https://dl.google.com/cloudagents/add-google-cloud-ops-agent-repo.sh
  sudo bash add-google-cloud-ops-agent-repo.sh --also-install

  # Create a configuration file for Ops Agent
  cat > /etc/google-cloud-ops-agent/config.yaml <<EOL
  logging:
    receivers:
      apache_access:
        type: apache_access
      apache_error:
        type: apache_error
    service:
      pipelines:
        apache:
          receivers:
            - apache_access
            - apache_error
  metrics:
    receivers:
      apache:
        type: apache
    service:
      pipelines:
        apache:
          receivers:
            - apache
  EOL

  # Restart the Ops Agent to apply the new configuration
  sudo service google-cloud-ops-agent restart
SCRIPT
}

resource "google_compute_firewall" "default" {
  name    = "apache-allow-http"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  target_tags = ["http-server"]
  source_ranges = ["0.0.0.0/0"]
}

resource "google_monitoring_uptime_check_config" "default" {
  display_name = "uptime-check"
  timeout       = "10s"
  period        = "60s"

  http_check {
    path = "/"
    port = "80"
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "snippets-390309"
      # Reference the external IP of the compute instance
      host       = google_compute_instance.default.network_interface[0].access_config[0].nat_ip
    }
  }

  # Explicitly specify dependency on the Compute Instance
  depends_on = [google_compute_instance.default]
}

module "alerts" {
  source = "./modules/alerts"
  email_address = var.email_address
}

module "dashboards" {
  source = "./modules/dashboards"
}

resource "google_monitoring_custom_service" "apache_service" {
  service_id = "apache-service"
  display_name = "Apache Web Server"

  # Define the service's custom resource type.
  # This example uses 'gce_instance' as the resource type for the Nginx web server.
  telemetry {
    resource_name = "gce_instance"
  }
}

module "servicelevels" {
  source      = "./modules/servicelevels"
  goal        = 0.99 # For example, 99% availability
  service     = "apache-service"
  slo_id      = "apache-availability-slo"
}