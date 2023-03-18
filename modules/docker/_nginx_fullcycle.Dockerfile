FROM nginx:latest

COPY /modules/docker/html /usr/share/nginx/html

ENTRYPOINT [ "/docker-entrypoint.sh" ]

CMD [ "nginx", "-g", "daemon off;" ]