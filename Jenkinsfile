// pipeline {
//     agent any

//     environment {
//         // Define environment variables if needed
//         DOCKER_IMAGE_NAME = "go-sample-api"
//     }

//     stages {
//         stage('Build Container Image') {
//             steps {
//                 script {
//                     // Build the Docker image
//                     sh "docker build -t $DOCKER_IMAGE_NAME ."
//                 }
//             }
//         }

//         stage('Deploy') {
//             steps {
//                 script {
//                     // Deploy via docker-compose
//                     sh "docker rm -f $DOCKER_IMAGE_NAME && docker-compose up -d"
//                 }
//             }
//         }

//         stage('Cleanup') {
//             steps {
//                 script {
//                     // Clean up, remove the local Docker image
//                     sh "docker rmi $DOCKER_IMAGE_NAME"
//                 }
//             }
//         }
//     }

//     post {
//         success {
//             echo 'Build and deployment succeeded!'
//         }
//         failure {
//             echo 'Build or deployment failed!'
//         }
//     }
// }

pipeline {
    agent any

    stages {
        stage("Helloow") {
            steps {
                echo "Hello Pipeline"
            }
        }
    }

    post {
        always {
            echo "I always called"
        }
        success {
            echo "Succeed bruh"
        }
        failure {
            echo "Nuf, failed :("
        }
        Cleanup {
            echo "I always called too, but afterwards instead"
        }
    }
}