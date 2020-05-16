pipeline {
  agent any
  stages {
    stage('checkout') {
      steps {
        git(url: 'https://github.com/r4bbitz/SayHello.git', branch: 'develop', changelog: true)
      }
    }

    stage('unit test') {
      steps {
        sh '''export GOROOT="/usr/local/go"
export PATH=$PATH:$GOROOT/bin
export GOPATH="/var/lib/jenkins/workspace/go"
go test -v'''
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

    stage('integrated test') {
      parallel {
        stage('integrated test') {
          steps {
            sh 'docker run -t postman/newman run "https://www.postman.com/collections/8a0c9bc08f062d12dcda"'
          }
        }

        stage('start app') {
          steps {
            sh '''export GOROOT="/usr/local/go"
export PATH=$PATH:$GOROOT/bin
export GOPATH="/var/lib/jenkins/workspace/go"




pwd'''
          }
        }

      }
    }

  }
}