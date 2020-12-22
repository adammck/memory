# Manual Testing

I'm trying to make sure that this works with both cgroups v1 and v2, whether
run via systemd or Docker or Kubernetes. One might think that would all be the
same thing, but apparently not. I suspect that each one performs its own
hackery on top of the actual cgroups interface.

So I'm using Vagrant to manually test every combination.
TODO(adammck): Automate this!

```
$ vagrant up focal
$ vagrant ssh focal

# Dump some info
$ uname -r
$ grep cgroup /proc/filesystems
$ systemctl --version | head -n 1

# TODO(adammck): Move this to provision.sh
$ curl -fsSL https://get.docker.com -o get-docker.sh
$ sudo sh get-docker.sh

# TODO(adammck): Move this to provision.sh
$ pushd src/memory/example
$ go build
$ popd

# Run it in three different ways
$ src/memory/example/example
$ sudo systemd-run --quiet --pipe --wait -p MemoryHigh=5M -p MemoryMax=10M src/memory/example/example
$ sudo docker run -it --rm --memory=20M --volume=$HOME/src/memory/example:/root/example bash:4.4 /root/example/example
```

## Results

1. ssh in
2. run via nothing
3. run via `systemd-run`
4. run via `docker run`

### Focal

```
$ vagrant ssh focal

$ uname -r
5.4.0-52-generic

$ grep cgroup /proc/filesystems
nodev   cgroup
nodev   cgroup2

$ systemctl --version | head -n 1
systemd 245 (245.4-4ubuntu3.2)

$ src/memory/example/example
limit: none
usage: 3792896B

$ sudo systemd-run --quiet --pipe --wait -p MemoryHigh=5M -p MemoryMax=10M src/memory/example/example
limit: 10485760B
usage: 778240B

$ sudo docker run -it --rm --memory=20M --volume=$HOME/src/memory/example:/root/example bash:4.4 /root/example/example
WARNING: Your kernel does not support swap limit capabilities or the cgroup is not mounted. Memory limited without swap.
limit: 20971520B
usage: 2314240B
```

### Bionic

```
$ vagrant ssh bionic

$ uname -r
4.15.0-122-generic

# we have cgroups v2 now
$ grep cgroup /proc/filesystems
nodev   cgroup
nodev   cgroup2

$ systemctl --version | head -n 1
systemd 237

$ src/memory/example/example
limit: none
usage: 433532928B

$ sudo systemd-run --quiet --pipe --wait -p MemoryHigh=5M -p MemoryMax=10M src/memory/example/example
limit: 10485760B
usage: 778240B

$ sudo docker run -it --rm --memory=20M --volume=$HOME/src/memory/example:/root/example bash:4.4 /root/example/example
WARNING: Your kernel does not support swap limit capabilities or the cgroup is not mounted. Memory limited without swap.
limit: 20971520B
usage: 1560576B
```

### Xenial

```
$ vagrant ssh xenial

$ uname -r
4.4.0-193-generic

# only v1
$ grep cgroup /proc/filesystems
nodev   cgroup

$ systemctl --version | head -n 1
systemd 229

$ src/memory/example/example
limit: none
usage: 663846912B

# no --pipe, no --wait, no MemoryHigh
$ sudo systemd-run -p MemoryLimit=10M ./example
$ journalctl -u run-r324f4b062c5a42bcb155f7e6660ab284.service
-- Logs begin at Tue 2020-12-22 07:36:55 UTC, end at Tue 2020-12-22 07:39:01 UTC. --
Dec 22 07:39:01 vagrant systemd[1]: Started /home/vagrant/src/memory/example/./example.
Dec 22 07:39:01 vagrant example[1665]: limit: 10485760B
Dec 22 07:39:01 vagrant example[1665]: usage: 524288B

$ sudo docker run -it --rm --memory=20M --volume=$HOME/src/memory/example:/root/example bash:4.4 /root/example/example
limit: 20971520B
usage: 1953792B
```

### Trusty

```
$ vagrant ssh trusty

$ uname -r
4.4.0-31-generic

$ grep cgroup /proc/filesystems
nodev   cgroup

# trusty is pre-systemd!
$ systemctl --version | head -n 1
-bash: systemctl: command not found

# ----
# before running get-docker.sh...
# ----

$ cat /proc/self/cgroup
1:name=systemd:/user/1000.user/3.session

$ src/memory/example/example
error reading limit: memory subsystem not present
error reading usage: memory subsystem not present

# ----
# after running get-docker.sh...
# ----

$ cat /proc/self/cgroup
13:pids:/
12:hugetlb:/
11:net_prio:/
10:perf_event:/
9:net_cls:/
8:freezer:/
7:devices:/
6:memory:/
5:blkio:/
4:cpuacct:/
3:cpu:/
2:cpuset:/
1:name=systemd:/user/1000.user/3.session

$ src/memory/example/example
limit: none
usage: 811102208B

# ----

# still no systemd!
$ sudo systemd-run -p MemoryLimit=10M ./example
sudo: systemd-run: command not found

$ sudo docker run -it --rm --memory=20M --volume=$HOME/src/memory/example:/root/example bash:4.4 /root/example/example
WARNING: Your kernel does not support swap limit capabilities or the cgroup is not mounted. Memory limited without swap.
limit: 20971520B
usage: 5271552B
```
