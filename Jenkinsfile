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
            sh '''export PATH=$PATH:/usr/local/bin
docker run -v "/Users/dumrongkiat/.jenkins/workspace/SayHello_master":"/etc/newman"  newman:1.0    run automationtest/SayHelloAutomationTest.postman_collection.json -d automationtest/input.json'''
          }
        }

        stage('start app') {
          steps {
            sh 'pwd'
          }
        }

      }
    }

  }
}