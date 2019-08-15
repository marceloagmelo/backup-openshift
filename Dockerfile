FROM ocpsatpvlbr01.dcbr01.corp:5000/santander-ocp3-rhel7

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

#######################################################################
##### We have to expose image metada as label and ENV
#######################################################################
LABEL br.com.santander.imageowner="Corporate Techonology" \
      br.com.santander.description="Golang 1.10.2 runtime for node microservices - backup openshift" \
      br.com.santander.components="Golang Server" \
      br.com.santander.image="registry.cmpn.paas.gsnetcloud.corp/santander/go-backup-openshift:2.0.3.RELEASE"

ENV br.com.santander.imageowner="Corporate Techonology"
ENV br.com.santander.description="Golang 1.10.2 runtime for node microservices - backup openshift"
ENV br.com.santander.components="Golang Server"
ENV br.com.santander.image="registry.cmpn.paas.gsnetcloud.corp/santander/go-backup-openshift:2.0.3.RELEASE"

EXPOSE 5000 8000

USER golang

WORKDIR $IMAGE_SCRIPTS_HOME

ENTRYPOINT [ "./control.sh" ]
CMD [ "start" ]
