pipeline {
  agent any
  stages {
    stage('checkout') {
      steps {
        git(url: 'https://github.com/r4bbitz/SayHello.git', branch: 'develop', changelog: true)
      }
    }

    stage('build') {
      steps {
        sh '''export GOROOT="/usr/local/go"
export PATH=$PATH:$GOROOT/bin
export GOPATH="/var/lib/jenkins/workspace/go"

go build'''
      }
    }

    stage('Test') {
      steps {
        sh 'go test'
      }
    }

  }
}