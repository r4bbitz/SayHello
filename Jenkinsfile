pipeline {
  agent any
  stages {
    stage('checkout') {
      steps {
        git(url: 'https://github.com/r4bbitz/SayHello.git', branch: 'develop', changelog: true)
      }
    }

    stage('build') {
      parallel {
        stage('build') {
          steps {
            sh '''export GOROOT="/usr/local/go"
export PATH=$PATH:$GOROOT/bin
export GOPATH="/var/lib/jenkins/workspace/go"

go build'''
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

      }
    }

    stage('Test') {
      steps {
        sh '''export GOROOT="/usr/local/go"
export PATH=$PATH:$GOROOT/bin
export GOPATH="/var/lib/jenkins/workspace/go"




go test'''
      }
    }

  }
}