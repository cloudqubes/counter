pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        echo 'Building container images'
        sh """docker build . """
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
