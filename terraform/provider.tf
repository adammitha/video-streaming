provider "google" {
  project     = "video-streaming-306005"
  region      = "us-west1"
  zone        = "us-west1-c"
}

provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "gke_video-streaming-306005_us-west1_video-streaming"
}