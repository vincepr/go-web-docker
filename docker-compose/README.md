# Docker compose

## basic structures
### Services
- computing components of an application. Basically our container images
- services communicate with each other trough Networks
### Networks
- in this setting a Network enables IP routes between containers
### VOLUMES
- persistent data that services may access if possible.
- basically exist on our actually system and can be accessed from there aswell
### Configs
- configs allow services to adapt their behavior without the need to rebuild the docker image.
### Secret
- configuration data, that shoud not be exposed without security consideration. 
- Secrets are made available to services as files mounted into their containers, just like configs
### Project
- A Project is an individual deployment of an application specification on a platfom.
- it is used to group ressources together and isolate them from other appliactions of the same Compose.
- A Compose implementation creating ressources on a platfrom MUST prefix resource names by project and set the label `com.docker.compose.project`
- Project name can be set explicit by top-level `name` attribute. Compose implementation MUST offer a way for users to set a custom project name and override this name. this allows running the same compose.yaml twice on the same infrastructure without changes.

https://docs.docker.com/compose/compose-file/