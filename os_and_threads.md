
# Mutex

pthread_mutex_init(&mutex,&attr);
pthread_mutex_destroy(&mutex);  /* undefined if mutex is locked at destroy ! */

pthread_once(&pthread_once_t,function_p);

pthread_mutex_lock(&mutex);
pthread_mutex_trylock(&mutex);
pthread_mutex_unlock(&mutex);

# Cond-wait

```c
int pthread_cond_wait(pthread_cond_t *cptr, pthread_mutex_t *mptr );
int pthread_cond_timedwait(pthread_cond_t *cptr, pthread_mutex_t *mptr , const timespec *);
int pthread_cond_signal (pthread_cond_t *cptr);
int pthread_cond_broadcast (pthread_cond_t *cptr);    /* wakes up all threads waiting */
```

eg: cond_wait in prod/consumer problem

```c
struct {
  pthread_mutex_t mutex;
  pthread_cond_t cond;
  int nready;                   /* number ready for consumer */
} nready = {
PTHREAD_MUTEX_INITIALIZER, PTHREAD_COND_INITIALIZER
};
```

## producer:

```c
pthread_mutex_lock(&nready.mutex);
if (nready.nready == 0) {
  pthread_cond_signal(&nready.cond);            /* Note -- Signalling does nothing about the mutex . You can signal outside the mutex too */
}
nready.nready++;
pthread_mutex_unlock(&nready.mutex);
```

### Alternate:

```c
auto int do_signal = 0;

pthread_mutex_lock(&nready.mutex);
if (nready.nready == 0) {
  do_signal = 1;
}
nready.nready++;
pthread_mutex_unlock(&nready.mutex);

if (do_signal) {
  pthread_cond_signal(&nready.cond);            /* Note -- You can signal outside the mutex too */
}
```

## consumer:

```c
pthread_mutex_lock(&nready.mutex);
while (nready.nready == 0)                            /* while(ready-var) helps in preventing spurious wake up calls */
  pthread_cond_wait(&nready.cond, &nready.mutex);     /* Note: wait releases the mutex. So always call wait within mutex.
                                                         And wait will re-acquire mutex when it returns */
nready.nready--;
pthread_mutex_unlock(&nready.mutex);
```

# read-write-locks

pthread_rwlock_t


# Thread creation

```c
pthread_t /* data-type to represent a thread's id */

int pthread_create(pthread_t *thread,                             /* result arg - in which the child's id is placed */
                   const pthread_attr_t *attr,
                   void *(*start_routine) (void *), void *arg);

pthread_t pthread_self(void);
int pthread_equal(pthread_t lhs, pthread_t rhs);
```

Termintate thread
* fall off its start routine
* pthread_exit
        void pthread_exit(void *retval);        /* ret-val is passed to the thread that will call pthread_join */
* pthread_cancel

Keeping a threads detached is good, as exited thread resources are help up (like zombie processes) till someother
thread called pthread_join()

# Difference between select/poll

Good write up [here]https://daniel.haxx.se/docs/poll-vs-select.html

Dev-poll splits the telling of which fd's and actual waiting into 2 sys-calls.

```c
int select(int nfds, fd_set *readfds, fd_set *writefds, fd_set *exceptfds, struct timeval *timeout);
void FD_CLR(int fd, fd_set *set);
int  FD_ISSET(int fd, fd_set *set);
void FD_SET(int fd, fd_set *set);
void FD_ZERO(fd_set *set);
int pselect(int nfds, fd_set *readfds, fd_set *writefds, fd_set *exceptfds, const struct timespec *timeout, const sigset_t *sigmask);
```
* select uses timeval, while pselect uses timespec
* select updates timeval to reflect how much time was left, while pselect doesn't change timeout.
  This is linux behavior. POSIX per se allows leaving timeval or updating it.
* if timeout is NULL, then select blocks. If both sec/usec is 0, then select returns immdly
* select returns the total of all descriptors in read/write/error
* pselect avoid the signalling race-cond by clubbing sigprocmask in one go.
* FD_SETSIZE is the max size of fd_set

```c
int poll(struct pollfd *fds, nfds_t nfds, int timeout);
int ppoll(struct pollfd *fds, nfds_t nfds, const struct timespec *timeout, const sigset_t *sigmask);
struct pollfd {
    int   fd;         /* file descriptor */
    short events;     /* requested events */
    short revents;    /* returned events */
};
```
* timeout is in milliseconds
* events is one of POLLIN/POLLOUT/+ lots of POLLERR, POLLHUP

## epoll

```c
int epoll_create(int size);
int epoll_create1(int flags);
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);
int epoll_wait(int epfd, struct epoll_event *events, int maxevents, int timeout);
int epoll_pwait(int epfd, struct epoll_event *events, int maxevents, int timeout, const sigset_t *sigmask);
```

* Create a epoll instance. size is old and ignored. flags is for setting O_CLOEXEC
* ctl sets up a fd, with one of the options and a usr-ata
* wait gets the list of fd along with the events in it.
* level triggered is default. Edge triggered is useful to detect change of status. This is useful,
  if we dont want to read/write right-away, note down that its ready. Such fd's arent notified as
  ready until we get a EGAGIN on a subsequent read/write.


# Different between mutex and semaphore

Good explanation at http://stackoverflow.com/questions/62814/difference-between-binary-semaphore-and-mutex/86021#86021

* At the outset, mutex is for exclusive access, semaphore is for a limiting access to a certain count.

```c
Mutex - Exclusive access.

Lock
Work
UnLock

Semaphore ( has an associated count against it).

take()
work();
give()
```

* Other characteristic is that mutex is used for exclusion(locking), while semaphare is used for signalling.
  The first task waits till something is ready, while the second task signals when that is available.
* Mutex is typically same-task should release. (true to spirit of locking).
  Semaphore can be given by one task, while taken by another task (true to spirit of signalling).
* A mutex is allowed to be re-entrant in some implementation(eg. recursive
  mutex is posix), but sem's aren't (as they can be signalled by any task)

If locking is what is required, use a mutex and not a semaphore with count 1. (Why? To Read)

## Other questions around here

1. Can a thread acquire more than one lock (Mutex)?
   * Yes
2. Can a mutex be locked more than once?
   * Normal mutex- no. Recursive mutex-yes. Must be unlocked as many times.
3. What happens if a non-recursive mutex is locked more than once.
   * Deadlock.
4. Binary semaphore and mutex is same?
   * No.
5. Mutext and critical section.
   * Synonyms, for purpose. BUt some OS crit-section means, disable interrupts
     and run. So cheap but may make CPU non-responsive!
6. Can we acquire mutex/semaphore in an Interrupt Service Routine?
   * Not recommended. However, a ISR can typically signal a semaphore.
7. Should a thread block if a resource is unavailable?
   * Not necessary. That's why posix has try_lock() which will return -ve immdly.
     The thread as per design can go do other things.

# Different types of synchronization primitives

atomic variables
spin locks
mutexes
 * simple
 * recursive
 * read-write
 * blocking/timed_wait
semaphores

--To read:
* critical section
* barrier

c++ 14

std::mutex
std::timed_mutex
std::recursive_mutex
std::recursive_timed_mutex
std::shared_timed_mutex

`std::lock_guard<>`
`std::unique_lock<>`
`std::shared_lock<>`

## Difference between Process and Threads

Resources that are shared
 (man clone(2) gives a bunch of them)
  - Stack / Execution unit / Registers
  - Memory space
  - Fds that are open
    - regular files
    - sockets
    - Pipes
    - memory mapped files
  - Signals that are pending
  - Signal handling disposition
  - Locks (memory locks/file locks)
  - Process stats
  - Stuff like atexit behavior().

When we say a thread, generally its just has only the stack that is private. Otherwise everything else is shared.
Processes on the other hand, have everything for itself.
 However, when a process forks, certain things are shared between parent & child.

# Context switch

In linux traditionally the ticks are tracked on time spent on user-space by a
process, kernel-space by a process. Sometimes interrupts are processed in the context
of a process in kernel-space. Typically the cpu of a process is the (user-time+kernel-time)/clock-time


# Dining Philosopher Issue

Problem:
* Circular table, Every philosopher needs both side chopsticks. If every one acquire left side
  and then try to take right side, it results in a dead lock.
Solution
* One idea is to release a left chopstick if right is not obtained. While this will solve the issue
  most of the time, still theoretically, each philosopher may still take the left first, drop left,
  take left again, drop left infinitely.
* Another idea is to number chopsticks 0 to n-1 and have the philosophers pick the lower number first
  and higher number next. Every philosopher will pick in one order except the last philosohper who
  will pick in the other order, thus breaking the cycle.

# Deadlock-free class

* Design a lock class that will lock, only if there is no dead-lock.
Solution
* Build a directed graph, where a edge A->B represents lock A is taken first and then lock B.
* Keep buiding the graph, based on lock order. You have a possible deadlock, if a cycle gets
  formed when adding a edge.
  (We can use DFS to detect the cycle).

# Lockless allocator

Basically its a fixed size mem-pool. But there aren't locks to alloc/free memory.

* The free/taken objects are tracked by bits. For ex, a 64-bit WORD represents the taken/free status of 64 blocks.
* When an alloc-is issued from thread-x, this walk over the bit-array. It does a atomic::exchange of 0 on a word.
  And then checks whether any 1 bit is available in this word. If so, it takes that location, clears the bit,
  and then does a atmoic::OR of the 1-bit cleared word back. If this word has no free bits, it moves on to the next word.
* When a free is issued from thread-y, this does a atomic:OR on the bit its free'ing.

Flip-sides:

* Each thread holds one entire 32/64-block - depending on the atomicity of the underlying word. If all free blocks are
  only on the last word, then a thread declares no more memory although there is.
* Each thread iterates the entire bit-field before finding out next free, instead of walking over a simple free list
  under a guard(!)

# Single-Producer/Single-Consumer Bit flag


# Call-Order

Assume a signleton class. It has 3 methods. first(), second(), third(). How do u ensure,
they are strictly called in order. first should finish, and then second should proceed etc..
assumed 3 threads will independantly endup calling the functions.

Hold 2 semaphores (cond_wait(). Signl them after first is complete, and second is complete()


# Synchronization Methods

We have a class with one serialized methond and one normal one. Can 2 threads execute the 2
methonds concurrently
Answer:
Basically serialized is like holding a instance level lock and then executing the fucntion.
So, while the call itself will happen, whether it will produce desired results depends on
the design on the class and what elements are accessed.

# FizzBuzz

You have 4 threads. One tread prints multiples of 3&5, one that of 3, one that of 5 and one
the result. How will you make all 4 synchronize.
Answer:
The base variable that is incrementing should be same. Each thread has the condition as a
local variable, and if its condition matches it preints that.

Here is the snippet

```c
  lock.acquire();
  if (my_condition_satisfies){
    print("%d - %s",current,mystring);
    current++;
  }
  lock.release();
```

Note that one thread will take the lock. But it will silently do nothing, if its conditions
didn't match.  The flipside is that there are too many spurious runs. To think of something
more efficient.

# Linux and threads

* pthread defines the notion of a processid and a thread id.
   process id is globally unique in system, while thread-id is local to the process.
  Linux typically allows us to create a thread using clone() system call.
   To emulate posix style process/thread -> linux does the thread grouping, and uses the
     pid of the prcess as the thread id and hte thread-group's pid as the process pid.

* Linux Threads was a first suppot for posix apis. This is superseded by NTPL from 2.6 onwards
* NTPL (Native Threads for Posix in Linux) is the new implentattion that is implemented over
  clone() and futex() system calls


# Fork in multi-threaded program

* fork() in a multi threaded program is mostly a disaster.
  -> only calling thread is alive in the child.
  -> only async functions can be called in child.
    -> malloc, printf, syslog etc.. aren't async!
  -> possible only useful reason to fork() is to exec(). Even then FD's should be closed at exec.
     There may be race between parent's one thread's open()and fctnl() before which another thread
     could have called fork.

One useful function that tries to solve the problem with fork(2) in
multi-threaded programs is pthread_atfork(). It has the following prototype:


int pthread_atfork(void (*prepare)(void), void (*parent)(void), void (*child)(void));

It allows to set handler functions that will be automatically executed on fork call:
prepare - Called just before a new process is created.
parent - Called after a new process is created in the parent.
child - Called after a new process is created in the child.

The purpose of this call is to deal with critical sections of a multi-threaded
program at the time fork(2) is called. A typical scenario is when in the
prepare handler mutexes are locked, in the parent handler unlocked and in the
child handler reinitialized.


# Signals

* In a process with many threads


# Time

time_t -> seconds
timeval -> micro sec
timespec -> nanonsec

rdtsc -> a monotic counter from cpu
  -> but not necessarily accurate. It may drift based on cpu's turbo speeding.

clock_gettime


memory barrier/compiler memory barrier

rdtsc
rdtscp

## Read/Write

sendmmsg() is a api to send multiple messages in one go.

```c
int sendmmsg(int sockfd, struct mmsghdr *msgvec, unsigned int vlen, unsigned int flags);

/*msg_hdr itself is as per sendmsg */
struct mmsghdr {
    struct msghdr msg_hdr;  /* Message header */
    unsigned int  msg_len;  /* Number of bytes transmitted */
};
```


