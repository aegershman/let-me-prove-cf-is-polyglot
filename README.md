# let-me-prove-cf-is-polyglot

It's not lock-in... It's digiorno. (Haha. Just kidding. It's actually Cloud Foundry.)

```bash
# push all the apps in all the directories
# does this in parallel subshells
for d in $(ls -d */); do
    (
        cd $d
        cf push
        cd ..
    ) &
done

# delete all the apps in all the directories
# does this in parallel subshells
for d in $(ls -d */); do
    (
        cd $d
        cf delete $(yq r manifest.yml "applications[0].name") -f -r
        cd ..
    ) &
done
```
