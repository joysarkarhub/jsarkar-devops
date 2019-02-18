#!/bin/sh
head -1 /proc/self/cgroup|cut -d/ -f3 > containerid
