pipeline {
    agent any

    environment {
        // Задайте переменные среды: название образа и базовая версия
        DOCKER_IMAGE = 'my-golang-app'
        BASE_VERSION = '1.0.0'
        DOCKER_HOST = "unix:///var/run/docker.sock"
    }

    stages {
        stage('Checkout') {
            steps {
                sh 'docker ps'
                // Клонирование репозитория
                checkout scm
            }
        }

        stage('Calculate Version') {
            steps {
                script {
                    // Подсчет коммитов в ветках
                    def featureCommits = sh(script: "git rev-list origin/feature --count", returnStdout: true).trim()
                    def mainCommits = sh(script: "git rev-list origin/main --count", returnStdout: true).trim()
                    def developCommits = sh(script: "git rev-list origin/develop --count", returnStdout: true).trim()

                    // Расчет версии (BASE_VERSION + main + разработка)
                    def calculatedVersion = "${env.BASE_VERSION}-${mainCommits}.${developCommits}.${featureCommits}"

                    echo "Calculated version: ${calculatedVersion}"

                    // Сохранение версии
                    env.APP_VERSION = calculatedVersion
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Сборка Docker-образа и тегирование
                    sh "docker build -t ${env.DOCKER_IMAGE}:${env.APP_VERSION} ."
                }
            }
        }

        stage('Test in Builder') {
            steps {
                script {
                    // Собираем промежуточный "builder" образ, где выполняются тесты
                    sh "docker build -t builder-test --target builder -f Dockerfile ."

                    // Запуск тестов в "builder" образе
                    sh "docker run --rm builder-test go test ./..."
                }
            }
        }


        stage('Deploy') {
            steps {
                script {
                    // Деплой контейнера
                    sh "docker run -d -p 8080:8080 --name app-container ${env.DOCKER_IMAGE}:${env.APP_VERSION}"

                    echo "Application deployed successfully. Running version: ${env.APP_VERSION}"
                }
            }
        }
    }

    post {
        // Очистка окружения после сборки
        always {
            script {
                // Удаление ненужных Docker-образов
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