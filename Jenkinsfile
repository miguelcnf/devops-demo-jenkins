pipeline {
  agent none
  stages {
  	stage('test') {
      agent {
        docker {
          image 'golang'
         }
      }
      steps {
        sh 'go test ./... -v'
      }
     }
    stage('build') {
      agent {
        docker {
          image 'golang'
        }
      }
      steps {
        sh 'go build -o hello hello.go'
      }
    }
    stage('deploy to staging') {
      agent any
      steps {
        sh 'echo deploying to staging'
      }
    }
    stage('wait for release approval') {
      agent none
      steps {
        timeout(time: 10, unit: 'MINUTES') {
          input(message: 'release build?')
        }
      }
    }
    stage('deploy to production') {
      agent any
      steps {
        sh 'echo deploying to production'
      }
    }
  }
}
