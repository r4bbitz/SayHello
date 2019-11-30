pipeline {
  agent {
    node {
      label 'go'
    }

  }
  stages {
    stage('checkout') {
      steps {
        git(url: 'https://github.com/r4bbitz/SayHello.git', branch: 'develop', changelog: true)
      }
    }

    stage('build') {
      steps {
        sh 'go build '
      }
    }

    stage('run') {
      steps {
        sh './SayHello'
      }
    }

  }
  environment {
    go = 'go'
  }
}