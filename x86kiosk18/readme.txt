
mix archive.install hex nerves_bootstrap
mix nerves.new x86kiosk #no deps
cd x86kiosk
rm -fr deps/ .nerves/ _build/
MIX_TARGET=x86_64 mix deps.unlock --all
#update app and host names wxkiosk
#update nerves_system_x86_64 path ../
MIX_TARGET=x86_64 mix deps.get
MIX_TARGET=x86_64 mix firmware
MIX_TARGET=x86_64 mix burn -d image.img
chown samuel:samuel image.img

sudo qemu-system-x86_64 -m 512 -enable-kvm -usb -hdb /dev/sdc \
    -netdev user,id=mynet0,hostfwd=tcp::8022-:22 -device \
    virtio-net-pci,netdev=mynet0

qemu-system-x86_64 \
 -enable-kvm -m 512M \
 -netdev user,id=mynet0,hostfwd=tcp::8022-:22,hostfwd=tcp::8090-:80 \
 -device virtio-net-pci,netdev=mynet0 \
 -drive file=image.img,format=raw \
 -serial mon:stdio

qemu-virgil -enable-kvm -m 512M -device virtio-vga,virgl=on \
    -display sdl,gl=on -netdev user,id=ethernet.0,hostfwd=tcp::8022-:22 \
    -device rtl8139,netdev=ethernet.0 image.img

ssh localhost -p 8022
