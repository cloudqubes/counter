pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        echo 'Building container images'
        sh """docker build -t counter:1.0.0 ."""
      }
    }
    stage('Publish') {
      steps {
        echo 'Deploying....'
      }
    }
    stage('Deploy') {
      steps {
        echo 'Deploying....'
      }
    }
  }
}
