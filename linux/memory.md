# 1. Memory

## 1.1 Free

메모리가 커널에서 제공하는 중요한 리소스 중 하나로 CPU가 프로세스의 연산 과정에 필요한 리소스라고 정의한다면, 메모리는 프로세스가 연산할 수 있는 공간 (`Area`)을 제공해 주는 리소스라고 볼 수 있다. 프로세스는 메모리라는 공간에 필요한 함수를 넣어 두거나 변수에 값을 저장하거나 하는 방식으로 연산을 위한 공간을 확보하고 작업을 진행하기 때문이다.

해당 사용중인 메모리가 부족하게 된다면 프로세스가 더 이상 연산을 위한 공간 확보를 할 수 없기 때문에, 시스템 호출 혹은 응답 불가 현상이나 성능 저하가 발생할 수 있기 때문에 지속적인 메모리 가용률을 확인해야 한다.

```bash
free -<options>
```

```bash
geonwoo@WORK-DESKTOP:~$ free
               total        used        free      shared  buff/cache   available
Mem:        15989336      658788    14933336        4172      599524    15330548
Swap:        4194304           0     4194304
```

<details>
<summary>Memory active size</summary>

| Options | Description |
| - | - |
| -b (`Byte`) | Memory의 사이즈 표시 단위를 `Byte`로 출력 |
| -k (`KiloByte`) | Memory의 사이즈 표시 단위를 `KiloByte`로 출력 |
| -g (`GigaByte`) | Memory의 사이즈 표시 단위를 `GigaByte`로 출력 |

```bash
geonwoo@WORK-DESKTOP:~$ free -k
               total        used        free      shared  buff/cache   available
Mem:        15989336      638884    15388888        4176      160060    15350452
Swap:        4194304           0     4194304
```

</details>


# 2. PID

## 2.1 Create PID unit count

```bash
top -<options> -<options> <int>
```

```bash
geonwoo@WORK-DESKTOP:~$ top
top - 13:30:09 up  4:14,  1 user,  load average: 0.03, 0.01, 0.00
Tasks:  36 total,   1 running,  35 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.0 us,  0.0 sy,  0.0 ni,100.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
MiB Mem :  15614.6 total,  14535.7 free,    663.6 used,    616.1 buff/cache
MiB Swap:   4096.0 total,   4096.0 free,      0.0 used.  14951.0 avail Mem

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
      1 root      20   0   24248  15548  11680 S   0.0   0.1   0:00.75 systemd
      2 root      20   0    3760   2336   2204 S   0.0   0.0   0:00.00 init-systemd(Ub
      3 root      20   0    3612   2092   2004 S   0.0   0.0   0:00.00 init-watcher
      7 root      20   0    3792   2324   2200 S   0.0   0.0   0:00.00 init
     47 root      19  -1   50688  17412  16112 S   0.0   0.1   0:00.35 systemd-journal
     65 systemd+  20   0   22416  14504  12012 S   0.0   0.1   0:00.08 systemd-resolve
     88 root      20   0   34768  12296   9228 S   0.0   0.1   0:00.41 systemd-udevd
<...>
```

`top`를 통해 시스템에서 실행 중인 프로세스들의 목록과 우선순위들을 볼 수 있는데 [1. Memory](#11-free)의 

```bash
sudo sysctl -a | grep -i pid_max
```

```bash
kernel.pid_max = 4194304
```

`sudo sysctl -a | grep -i pid_max`은 사용중인 시스템에서 생성가능한 PID 갯수를 출력해주지만 좀비 프로세스가 많으면 많을 수록 PID 자체 할당이 불가능하여 지속적인 정리를 권장한다. 그렇다면 해당 부하의 문제를 확인하기 위해서 `vmstat`를 활용하는데 `uptime`과 다르게 I/O를 발생하는 스크립트를 실행시키고 결과를 출력해준다

```bash
geonwoo@WORK-DESKTOP:/tmp$ vmstat 1
```

```bash
procs -----------memory---------- ---swap-- -----io---- -system-- -------cpu-------
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st gu
 0  0      0 14888776  19096 611956    0    0    79  1219  126    0  0  0 100  0  0  0
 1  0      0 14888524  19096 611996    0    0     0     0  235  125  0  0 100  0  0  0
 0  0      0 14888524  19096 611996    0    0     0     0  133   75  0  0 100  0  0  0
 0  0      0 14888524  19096 611996    0    0     0     0  119   52  0  0 100  0  0  0
 0  0      0 14888524  19096 611996    0    0     0     0  118   50  0  0 100  0  0  0
 0  0      0 14888524  19096 611996    0    0     0     0  132   94  0  0 100  0  0  0
 <...>
```