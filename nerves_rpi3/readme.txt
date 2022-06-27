MIX_TARGET=rpi3 mix deps.get
MIX_TARGET=rpi3 mix firmware
MIX_TARGET=rpi3 mix burn
MIX_TARGET=rpi3 mix upload #works

Linux version 5.10.88 (buildroot@buildroot) (armv7-nerves-linux-gnueabihf-gcc (crosstool-NG 1.24.0.498_5075e1f) 10.3.0, GNU ld (crosstool-NG 1.24.0.498_5075e1f) 2.37) #1 SMP PREEMPT Wed May 25 11:14:32 UTC 2022

Kernel command line: coherent_pool=1M snd_bcm2835.enable_compat_alsa=0 snd_bcm2835.enable_hdmi=1 bcm2708_fb.fbwidth=640 bcm2708_fb.fbheight=480 bcm2708_fb.fbswap=1 smsc95xx.macaddr=B8:27:EB:CF:EE:0C vc_mem.mem_base=0x3ec00000 vc_mem.mem_size=0x40000000  console=tty1 console=ttyAMA0,115200 fbcon=scrollback:1024k root=/dev/mmcblk0p2 rootwait consoleblank=0 quiet

WORKS OUT OF THE BOX.


NervesMOTD.print()

ping rpi3-ee0c
ssh rpi3-ee0c
