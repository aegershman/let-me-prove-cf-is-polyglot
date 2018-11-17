# java-boot

If your `.m2/settings.xml` file is pointing to a maven mirror (e.g., corporate maven), you _might_ need to temporary remove it. E.g., change the name to something like `.m2/settings-corp.xml`, then when you're done with this project you can change it back to `.m2/settings.xml`.

local:

```bash
./mvnw clean install
./mvnw spring-boot:run
# http://localhost:8080
```

cloud-foundry:

```bash
./mvnw clean install
cf push
```
