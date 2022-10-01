# Go Linux Container

## Notes

- Namespaces - provide environmental isolation - allows processes to act without affecting the host system (unless intended) i.e. isolated to the namespace
  - PID - The PID namespace enables a view of mapped ids (as opposed to the real pid) for that process and child processes
  - MNT - Enables processes to have their own mount table, without affecting other namespaces
  - NET - Enabled processes to have their own networking stack - can be virtually linked with another namespace, with routing, which facilitates communications outside the container
  - UTS - Enables a view of the systems hostname and domain name
  - IPC - Enables a view of IPC objects/POSIX message queues (not identified by filesystem pathnames) 
  - USR - Maps the uids/gids to a different set than the host - can allow different permission e.g. pseudo-root or even allow root level permissions on the namespace resources only
  - TIME - virtualises the two system clocks

- CGroups
  - Limits on collections of processes or tasks - rules for resource sharing

- Layered Filesystems
  - Copy of root filesystem for each container e.g. btrfs, uionfs, aufs etc 



#### Sources & Resources
https://man7.org/linux/man-pages/man2/clone.2.html  
https://man7.org/linux/man-pages/man7/namespaces.7.html  
https://www.infoq.com/articles/build-a-container-golang/
