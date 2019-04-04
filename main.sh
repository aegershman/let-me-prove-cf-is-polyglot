#!/usr/bin/env bash

case "$1" in
push-all)
    # push all the apps in all the directories
    # does this in parallel subshells
    for d in $(ls -d */); do
        (
            cd $d
            cf push
            cd ..
        ) &
    done
    ;;

delete-all)
    # delete all the apps in all the directories
    # does this in parallel subshells
    for d in $(ls -d */); do
        (
            cd $d
            cf delete $(yq r manifest.yml "applications[0].name") -f -r
            cd ..
        ) &
    done
    ;;

create-a-bunch-of-service-instances)
    arr=(
        "app-autoscaler standard autoscaler"
        "p.rabbitmq single-node-3.7 rmq-single"
        "p.rabbitmq clustered-3.7 rmq-cluster"
        "p.mysql db-small mysql-small"
        "p.mysql leader-follower mysql-leader-follower"
        "metrics-forwarder unlimited metrics-forwarder"
        "p.redis cache-small redis-small"
        "scheduler-for-pcf standard scheduler"
        "aws-s3 standard s3"
        "p-config-server standard config"
        "p-service-registry standard registry"
    )

    for i in "${arr[@]}"; do
        cf create-service ${i} &
    done
    ;;

*)
    exit 1
    ;;
esac
