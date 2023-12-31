# Go-Inca

---

### Features 💡

```
1. Set (key, val)
2. Get (key)
3. Set TTL(key, time)
4. Remove TTL(key)
5. Delete (key)
6. Set Eviction Policy (LRU/LFU)
    - LRU:
        - Capacity[MB]
    - LFU
        - Capacity[MB]
```

### Currently working on ⚙️

| Feature                                                      | Done(✅/❌) |
| ------------------------------------------------------------ | ----------- |
| Implement LRU                                                | ✅          |
| Implement LFU                                                | ❌          |
| Implement Set, Get and Delete for LRU                        | ✅          |
| Implement Set, Get and Delete for LFU                        | ❌          |
| Implement TTL                                                | ✅          |
| Optimize caching system                                      | ✅          |
| Add support for more data structures                         | ❌          |
| Make a network interface to communicated via TCP(Text Based) | ❌          |
| Make it work flawless with multiple concurrent users         | ❌          |
| Refactor DLL                                                 | ✅          |
| Refactor LRU                                                 | ✅          |
| Refactor Parser                                              | ✅          |
| Implement a parser to read text commands                     | ✅          |

### Dev logs 🗒️

```
Tue 04 Jul 2023 12:47:42 PM IST
I have no idea about what I'm getting into, but yeah.
I'mma try my best to learn as much as I can, no matter what :)
```

```
Thu 06 Jul 2023 02:35:09 AM IST
Just implemented a doubly linkedlist and lru cache with basic functionality. Still hasn't found out the way to implement TTL.
This shit is fun af.
```

```
Thu 06 Jul 2023 01:43:52 PM IST
Everything thing is fast af, but I feel like I can improve it. Benchmarks are showing good results. I should do something to improve setting keys after capacity is over.
```

```
Thu 06 Jul 2023 02:37:05 PM IST
Just found out that I'm doing a lot of allocation,
which can worsen the performance. I think a better way to approach this issue is oncee you reach the capacity, then you should just reuse the allocated Node instead of just allocating new one.
```

```
Thu 06 Jul 2023 02:55:01 PM IST
Caught the culprit which made my code slower.
It's the fucking `Pop()` method. It uses 'Tail()' method which is an `O(n)` method,
which made setting new values relatively slow. I'm such a dumb fuck.
I didn't notice it there. Now the better way to implement it is by
keeping track of the `TailNode` in LRUCache struct.
```

```
Thu 06 Jul 2023 03:11:37 PM IST
DIGGITY DONG! Just made that silly change and now my code is 3.53438442983 × 10^9 times faster, like fuck, that's nearly 3.5  × 10^9 times faster, wtf?!?
```

```
Thu 06 Jul 2023 03:35:40 PM IST
LOL, I messed up. I think i fucked up the whole dll by doing
some stupid shit. My brain is quite literally fried from coding for the past 10 hours. I need a break :) A quick guess tho. If my current code is working fine, then it's 1.96x faster than my previous working version(still, that's a lot of improvement even tho it's not in the magnitude of 9)
```

```
Thu 06 Jul 2023 04:30:00 PM IST
Haven't stopped coding yet, and just found out the list provided by go ain't gonna be that useful for my purpose
```

```
Thu 08 Jul 2023 04:30:00 AM IST
Implemented a TTL system. Also found many bugs in the Set method as I was in a hurry to create a TTL system. Learned that always be very careful when you are modifying a struct in-place and also be careful with accessing struct fields in if-else statements. Both of these things can land you in very bad situations
```

```
Mon 17 Jul 2023 03:49:40 PM IST
Forgot to log days. Implemented a simple command parser. There might be some uncaught bugs, but it's working as of now. Should've made a rule/syntax layout before hand starting the parser. but nonetheless, it's working!
```

```
Tue 18 Jul 2023 04:05:41 PM IST
Found a flaw in the system. If I'm trying to access length and if a key's TTL is already expired, the key won't be removed until it's accessed or get evicted by the LRU system.
This means if my capacity(C) = 3 and key1 is already expired, the length would still be 3 since there's no background service checking for eviction
```

```
Tue 18 Jul 2023 04:43:26 PM IST
I don't know whether this is true or not, but i just did a bench mark test and the results are pretty, good, atleast.
for the `IsFull()` method with O(N) complexity, BenchmarkSetKeyfor10xCapacity-16 gave around `1690250273 ns/op`, which is atrocious for a cache. But ever since if made the method O(1) by keeping track of the length, the bench mark gave me `0.2400 ns/op`, which is wayyy better than the previous one. Since there's no O(N) operation being done, it must've made a real difference in the algorithm. I don't know how to check it properly. I guess I should ask someone from discord to help me with it.
```
