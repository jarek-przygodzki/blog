Vaadin CDI application build with Gradle.

Related to blog post [WildFly, Vaadin CDI i WebSockety](https://jarekprzygodzki.wordpress.com/2015/02/22/wildfly-vaadin-cdi-i-websockety/)

# Build WAR
```
gradle war
```

# Import project into Eclipse
- completely rewrite existing Eclipse config files (if any)
```
gradle cleanEclipse eclipse
```
- import into Eclipse *File -> Import... -> General -> Existing Project into Workspace*