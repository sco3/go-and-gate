resource "kubernetes_deployment" "and_gate" {
  metadata {
    name = "and-gate-server"
    labels = {
      app = "and-gate"
    }
  }

  spec {
    replicas = 3

    selector {
      match_labels = {
        app = "and-gate"
      }
    }

    template {
      metadata {
        labels = {
          app = "and-gate"
        }
      }

      spec {
        container {
          image             = "and-gate-server:latest"
          image_pull_policy = "IfNotPresent"
          name              = "server"
          args              = ["-start", "-address", ":5555"]

          port {
            container_port = 5555
            name           = "grpc"
          }

          resources {
            limits = {
              cpu    = "500m"
              memory = "128Mi"
            }
            requests = {
              cpu    = "100m"
              memory = "64Mi"
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "and_gate" {
  metadata {
    name = "and-gate-service"
  }

  spec {
    selector = {
      app = kubernetes_deployment.and_gate.spec[0].template[0].metadata[0].labels.app
    }

    port {
      port        = 5555
      target_port = 5555
      name        = "grpc"
    }

    type = "ClusterIP"
  }
}