resource "kubernetes_deployment" "redis" {
  metadata {
    name = "redis"
    labels = {
      app = "redis"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "redis"
      }
    }

    template {
      metadata {
        labels = {
          app = "redis"
        }
      }

      spec {
        container {
          image = "redis:alpine"
          name  = "redis"

          port {
            container_port = 6379
            name           = "redis"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "redis" {
  metadata {
    name = "redis-service"
  }

  spec {
    selector = {
      app = kubernetes_deployment.redis.spec[0].template[0].metadata[0].labels.app
    }

    port {
      port        = 6379
      target_port = 6379
      name        = "redis"
    }

    type = "ClusterIP"
  }
}