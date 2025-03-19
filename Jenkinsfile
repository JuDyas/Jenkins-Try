pipeline {
    agent any

    environment {
        VERSION = ''
        MAJOR_VERSION = '1'
        MINOR_VERSION = '0'
        PATCH_VERSION = '0'
    }

    stages { // <-- Единственная секция для стадий
        stage('Checkout') {
        steps {
            script {
                def branchName = env.BRANCH_NAME
                    if (branchName == 'main') {
                    // Подсчёт версий
                        MAJOR_VERSION = sh(script: "git rev-list --count origin/main || echo '0'", returnStdout: true).trim()
                        echo "MAJOR_VERSION: ${MAJOR_VERSION}"

                        PATCH_VERSION = sh(script: """
                        git for-each-ref --format='%(refname:short)' refs/remotes/origin/feature/* | wc -l || echo '0'
                        """, returnStdout: true).trim()
                        echo "PATCH_VERSION: ${PATCH_VERSION}"

                        MINOR_VERSION = sh(script: "git rev-list --count origin/develop || echo '0'", returnStdout: true).trim()
                        echo "MINOR_VERSION: ${MINOR_VERSION}"

                        // Формирование итоговой версии
                        VERSION = "${MAJOR_VERSION}.${MINOR_VERSION}.${PATCH_VERSION}"
                        echo "VERSION: ${VERSION}"

                        if (!VERSION?.trim()) {
                        error("VERSION is empty! MAJOR_VERSION=${MAJOR_VERSION}, MINOR_VERSION=${MINOR_VERSION}, PATCH_VERSION=${PATCH_VERSION}")
                        }
                    }
                }
                checkout scm
            }
        }

        stage('Check Docker') {
        steps {
            script {
                sh 'docker --version'
                }
            }
        }

        stage('Build') {
        steps {
            script {
                echo "Building Docker image with VERSION: ${VERSION}"

                    // Проверяем, что VERSION не пустая
                    if (!VERSION?.trim()) {
                    error("VERSION is empty! Unable to proceed with Docker build.")
                    }

                    sh "docker build -t myapp:${VERSION} ."
                }
            }
        }

        stage('Test') {
        steps {
            script {
                echo 'Running tests...'
                }
            }
        }

        stage('Deploy') {
        steps {
            script {
                echo "Deploying Docker image..."
                    sh """
                    docker stop myapp || true
                    docker rm myapp || true
                    docker run -d --name myapp -p 8082:8082 myapp:${VERSION}
                    """
                }
            }
        }
    }

    post {
        success {
            echo "Build and deploy succeeded!"
        }
        failure {
            echo "Build failed."
        }
    }
}