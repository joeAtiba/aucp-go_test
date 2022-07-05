pipeline {
    agent any
    
    tools {
        go 'go1.18'
    }
    
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    
    stages {
        stage('Git') {
            steps {
                git branch: 'main', url: 'https://github.com/joeAtiba/aucp-go_test'
            }
        }
        
        stage('Debugging') {
            steps {
                echo 'Here is JENKINS_HOME:'
                echo "${JENKINS_HOME}"
                echo 'Here is GOPATH:'
                echo "${GOPATH}"
                sh 'ls -lasF $GOPATH'
            }
        }
        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }

        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go build'
            }
        }
        
        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet .'
                    echo 'Running linting'
                    sh 'golint .'
                }
            }
        }
    }
}
