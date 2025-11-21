pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'sanjith98/ip-service:latest'
        K8S_NAMESPACE = 'ip-service'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Unit Tests') {
            steps {
                sh 'go test -v ./...'
            }
        }

        stage('Build & Push Docker Image') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
                    sh 'docker build -t $DOCKER_IMAGE .'
                    sh 'echo $PASS | docker login -u $USER --password-stdin'
                    sh 'docker push $DOCKER_IMAGE'
                }
            }
        }

        stage('Deploy to Kubernetes Cluster') {
            steps {
                sh '''
                kubectl config use-context prod-cluster
                kubectl create namespace $K8S_NAMESPACE --dry-run=client -o yaml | kubectl apply -f -
                kubectl apply -f k8s/
                kubectl set image deployment/ip-service ip-service=$DOCKER_IMAGE -n $K8S_NAMESPACE
                kubectl rollout status deployment/ip-service -n $K8S_NAMESPACE
                echo "Deployed successfully!"
                echo "Live URL will be available after Ingress TLS cert is issued"
                '''
            }
        }
    }

    post {
        success {
            echo 'Jenkins CI/CD pipeline completed successfully!'
            echo 'Your service is now live with TLS (Let's Encrypt) on the Ingress host'
        }
    }
}
