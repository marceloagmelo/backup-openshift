FROM centos:7

USER root

ENV GID 23550
ENV UID 23550

ENV APP_HOME /opt/app
ENV IMAGE_SCRIPTS_HOME /opt/scripts

RUN mkdir -p $APP_HOME && \
    mkdir $IMAGE_SCRIPTS_HOME

COPY Dockerfile $IMAGE_SCRIPTS_HOME/Dockerfile
ADD go-backup-openshift $APP_HOME
COPY scripts $IMAGE_SCRIPTS_HOME

RUN groupadd --gid $GID golang && useradd --uid $UID -m -g golang golang && \
    chmod 755 $APP_HOME/go-backup-openshift && \
    chown -R golang:golang $APP_HOME && \
    chown -R golang:golang $IMAGE_SCRIPTS_HOME

EXPOSE 5000 8000

USER golang

WORKDIR $IMAGE_SCRIPTS_HOME

ENTRYPOINT [ "./control.sh" ]
CMD [ "start" ]
