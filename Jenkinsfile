pipeline {
    agent any

    environment {
        // Define environment variables if needed
        DOCKER_IMAGE_NAME = "my-golang-app"
        DOCKER_REGISTRY_URL = "docker.io"
        DOCKER_REGISTRY_USERNAME = credentials('docker-hub-username')
        DOCKER_REGISTRY_PASSWORD = credentials('docker-hub-password')
    }

    stages {
        stage('Build') {
            steps {
                script {
                    // Build the Go application
                    sh 'go build -o myapp'
                }
            }
        }

        stage('Dockerize') {
            steps {
                script {
                    // Build the Docker image
                    sh "docker build -t $DOCKER_REGISTRY_URL/$DOCKER_REGISTRY_USERNAME/$DOCKER_IMAGE_NAME ."
                }
            }
        }

        stage('Push to Docker Registry') {
            steps {
                script {
                    // Push the Docker image to the registry
                    sh "docker login -u $DOCKER_REGISTRY_USERNAME -p $DOCKER_REGISTRY_PASSWORD $DOCKER_REGISTRY_URL"
                    sh "docker push $DOCKER_REGISTRY_URL/$DOCKER_REGISTRY_USERNAME/$DOCKER_IMAGE_NAME"
                }
            }
        }

        stage('Cleanup') {
            steps {
                script {
                    // Clean up, remove the local Docker image
                    sh "docker rmi $DOCKER_IMAGE_NAME"
                }
            }
        }
    }

    post {
        success {
            echo 'Build and deployment succeeded!'
        }
        failure {
            echo 'Build or deployment failed!'
        }
    }
}
