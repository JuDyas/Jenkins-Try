pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'my-golang-app'
        DOCKER_HOST = "unix:///var/run/docker.sock"
    }

    stages {
        stage('Checkout') {
            steps {
                sh 'docker ps'
                checkout scm
            }
        }

        stage('Calculate Version') {
            steps {
                script {
                    def featureCommits = sh(script: "git rev-list origin/feature --count", returnStdout: true).trim()
                    def mainCommits = sh(script: "git rev-list origin/main --count", returnStdout: true).trim()
                    def developCommits = sh(script: "git rev-list origin/develop --count", returnStdout: true).trim()

                    def calculatedVersion = "${mainCommits}.${developCommits}.${featureCommits}"

                    echo "Calculated version: ${calculatedVersion}"
                    env.APP_VERSION = calculatedVersion
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh "docker build -t ${env.DOCKER_IMAGE}:${env.APP_VERSION} ."
                }
            }
        }

        stage('Test in Builder') {
            steps {
                script {
                    sh "docker build -t builder-test --target builder -f Dockerfile ."
                    sh "docker run --rm builder-test go test ./..."
                }
            }
        }


        stage('Deploy') {
            steps {
                script {
                    sh 'docker rm -f app-container || true'
                    sh "docker run -d -p 8081:8081 --name app-container ${env.DOCKER_IMAGE}:${env.APP_VERSION}"
                    echo "Application deployed successfully. Running version: ${env.APP_VERSION}"
                }
            }
        }
    }

    post {
        always {
            script {
                sh "docker image prune -f"
            }
        }
        success {
            echo 'Build completed successfully!'
        }
        failure {
            echo 'Build failed!'
        }
    }
}