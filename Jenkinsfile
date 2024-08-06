pipeline {
    agent any

    environment {
        GO_VERSION = '1.19'
    }

    stages {
        stage('Setup') {
            steps {
                echo 'Setting up...'
                sh 'go version'
            }
        }
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go build -o my-go-app'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'go test ./...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
                // Add your deployment steps here
                // For example, copying the binary to a server
                // sh 'scp my-go-app user@server:/path/to/deploy'
            }
        }
    }
    post {
        always {
            echo 'Cleaning up...'
            sh 'rm -f my-go-app'
        }
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
