FROM google/cloud-sdk:alpine

RUN apk --update add openjdk8-jre
RUN gcloud components install --quiet beta pubsub-emulator
RUN mkdir -p /var/pubsub

EXPOSE 8086

CMD [ "gcloud", "beta", "emulators", "pubsub", "start", "--host-port=0.0.0.0:8086", "--data-dir=/var/pubsub"]
