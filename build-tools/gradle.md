* Convert Maven build to gradle

  ```gradle init```

* Initialise new project folder structure (Java example)

  ```gradle init --type java-library```

* List all projects in the build

  ```gradle -q projects```

* List all properties on the build

```gradle -q properties```

* List build configuration model

  ```gradle -q model```

* Execute one integration test:

  ```./gradlew integrationTest --tests=*SomeTestClass```

* Execute tasks parallel:

  ```./gradlew clean build --parallel```

* Skip test:

  ```gradle build -x test```

* Include console and logging output for container code and tests

  ```gradle functionalTest --info```

* Run application tests with debugger agent

  ```gradle test --debug-jvm```

* Refresh dependencies

  ```./gradlew clean build --refresh-dependencies```

* View dependencies for specific jar (and configuration)

  ```gradle my-sub-project:dependencyInsight --configuration testCompile --dependency hamcrest-core```

* Run in custom environment

  ```-Dspring.profiles.active=dev ```


## Configuration Cheatsheet

* Better logging in tests:

```
test {
    testLogging {
        exceptionFormat = 'full'
    }
}
```
