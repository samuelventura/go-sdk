MIX_TARGET=rpi4 mix deps.get
MIX_TARGET=rpi4 mix firmware
MIX_TARGET=rpi4 mix burn
MIX_TARGET=rpi4 mix upload #works

Linux version 5.10.88-v8 (buildroot@buildroot) (aarch64-nerves-linux-gnu-gcc (crosstool-NG 1.24.0.498_5075e1f) 10.3.0, GNU ld (crosstool-NG 1.24.0.498_5075e1f) 2.37) #1 SMP PREEMPT Wed May 25 11:14:32 UTC 2022

Kernel command line: coherent_pool=1M 8250.nr_uarts=1 snd_bcm2835.enable_compat_alsa=0 snd_bcm2835.enable_hdmi=1  smsc95xx.macaddr=E4:5F:01:57:48:89 vc_mem.mem_base=0x3ec00000 vc_mem.mem_size=0x40000000  dwc_otg.lpm_enable=0 console=tty1 console=ttyS0,115200 fbcon=scrollback:1024k root=/dev/mmcblk0p2 rootfstype=squashfs rootwait consoleblank=0 quiet

WORKS OUT OF THE BOX.

Improvements

- tio says shell will run on tty1 and ask to set -c ttyS0 in erlinit.confg to run on that uart

NervesMOTD.print()
ping rpi4-4ad8
