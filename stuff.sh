dir=$(mktemp -d /tmp/bootedseeds.XXXX)
mkdir -p $dir/bootedseeds
rg --files |
    grep -viE 'README|makefile' |
    grep -viE 'stuff.sh' |
    grep -viE '^bootedseeds$' |
    tar czf /tmp/out.tgz -T -

tar xzf /tmp/out.tgz -C $dir/bootedseeds
rg --files $dir/bootedseeds
txtar-c $dir >/tmp/bootedseeds.txtar
cat /tmp/bootedseeds.txtar | pbcopy
