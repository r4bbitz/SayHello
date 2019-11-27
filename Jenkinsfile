pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        git(url: 'https://github.com/r4bbitz/SayHello.git', branch: 'develop', changelog: true)
      }
    }

  }
}