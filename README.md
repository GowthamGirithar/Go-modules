# Go-modules

Few points to be remembered while working on Go Modules

1. GOROOT is the installation path 
2. GOPATH is normally /users/<>/go
3. Set the environment variable for both if it is not there and also add PATH. so we can run the program from any workspace. (Normally Go Lang all the modules will be under the $GOPATH/src/ folder)
4. Package name will be unique under the folder. (one folder will have only one package name for all the go files inside that folder)
   Consider you have the module as employee-service , all the packages under the folder will be of same package and if you want to have another package name , create the folder under employee-service and then create the file with different package name
5. Package name for the test files can follow same package as file (internal test files) or new package name as <package name of the file>_test (external test file)
6. Any folder named as internal will be internal to the module and which cannot be accessed by others outside the module (way to control  external access)
7. GOPROXY -> <ul>direct or empty string - directly download the dependencies from the path mentioned in import</ul>
              <ul>off - dont download the dependencies</ul>
              <ul>any proxy url - dowload from that proxy server</ul>
8. Create the module : go mod init <path to your module directoy>      (path is only required if your module folder is in different location from GOPATH)
9. go build -> build the module and generate the executable file for the module
10. go run -> run the go files and for this you have to pass all the main package files
11. vendor :
      * It will have the local copy of all the external import files
      * To create the vendor , we have to use the command go mod vendor
      * When using modules, GOPATH is no longer used for resolving imports.
However, it is still used to store downloaded source code (in GOPATH/pkg/mod/)
and compiled commands (in GOPATH/bin)(for normal projects it will take from GOPATH/src/)
      * When you build the go module , it will take the files from GOPATH/pkg/mod
      * To look into vendor instead GOPATH/pkg/mod , you have to use the command go build -mod vendor
12. go install -> which also create the executable file but in GOPATH/bin location
13. go.sum file which will have the checksum of all the imported external dependencies
14. go.mod file which contain the imported external dependencies verison information and go version
15. go get -u or go get -u =patch commands will get the latest dependencies
16. Semantic Import Versioning : 
               Go follows version in major.minor.patch format. Major increments only if there are no backward compatability , minor for new feature , patch for bug fixes.
               When we do the major changes , version will be incremented and packages should have different structure.(for version 2.0.0 , it should have /v2 folder structure - no need any version importing till v2)
17. we cannot import the folder and only can import the package
18. To get the modules with version, we can execute the command as go get <path to the location>@v1.3.4
19. Versioning is important and so when we push the files to github, please do tagging
   *  $ git tag v1.0.0
   *  $ git push --tags
               
          
