pipeline {
    agent any

    environment {
        VERSION = ''
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    def branchName = env.GIT_BRANCH.split('/')[1]

                    if (branchName == 'features') {
                        VERSION = '0.0.1'
                    } else if (branchName == 'main') {
                        VERSION = '1.0.0'
                    } else {
                        VERSION = '0.1.0'
                    }

                    echo "Building version: ${VERSION}"
                }
                checkout scm
            }
        }

        stage('Build') {
            steps {
                script {
                    // Здесь можно добавить команду для сборки Docker-образа
                    sh 'docker build -t myapp:${VERSION} .'
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    // Здесь добавь команды для тестирования, если нужно
                    echo 'Running tests...'
                    // Например: sh './run_tests.sh'
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    // Разворачиваем контейнер в Docker
                    sh """
                    docker stop myapp || true
                    docker rm myapp || true
                    docker run -d --name myapp -p 8080:8080 myapp:${VERSION}
                    """
                }
            }
        }

        stage('Publish') {
            steps {
                script {
                    // Тут можно публиковать артефакты или пушить образ в Docker Hub или Registry
                    echo "Pushing Docker image to registry..."
                    sh "docker push myapp:${VERSION}"
                }
            }
        }
    }

    post {
        success {
            echo "Build and deploy succeeded!"
        }
        failure {
            echo "Build failed. Check logs for errors."
        }
    }
}
