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
                    // this fails with 'golint: command no found'
                    //     sh 'golint .'
                    sh '/var/lib/jenkins/jobs/pipeline-jb-test-02/builds/12/pkg/mod/golang.org/x/lint@v0.0.0-20210508222113-6edffad5e616/golint/golint .'
                }
            }
        }
    }
}
