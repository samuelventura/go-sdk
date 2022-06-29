#https://github.com/nerves-project/nerves_system_x86_64
mix archive.install hex nerves_bootstrap
mix nerves.new x86nerves #no deps
cd x86nerves
rm -fr deps/ .nerves/ _build/
MIX_TARGET=x86_64 mix deps.unlock --all
#update app and host names wxkiosk
#update nerves_system_x86_64 path ../
MIX_TARGET=x86_64 mix deps.get
MIX_TARGET=x86_64 mix firmware
MIX_TARGET=x86_64 mix burn
MIX_TARGET=x86_64 mix burn -d image.img
chown samuel:samuel image.img

#from https://github.com/nerves-project/nerves_system_x86_64/issues/129
qemu-system-x86_64 \
    -drive file=image.img,if=virtio,format=raw \
    -net nic,model=virtio \
    -net user,hostfwd=tcp::8022-:22 \
    -serial stdio

#works as well
sudo qemu-system-x86_64 \
    -drive file=/dev/sdc,if=virtio,format=raw \
    -net nic,model=virtio \
    -net user,hostfwd=tcp::8022-:22 \
    -serial stdio

#IMAGE BOOTS IN GIGABYTE BRIX AS WELL WITH EXACT SAME ISSUES AND NO ETH0 DETECTION

ssh localhost -p 8022

kex_exchange_identification: read: Connection reset by peer
erlinit: Cannot mount /dev/rootdisk0p4 at /root: Invalid argument (FIXED ON REBOOT)
warning: the VM is running with native name encoding of latin1 which may cause Elixir 
    to malfunction as it expects utf8. Please ensure your locale is set to UTF-8
    (which can be verified by runnig "locale" in your shell)
