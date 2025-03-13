pipeline {
    agent any

    environment {
        VERSION = ''
        MAJOR_VERSION = '1'
        MINOR_VERSION = '0'
        PATCH_VERSION = '0'
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    def branchName = env.BRANCH_NAME
                    if (branchName == 'main') {
                        MAJOR_VERSION = sh(script: "git rev-list --count origin/main", returnStdout: true).trim()
                        PATCH_VERSION = sh(script: "git rev-list --count origin/feature/*", returnStdout: true).trim()
                        MINOR_VERSION = sh(script: "git rev-list --count origin/develop", returnStdout: true).trim()

                        VERSION = "${MAJOR_VERSION}.${MINOR_VERSION}.${PATCH_VERSION}"

                        echo "Building version: ${VERSION}"
                    }
                }
                checkout scm
            }
        }

        stage('Build') {
            steps {
                script {
                    sh 'docker build -t myapp:${VERSION} .'
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
