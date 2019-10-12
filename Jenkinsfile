pipeline {
  agent any
  stages {
  	stage('test') {
      steps {
        sh 'echo testing'
      }
     }
    stage('build') {
      steps {
        sh 'echo building'
      }
    }
    stage('deploy to staging') {
      steps {
        sh 'echo deploying to staging'
      }
    }
    stage('wait for release approval') {
      steps {
        timeout(time: 10, unit: 'MINUTES') {
          input(message: 'release build?')
        }
      }
    }
    stage('deploy to production') {
      steps {
        sh 'echo deploying to production'
      }
    }
  }
}
