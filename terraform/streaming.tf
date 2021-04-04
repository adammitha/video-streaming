locals {
  service_name = "streaming"
  image_tag = "gcr.io/${var.project}/${local.service_name}:${var.app_version}"
}