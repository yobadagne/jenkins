pipeline {
    agent any

    environment {
        GO_VERSION = '1.19'
    }

    stages {
         stage('Checkout') {
            steps {
                script {
                    withCredentials([string(credentialsId: 'github-token', variable: 'GITHUB_TOKEN')]) {
                        git url: 'https://$GITHUB_TOKEN@github.com/yobadagne/jenkins.git', credentialsId: 'github-token'
                    }
                }
            }
        }
        stage('Setup') {
            steps {
                echo 'Setting up...'
                sh 'go version'
            }
        }
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go build -o my-go-app cmd/main.go'
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
            // sh 'rm -f my-go-app'
        }
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
