FROM alpine:3.23 AS build
WORKDIR /app

RUN apk add --no-cache openjdk11 maven

COPY ./pom.xml .
RUN mvn dependency:go-offline -B

COPY ./src ./src
RUN mvn clean package -DskipTests && rm -rf /root/.m2/repository


FROM alpine:3.23
WORKDIR /app

RUN apk add --no-cache openjdk11-jre curl && rm -rf /var/cache/apk/*

RUN adduser -S -D -H -u 10001 appuser

COPY --from=build /app/target/*.jar app.jar

USER appuser

EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
